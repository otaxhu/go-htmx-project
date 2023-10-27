build_project_windows: build_templ build_tailwind
	go build -o ./dist/app_windows.exe .

build_project_linux: build_templ build_tailwind
	go build -o ./dist/app_linux .

build_templ:
	templ generate -path ./internal/web/templates

build_tailwind:
	.\tailwind.exe -i ./tailwind.input.css -o ./static/css/tailwind.css

# This commands must be executed only once
install_external_tools_windows: \
	install_tailwind_cli_windows \
	install_templ_cli

install_tailwind_cli_windows:
	curl -sL -o ".\tailwind.exe" "https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.5/tailwindcss-windows-x64.exe"

install_templ_cli:
	go install github.com/a-h/templ/cmd/templ@latest