# web
Simple HTTPS web server with HTTP redirection and without running as root

```go
package main

import (
  "github.com/jdcio/web"
  "net/http"
)

func main() {
  web.Boot(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello World!"))
  }),
    "./keys/server.*",
    "web", // run as user web after opening ports
  )
}
```
