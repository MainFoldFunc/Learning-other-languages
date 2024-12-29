package utils

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
)

func ParseBody(r *http.Request, x interface{}) {
    body, err := ioutil.ReadAll(r.Body)
    if err == nil {
        if err := json.Unmarshal([]byte(body), x); err != nil {
            return
        }
    }
}
