[![Go Report Card](https://goreportcard.com/badge/github.com/orasik/gocomparejson)](https://goreportcard.com/report/github.com/orasik/gocomparejson)
[![Build Status](https://travis-ci.org/orasik/gocomparejson.svg?branch=master)](https://travis-ci.org/orasik/gocomparejson)

# ComapreJSON

Compare two json strings in go lang


### Installation

```bash
go get github.com/orasik/gocomparejson
```

### Usage

#### Example

```go
package main

import (
    "fmt"
    "github.com/orasik/gojsoncompare"
)

json1 := []byte(`{"hello":"world"}`)
json2 := []byte(`{
"hello":"world"
}`)

val, err := CompareJSON(string(json1), string(json2))

fmt.Print(val)
fmt.Print(err)
```

### Unit test

```bash
go test -v
```
