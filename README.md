# ErrX - Enhanced Error Handling for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/hamedcodelab/errx)](https://goreportcard.com/report/github.com/hamedcodelab/errx)
[![GoDoc](https://pkg.go.dev/badge/github.com/hamedcodelab/errx)](https://pkg.go.dev/github.com/hamedcodelab/errx)

## Overview
ErrX is a structured and extensible error handling package for Go, designed to provide standardized error messages, HTTP status codes, and improved stack tracing.

## Features
- Standardized error types
- Support for HTTP and gRPC error mapping
- Enhanced stack trace handling
- Easy integration into Go projects

## Installation
To install ErrX, run:

```sh
go get github.com/hamedcodelab/errx
```

## Usage
### Creating and Handling Errors
```go
package main

import (
	"fmt"
	"github.com/hamedcodelab/errx"
)

func main() {
	err := errx.New(errors.New("error validation").WithType(errx.ErrorValidation).WithCode(http.StatusBadRequest)
	fmt.Println(err.Error())
}
```

## Contributing
We welcome contributions! To get started:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to your fork (`git push origin feature-branch`)
5. Open a pull request

### Guidelines
- Follow Go best practices
- Write unit tests for new features
- Keep documentation up-to-date

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
For questions or support, open an issue or reach out to [@hamedcodelab](https://github.com/hamedcodelab).

