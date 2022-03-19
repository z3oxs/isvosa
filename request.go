package isvosa

import (
    "net/http"
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "io"
)

func request(Struct interface{}, token, endpoint string) {
    var reqError Error
    data, _ := json.Marshal(Struct)

    r, err := http.Post(fmt.Sprintf("%s/bot%s/%s", baseURL, token, endpoint), "application/json", bytes.NewBuffer(data))
    if err != nil {
        log.Fatal(err)

    }

    defer r.Body.Close()

    if r.StatusCode != 200 {
        b, _ := io.ReadAll(r.Body)
        json.Unmarshal(b, &reqError)

        log.Printf("%s [%d]", reqError.Description, reqError.ErrorCode)
    }
}
