# middleware

middleware usage

## authuser

see [README](authuser/README.md)

```
package server

import (
    "github.com/gin-gonic/gin"
    "git.llsapp.com/zhenghe/pkg/http/middleware/authuser"
)

func main() {
    router := gin.New()
    router.Use(m.Auth())
    
    router.GET("/public", auther.IgnoreAuth(), func(c *gin.Context) {
        c.String(http.StatusOK, "OK")
    })
}
```


## CORS

```
package server

import (
    "github.com/gin-gonic/gin"
    m "git.llsapp.com/zhenghe/pkg/http/middleware"
)

func main() {
    router := gin.New()
    router.Use(m.CORS())
}
```

## newrelic

```
import (
    "github.com/gin-gonic/gin"
    m "git.llsapp.com/zhenghe/pkg/http/middleware"
)

func main() {
    router := gin.New()
    router.Use(gin.Logger(), m.WithNewrelic(app), m.ErrorHandler(), m.Recovery())
}
```

## timeoffset

server:

```
package server

import (
    "github.com/gin-gonic/gin"
    m "git.llsapp.com/zhenghe/pkg/http/middleware"
)

func main() {
    router := gin.New()
    router.Use(m.TimeOffset())
}
```

handler:

```
package handler

import (
    "github.com/gin-gonic/gin"
    m "git.llsapp.com/zhenghe/pkg/http/middleware"
)

func Hello(c *gin.context) {
    timeOffset := m.GetTimeOffset(c)
}
```

## ratelimit

see [README](ratelimit/README.md)

```
package server

import (
    "time"

    "github.com/go-redis/redis"
    "github.com/gin-gonic/gin"
    "git.llsapp.com/zhenghe/pkg/http/middleware/ratelimit"
    m "git.llsapp.com/zhenghe/pkg/http/middleware"
)

func main() {
    router := gin.New()
    router.Use(m.Auth())
    
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    helloWorld := func(c *gin.Context) {
        c.String(http.StatusOK, "OK")
    }

    router.GET("/public", ratelimit.RateLimit(client, ratelimit.SetRate(5, time.Minute)), helloWorld)
}
```