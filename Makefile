build_project_windows: build_tailwind
	go build -o ./dist/app_windows.exe .
	xcopy public dist\public /E /Y

build_project_linux: build_tailwind
	go build -o ./dist/app_linux .
	cp -rf ./public ./dist/public

build_tailwind:
	./tailwind.exe -i ./tailwind.input.css -o ./public/css/tailwind.css