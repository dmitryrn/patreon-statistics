package ws_server

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/websocket"
//	"github.com/pkg/errors"
//	"github.com/reactivex/rxgo/v2"
//	"log"
//	"net/http"
//	"patreon-statistics/internal/events"
//	"sync"
//)
//
//type IWSServer interface {
//}
//
//type WSServer struct {
//	ch       chan rxgo.Item
//	ob       rxgo.Observable
//	clients  map[int]*WSClient
//	clientsX sync.Mutex
//}
//
//var upgrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//type wsMessage struct {
//	clientId int
//	id       string
//	data     interface{}
//}
//
//type clientWsMessage struct {
//	Id   string      `json:"id"`
//	Data interface{} `json:"data"`
//}
//
//func parseMessage(messageBytes []byte) (*wsMessage, error) {
//	clientMessage := clientWsMessage{}
//
//	err := json.Unmarshal(messageBytes, &clientMessage)
//	if err != nil {
//		return nil, err
//	}
//	log.Println("parsed client message", clientMessage)
//
//	return &wsMessage{
//		id:   clientMessage.Id,
//		data: clientMessage.Data,
//	}, nil
//}
//
//func (ws WSServer) handler(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Print("ws: failed to upgrade", err)
//		return
//	}
//
//	ws.clientsX.Lock()
//	defer ws.clientsX.Unlock()
//
//	var clientId int
//	i := 0
//	for {
//		if _, ok := ws.clients[i]; !ok {
//			clientId = i
//			break
//		}
//		i++
//	}
//
//	client := NewWSClient(
//		conn,
//		func() { ws.onCloseConnection(clientId) },
//		func(message *wsMessage) { ws.onMessageFromClient(clientId, message) },
//	)
//	client.Listen()
//
//	ws.clients[clientId] = client
//}
//
//func (ws WSServer) onMessageFromClient(clientId int, message *wsMessage) {
//	message.clientId = clientId
//
//	appItem := events.Message{
//		Id:   events.MID_incoming_ws_message,
//		Data: message,
//	}
//
//	ws.ch <- rxgo.Item{V: appItem}
//}
//
//func (ws WSServer) onCloseConnection(clientId int) {
//	ws.clientsX.Lock()
//	defer ws.clientsX.Unlock()
//
//	// WSClient should close ws connection itself, here we only delete it from map
//	delete(ws.clients, clientId)
//	log.Println("ws: connection closed, clientId", clientId)
//}
//
//func (ws *WSServer) sendMessage(message events.Data_outgoing_ws_message) error {
//	ws.clientsX.Lock()
//	defer ws.clientsX.Unlock()
//
//	client, ok := ws.clients[message.WSClientId]
//	if !ok {
//		return errors.New(fmt.Sprint("no such client, id:", message.WSClientId))
//	}
//
//	err := client.Send(message)
//	if err != nil {
//		return errors.Wrap(err, fmt.Sprint("fail wsClient.Send, message:", message))
//	}
//
//	return nil
//}
//
//func NewWSServer(ch chan rxgo.Item, ob rxgo.Observable) error {
//	ws := WSServer{
//		ch: ch,
//		ob: ob,
//		clients: make(map[int]*WSClient),
//	}
//
//	ob.ForEach(func(item interface{}) {
//		message, ok := item.(events.Data_outgoing_ws_message)
//		if !ok {
//			return
//		}
//		err := ws.sendMessage(message)
//		if err != nil {
//			log.Println(errors.Wrap(
//				err,
//				fmt.Sprintf("ws: fail send message. clientId: %v, data: %v", message.WSClientId, message.Data),
//			))
//		}
//	}, nil, nil)
//
//	http.HandleFunc("/echo", ws.handler)
//
//	log.Println("ws: starting ws server")
//
//	err := http.ListenAndServe("localhost:8080", nil)
//	return errors.Wrap(err, "fail to start ws server")
//}
