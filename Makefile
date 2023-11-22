build_project_linux:
	./bin/cssmodules \
		-input-dir ./internal/web/views/ \
		-output-dir ./compiled-templates \
		-output-css-path ./static/css/styles.css

	go build -o ./dist/app_linux .

# This commands must be executed only once
install_external_tools_linux: \
	install_bootstrap_icons \
	build_cssmodules

install_bootstrap_icons:
	curl -sL --create-dirs -o "./static/icons/bs.zip" "https://github.com/twbs/icons/releases/download/v1.11.1/bootstrap-icons-1.11.1.zip"
	unzip -q -j -d "./static/icons/bootstrap-icons" "./static/icons/bs.zip" "*.svg"
	rm -f "./static/icons/bs.zip"

build_cssmodules:
	go build -o ./bin/cssmodules ./cmd/cssmodules