build: ensure-dir build-linux build-windows build-darwin

ensure-dir:
	mkdir bin

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scraper_linux *.go

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scraper_windows *.go

build-darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scraper_mac *.go
