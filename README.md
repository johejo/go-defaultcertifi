# go-defaultcertifi

[![ci](https://github.com/johejo/go-defaultcertifi/workflows/ci/badge.svg?branch=main)](https://github.com/johejo/go-defaultcertifi/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/johejo/go-defaultcertifi/branch/main/graph/badge.svg)](https://codecov.io/gh/johejo/go-defaultcertifi)
[![Go Report Card](https://goreportcard.com/badge/github.com/johejo/go-defaultcertifi)](https://goreportcard.com/report/github.com/johejo/go-defaultcertifi)

## Description

Automatic certificate bundler for `http.DefaultTransport`.

Intercepts the TLS config of `http.DefaultTransport` so helps you make HTTPS requests in any environment with a single binary without installing ca-certificates etc.

You don't have to add certificates into image in Dockerfile such as `apk add --no-cache ca-certificates` for making HTTPS Requests.

## Usage

```go
import (
  _ "github.com/johejo/go-defaultcertifi/http"
)
```

## License

MIT

## Author

@johejo
