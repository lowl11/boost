<pre>
| |__   ___   ___  ___| |_
| '_ \ / _ \ / _ \/ __| __|
| |_) | (_) | (_) \__ \ |_
|_.__/ \___/ \___/|___/\__|
</pre>
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

### Features
<table>
    <thead>
        <th>Feature</th>
        <th>Description</th>
        <th>Status</th>
    </thead>
    <tbody>
        <tr>
            <td>Logging</td>
            <td>Static call of log functions from anywhere</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Config</td>
            <td>Get environment & .yml config file from anywhere</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Errors</td>
            <td>Custom errors with rich context</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Easy controllers</td>
            <td>Base controller to easy returning responses</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>HTTP Requests</td>
            <td>Sending HTTP requests client with retries</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Request body validator</td>
            <td>Validating incoming request body (JSON) based on go-playground validator</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Health checker</td>
            <td>Health checking service (self-app + any other)</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Destroyer</td>
            <td>Close connections or destroy something at the end of application</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Cache</td>
            <td>Memory & Redis cache client</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>RabbitMQ / MessageBus</td>
            <td>Message bus pattern + RMQ support</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>DI</td>
            <td>Dependency Injection</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Cron</td>
            <td>Cron Job actions support</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>gRPC</td>
            <td>gRPC server & client support</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Database (external)</td>
            <td>Query builder & ORM repositories</td>
            <td>:white_check_mark:</td>
        </tr>
    </tbody>
</table>

### TODO

- Swagger support
- gateway (proxy router)
- Rate limiter
- DB ORM
