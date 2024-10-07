# Go Advanced Admin - Echo Integration

[Echo](https://echo.labstack.com/) framework integration for the Go Advanced Admin Panel.

[![Go Report Card](https://goreportcard.com/badge/github.com/go-advanced-admin/web-echo)](https://goreportcard.com/report/github.com/go-advanced-admin/web-echo)
[![Go](https://github.com/go-advanced-admin/web-echo/actions/workflows/tests.yml/badge.svg)](https://github.com/go-advanced-admin/web-echo/actions/workflows/tests.yml)
[![License: Apache-2.0](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/go-advanced-admin/web-echo?tab=doc)

This package provides integration with the Echo web framework for the Go Advanced Admin Panel, enabling you to use Echo as your web framework.

## Installation

Add the module to your project by running:

```sh
go get github.com/go-advanced-admin/echo-integration
```

## Documentation

For detailed documentation on how to use the Echo integration, please visit the [official documentation website](https://goadmin.dev/integrations/echo).

## Quick Start

```go
import (
    "github.com/go-advanced-admin/admin"
    "github.com/go-advanced-admin/echo-integration"
    "github.com/labstack/echo/v4"
)

func main() {
    // Initialize Echo
    e := echo.New()

    // Initialize the web integrator
    webIntegrator := adminecho.NewIntegrator(e.Group("/admin"))

    // Use webIntegrator when initializing the admin panel
}
```

For more detailed examples and configuration options, please refer to the [Echo Integration Guide](https://goadmin.dev/integrations/echo).

## Contributing

Contributions are always welcome! Please refer to the [Contributing Guidelines](https://github.com/go-advanced-admin/admin/blob/main/CONTRIBUTING.md) in the main repository.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
