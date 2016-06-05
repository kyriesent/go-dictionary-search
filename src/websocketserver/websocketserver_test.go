package websocketserver

import "testing"
import "fmt"
import "time"
import "net/http"
import "golang.org/x/net/websocket"

func TestWebsocketServer (t *testing.T) {
	NewHandler("echo", dataTransformer)

    go func () {
		fmt.Println("Listening on port " + "8080")
    	err := http.ListenAndServe(":" + "8080", nil)
	    if err != nil {
	        panic("ListenAndServe failed: " + err.Error())
	    }
    } ()

	fmt.Println("Running test")

	time.Sleep(1000 * time.Microsecond)

	origin := "http://localhost/"
	url := "ws://localhost:8080/echo"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
	    t.Error(err.Error())
	}
	
	var msg = make([]byte, 512)
	var result string
	var n int

	if _, err = ws.Write([]byte("hello 1234")); err != nil {
	    t.Error(err.Error())
	}

	if n, err = ws.Read(msg); err != nil {
	    t.Error(err.Error())
	}

	result = string(msg[:n])

	if result != "hello 1234" {
	    t.Error("Expected result to be \"hello 1234\", got " + result)
	}



	if _, err = ws.Write([]byte("world")); err != nil {
	    t.Error(err.Error())
	}


	if n, err = ws.Read(msg); err != nil {
	    t.Error(err.Error())
	}

	result = string(msg[:n])

	if result != "world" {
	    t.Error("Expected result to be \"world\", got " + result)
	}

	return
}

func dataTransformer (in []byte) []byte {
	return in
}