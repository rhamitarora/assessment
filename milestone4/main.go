// websockets.go
package main

import (    
    "net/http"
	
    "github.com/gorilla/websocket"
	t "milestone4/swt"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/timeupdating", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

        for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }
			
            if string(msg) == "true" {
				// Write message back to browser
				if err = conn.WriteMessage(msgType, []byte(t.CurrentTime())); err != nil {
					return
				}
            }
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "websockets.html")
    })

    http.ListenAndServe(":9802", nil)
}