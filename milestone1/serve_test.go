package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_helloHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:8596/", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    helloHandler(res, req)

    exp := "Hello World!"
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s gog %s", exp, act)
    }
}