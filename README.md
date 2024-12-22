# goenv

goenv is a simple Go module for work with environment variables throughout
your application through a functional interface.

## Installation

To install goenv, use `go get`:

```sh
go get github.com/sepdotgg/goenv
```

## Usage

Import the goenv package in your Go code:

```go
import "github.com/yourusername/goenv"

env := goenv.NewDefaultEnvironment()
hostname, err := env.Get("DB_HOSTNAME")

defaultPort := env.GetOrDefault("DB_PORT", "3000")

dbUser := env.MustGet("DB_USER") // panic if not set

```

## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.
