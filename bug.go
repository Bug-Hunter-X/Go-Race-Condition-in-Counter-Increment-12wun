```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var counter int
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        counter++
                }()
        }
        wg.Wait()
        fmt.Println("Counter:", counter)
}
```