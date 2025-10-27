package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "io/ioutil"
)

func TestGetBooks(t *testing.T) {
    req := httptest.NewRequest("GET", "/books", nil)
    w := httptest.NewRecorder()
    getBooks(w, req)
    res := w.Result()
    if res.StatusCode != 200 {
        t.Fatalf("expected status 200 got %d", res.StatusCode)
    }
    body, _ := ioutil.ReadAll(res.Body)
    var arr []map[string]interface{}
    json.Unmarshal(body, &arr)
    if len(arr) == 0 {
        t.Fatalf("expected at least one book in response")
    }
}
