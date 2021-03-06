build:
	GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -o bin/cloudadaptor-x86.linux
	GOOS=darwin GOARCH=amd64 go build -o bin/cloudadaptor-x86.darwin
release: build
	ossutil cp -r -u bin/ oss://grstatic/binary