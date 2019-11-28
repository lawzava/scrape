build: ensure-dir build-linux build-windows build-darwin compress

ensure-dir:
	rm -rf bin
	mkdir bin

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scrape.linux-amd64 *.go

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scrape.windows-amd64.exe *.go

build-darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scrape.darwin-amd-64 *.go

compress:
	cd ./bin && find . -name 'scrape*' | xargs -I{} tar czf {}.tar.gz {}

snap-build:
	rm -f scrape_*_amd64.snap
#	snapcraft clean scrape -s build
	snapcraft clean
	snapcraft

snap-publish:
	snapcraft push --release=stable scrape_*_amd64.snap
