package PKG

import (
    "encoding/json"
    "io"
    "net/http"
)

// ParseBody reads the body from an HTTP request and unmarshals it into the provided interface
func ParseBody(r *http.Request, x interface{}) {
    body, err := io.ReadAll(r.Body) // Using io.ReadAll instead of ioutil.ReadAll (deprecated in Go 1.16)
    if err != nil {
        return
    }
    if err := json.Unmarshal(body, x); err != nil {
        return
    }
}

