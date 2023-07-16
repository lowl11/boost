# boost
Minimalist Go framework based on FastHTTP

### Boosted Get Started
```go
import (
    "github.com/lowl11/boost"
    "net/http"
)

func main() {
    app := boost.New()
    
    app.GET("/ping", func(ctx boost.Context) error {
        return ctx.
            Status(http.StatusOK).
            String("pong")
    })
    
    app.Run(":8080")	
}
```

### External TODO

- gRPC support
- cron support
- RMQ support
- request validation support
- DB support
