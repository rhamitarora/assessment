package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    swtTime "milestone2/swt"
)

func Test_timeHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:8080/time", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    timeHandler(res, req)

    exp := swtTime.CurrentTime()
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s gog %s", exp, act)
    }
}
