# Tiny Go Rate Limiter 

Tiny rate limiter 

```go
import (
	"fmt"
	"time"
	"github.com/xsurge83/tinyratelimiter"
)

func main() {

    ratelimiter := tinyratelimiter.NewRateLimiter(2, time.Second)

	for i := 0; i < 2; i++ {
        fmt.Println(i, ratelimiter.Allow())
	}
	mt.Println(ratelimiter.Allow())
    
    // Output:
    // 0 true
    // 1 true
    // false

}
```
