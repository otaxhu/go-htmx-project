package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/otaxhu/go-cssmodules"
)

func main() {
	inputDir := flag.String("input-dir", "", "`directory` where all of the input templates and css files are.\nThe directory will be walked recursively")
	outputDir := flag.String("output-dir", "", "`directory` of output templates")
	outputCSSPath := flag.String("output-css-path", "", "`output-path` of your css file")

	flag.Parse()

	if *inputDir == "" || *outputCSSPath == "" || *outputDir == "" {
		fmt.Fprint(os.Stderr, "error: -input-dir, -output-dir and -output-css-path flags are required")
		os.Exit(2)
	}

	absInputDirPath, err := filepath.Abs(*inputDir)
	if err != nil {
		log.Fatal(err)
	}

	absOutputDir, err := filepath.Abs(*outputDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll(absOutputDir); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(absOutputDir, 0775); err != nil {
		log.Fatal(err)
	}

	absOutputCSSPath, err := filepath.Abs(*outputCSSPath)
	if err != nil {
		log.Fatal(err)
	}

	outputCSSFile, err := os.Create(absOutputCSSPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputCSSFile.Close()

	err = filepath.WalkDir(absInputDirPath, func(pathFile string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		file, err := os.Open(pathFile)
		if err != nil {
			return err
		}
		defer file.Close()

		if strings.HasSuffix(pathFile, ".html") {
			if !strings.HasSuffix(pathFile, ".module.html") {

				name := strings.TrimSuffix(d.Name(), ".html")

				noModuleHtmlFile, err := os.Create(filepath.Join(absOutputDir, name+"_"+uuid.NewString()+".html"))
				if err != nil {
					return err
				}
				defer noModuleHtmlFile.Close()

				if _, err := noModuleHtmlFile.ReadFrom(file); err != nil {
					return err
				}
				return nil
			}
			return nil
		}

		if !strings.HasSuffix(pathFile, ".css") {
			return nil
		}

		basenameAbsPath, found := strings.CutSuffix(pathFile, ".module.css")
		if !found {
			if _, err := outputCSSFile.ReadFrom(file); err != nil {
				return err
			}
			outputCSSFile.WriteString("\n")
			return nil
		}

		htmlFile, err := os.Open(basenameAbsPath + ".module.html")
		if errors.Is(err, os.ErrNotExist) {
			_, err := cssmodules.NewCSSModulesParser(file).ParseTo(outputCSSFile)
			outputCSSFile.WriteString("\n")
			return err
		} else if err != nil {
			return err
		}
		defer htmlFile.Close()

		scopedClasses, err := cssmodules.NewCSSModulesParser(file).ParseTo(outputCSSFile)
		if err != nil {
			return err
		}
		outputCSSFile.WriteString("\n")

		parsedHtmlFile, err := os.Create(filepath.Join(absOutputDir, filepath.Base(basenameAbsPath)+"_"+uuid.NewString()+".html"))
		if err != nil {
			return err
		}

		return cssmodules.NewHTMLCSSModulesParser(htmlFile, scopedClasses).ParseTo(parsedHtmlFile)
	})
	if err != nil {
		if err := os.RemoveAll(absOutputDir); err != nil {
			log.Fatal(err)
		}
		if err := os.Remove(absOutputCSSPath); err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
}
