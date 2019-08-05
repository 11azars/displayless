install_tooling:
	go get github.com/rakyll/statik

static_assets:
	cd ui; npm run build
	statik -src=ui/dist -p=ui

clean:
	-rm ui/statik.go
	-rm -r bin

build: clean static_assets
	mkdir -p bin
	go build -o bin/displayless cmd/main.go
