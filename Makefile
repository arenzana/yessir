version=`cat VERSION`
all: build
build:
	go build -ldflags "-X 'github.com/arenzana/yessir/cmd.ApplicationVersion=$(version)'"
install:
	go install -ldflags "-X 'github.com/arenzana/yessir/cmd.ApplicationVersion=$(version)'"
buildall:
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-X 'github.com/arenzana/yessir/cmd.ApplicationVersion=$(version)'" -o yessir_darwin_amd64
	env GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/arenzana/yessir/cmd.ApplicationVersion=$(version)'" -o yessir_linux_amd64
docker:
	docker build -t arenzana/yessir:$(version) -t arenzana/yessir:latest .
clean:
	rm yessir
