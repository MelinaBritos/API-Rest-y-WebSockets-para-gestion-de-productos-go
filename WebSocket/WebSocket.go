package WebSocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func Init() {
	go HandleMessages()
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	Conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error de actualización de WebSocket: %v", err)
		return
	}
	defer Conn.Close()

	clients[Conn] = true

	for {
		_, msg, err := Conn.ReadMessage()
		if err != nil {
			log.Printf("Cliente desconectado: %v", err)
			delete(clients, Conn)
			break
		}

		Conn.WriteMessage(websocket.TextMessage, []byte("Servidor recibió: "+string(msg)))
	}

}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("Error al enviar mensaje: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func Emit(eventJSON []byte) {
	broadcast <- eventJSON
}
