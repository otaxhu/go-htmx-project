build_project: build_tailwind
	go build -o ./dist/app.exe .
	xcopy public dist\public /E /Y

build_tailwind:
	./tailwind.exe -i ./tailwind.input.css -o ./public/css/tailwind.css