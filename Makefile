default: clean bindir
	go build -o bin/scrape *.go

amd64: clean bindir linux windows darwin compress

bindir:
	mkdir bin

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scrape.linux-amd64 *.go

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scrape.windows-amd64.exe *.go

darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/scrape.darwin-amd-64 *.go

compress:
	cd ./bin && find . -name 'scrape*' | xargs -I{} tar czf {}.tar.gz {}

clean:
	rm -rf bin

snap-clean:
	rm -f scrape_*_amd64.snap*
	snapcraft clean

snap-build:
	snapcraft

snap-publish:
	snapcraft push --release=edge scrape_*_amd64.snap
