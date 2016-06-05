package websocketserver

import "fmt"
import "net/http"
import "golang.org/x/net/websocket"

type websocketDataTransformer func ([]byte) []byte

func NewHandler (route string, handler websocketDataTransformer) {
	server := func (ws *websocket.Conn) {
	    msg := make([]byte, 24)
	    for {
	    	fmt.Println("waiting for message from websocket")
	    	err := websocket.Message.Receive(ws, &msg)
		    if err != nil {
		    	if err.Error() == "EOF" {
	    			fmt.Println("Socket input ended. Closing.")
		    		ws.Close()
		    		break
		    	}
		        panic("Reading from Websocket Failed: " + err.Error())
		    }
		    text := string(msg)

		    fmt.Println("Received bytes: ", msg)
		    fmt.Println("Received message: ", text)

		    response := handler(msg)

		    ws.Write(response)
	    }
	}

	fmt.Println("Listening on route: " + route)
    http.Handle("/" + route, websocket.Handler(server))    
	return
}
