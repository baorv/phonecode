# Phone code reader for 

[![Build Status](https://travis-ci.org/baorv/phonecode.svg?branch=master)](https://travis-ci.org/baorv/phonecode)

This library helps to lookup calling code of all countries in the world

## Installation

* Go get way

```bash
go get github.com/baorv/phonecode
```

* [Glide](https://github.com/Masterminds/glide) way

```bash
glide get github.com/baorv/phonecode
```

## Usage

```go
package main

import (
	"github.com/baorv/phonecode"
	"log"
)

func main() {
    // You can use custom data
	// reader, err := phonecode.New("/path/to/customize.csv")
	reader, err := phonecode.New()
	if err != nil {
		log.Panicf("Error when open reader with message: %s", err.Error())
	}
	record := reader.LookupPhoneCode("VN")
	log.Printf("%s as calling code is: %s", record.CountryCode, record.Code)
}
```

## Testing

Execute test suite:

```bash
go test
```

## Contributing

Contributions welcome! Please fork the repository and open a pull request with your changes.

## License

This project is licensed under the [MIT License](LICENSE).