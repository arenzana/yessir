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

## Usage

Below is the usage of the `run` command.

```bash
Start serving!

Usage:
  yessir run [flags]

Flags:
  -c, --cert string     Path to the server TLS certificate file (only for https)
  -h, --help            help for run
  -k, --key string      Path to the server TLS certificate key file (only for https)
  -p, --port int        Port to listen on (default 8888)
  -r, --return int      HTTP return code (200,404,500) (default 200)
  -s, --scheme string   Scheme http|https (default "http")
  ```
  
Noteworthy options are `-r` to return a different http code rather than 200. And `-s` to pick https (which works inconjunction with `-k` and `-c` for TLS) as the scheme.

### Examples

```bash
yessir run
```

Is the simplest way to run `yessir`. It will start a server on port `8888`  that will listen for all requests.

```bash
yessir -p 8877 -s https -c ~/Downloads/example.com.crt -k ~/Downloads/example.com.key -r 500
```

The example above will listen for `https` requests on port `8877` and, instead of returning `200`, it will always return a `500` error.
