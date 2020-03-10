# pconf

Разбор конфигурационного файла

## Example
```go
package main

import (
        "fmt"
        "github.com/uzhinskiy/lib.go/pconf"
)
var Config     pconf.ConfigType
func main() {
        Config = make(pconf.ConfigType)
	err := Config.Parse("./example/main.cfg")
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(Config)
}

```
