# Option

A JSON and SQL-compatible Option type for Go.

Works with:

- `database/sql`
  - and any driver that supports `sql.Scan` and `driver.Value`
- `encoding/json`
  - and any other JSON library that supports `Marshal` and `Unmarshal`
- `go-playground/validator`

## Installation

```bash
go get -u github.com/EwenQuim/option
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/EwenQuim/option"
)

func main() {
	// Option is set
	opt := option.Some("Hello")
	fmt.Println(opt) // Some(Hello)

	value, ok := opt.GetValue()
	fmt.Println(value, ok) // Hello true

	// Option not set
	opt2 := option.None[string]()
	fmt.Println(opt2) // None

	value, ok = opt2.GetValue()
	fmt.Println(value, ok) // <nil> false
}
```

### With encoding/json

... to be continued ...

### With database/sql

... to be continued ...

The idea is based on the unmaintained project mo, so thanks to samber for the idea and some parts of the code!
