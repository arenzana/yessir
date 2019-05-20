version=`cat VERSION`
all: build
build:
	go build -ldflags "-X 'github.com/iarenzana/yessir/cmd.ApplicationVersion=$(version)'"
install:
	go install -ldflags "-X 'github.com/iarenzana/yessir/cmd.ApplicationVersion=$(version)'"
buildall:
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-X 'github.com/iarenzana/yessir/cmd.ApplicationVersion=$(version)'" -o yessir_darwin_amd64
	env GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/iarenzana/yessir/cmd.ApplicationVersion=$(version)'" -o yessir_linux_amd64
docker:
	docker build -t github.com/iarenzana/yessir:$(version) -t github.com/iarenzana/yessir:latest .
clean:
	rm yessir
