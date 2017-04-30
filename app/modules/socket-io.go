package modules

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

const SOCKET_EVENT_CONNECTION = "connection"
const SOCKET_EVENT_DISCONNECTION = "disconnection"
const SOCKET_EVENT_TASK = "task"
const SOCKET_EVENT_COMPLETED = "completed"

const CLIENT_STATE_FREE = "FREE"
const CLIENT_STATE_WORK = "WORK"

type SocketClient struct {
	socket socketio.Socket
	state string
}

func (this *SocketClient) IsFree () bool {
	return this.state == CLIENT_STATE_FREE
}

func (this *SocketClient) SendTask (task *Task) {
	this.state = CLIENT_STATE_WORK
	this.socket.Emit(SOCKET_EVENT_TASK, task.Data)
}

func StartServer() {
	startServer()
}

func GetSocketClients() []*SocketClient {
	return clients
}


// --- PRIVATE ---------------------------------------------------------------------------------------------------------
var clients = make([]*SocketClient, 0)

func addClient(socket socketio.Socket) {
	client := &SocketClient{
		socket: socket,
		state: CLIENT_STATE_FREE,
	}

	socket.On(SOCKET_EVENT_COMPLETED, func(data string) {
		client.state = CLIENT_STATE_FREE
		ThatClientCompletedTask(client, data)
	})

	socket.On(SOCKET_EVENT_DISCONNECTION, func() {
		log.Println("Disconnect Slave")

		// Remove item in array
		clientIndex := -1
		for index, clientInSlice := range clients {
			if (client == clientInSlice) {
				clientIndex = index
				break
			}
		}

		if (-1 < clientIndex) {
			clients = append(clients[:clientIndex], clients[clientIndex + 1:]...)
		}
	})

	clients = append(clients, client)
}

func startServer() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	server.On(SOCKET_EVENT_CONNECTION, func(socket socketio.Socket) {
		log.Println("Connection Slave")

		addClient(socket)
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:15666...")
	log.Fatal(http.ListenAndServe(":15666", nil))
}
