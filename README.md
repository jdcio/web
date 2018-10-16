# web
Simple HTTPS web server with HTTP redirection and without running as root

```go
package main

import (
  "github.com/jdcio/web"
  "net/http"
)

func main() {
  web.Load("./keys/server.*", // ssl certs blob
    "web", // After server boots as root to open ports 80 and 443 it will drop down to user 'web'
  )
  
  web.Serve(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello World!"))
  })
}
```
