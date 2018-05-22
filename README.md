# Utils For Go


## Strings


### Substring


```go
func Substring(s string, start int64, end int64) string
```

Get the substring of a string from the index of start to end



Usage

```
package main

import (
    "fmt"

    "github.com/mrasong/util"
)

func main() {
    s := "我是Gopher"
    fmt.Println(util.Substring(s, 0, 2))
}
```