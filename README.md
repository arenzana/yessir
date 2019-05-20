[![Build Status](https://travis-ci.org/iarenzana/urbanobot.png)](https://travis-ci.org/iarenzana/yessir)
# yessir
`yessir` creates a mock API server that will accept anything you throw at it (or reject, see below) for testing purposes.

## Install
### Build from source

Make sure Go 1.11+ is installed on your machine. You can follow [this guide](https://golang.org/doc/install) to do so. On a Mac, just set your $GOPATH and run `brew install go`.

```bash
git clone https://gitlab.com/iarenzana/yessir.git
cd yessir
make
```

Now you can execute `yessir`

### Docker

```bash
make docker
docker run --rm -it -p 8888:8888 github.com/iarenzana/yessir:latest
```
