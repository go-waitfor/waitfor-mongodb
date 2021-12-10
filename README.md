# waitfor-mongodb
MongoDB resource readiness assertion library

# Quick start

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-mongodb"
	"os"
)

func main() {
	runner := waitfor.New(mongodb.Use())

	err := runner.Test(
		context.Background(),
		[]string{"mongodb://user:passsword@my-server.net/my-db"},
		waitfor.WithAttempts(5),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```
