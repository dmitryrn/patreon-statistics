package ws_server

//
//import (
//	"github.com/gorilla/websocket"
//	"github.com/pkg/errors"
//	"log"
//	"patreon-statistics/internal/events"
//)
//
//type WSClient struct {
//	connection *websocket.Conn
//	onClose func()
//	onMessage func(*wsMessage)
//}
//
//func (c WSClient) Send(message events.Data_outgoing_ws_message) error {
//	return c.connection.WriteJSON(message)
//}
//
//func (c *WSClient) Listen()  {
//	log.Println("client connected")
//
//	for {
//		mt, messageString, err := c.connection.ReadMessage()
//		if mt != websocket.TextMessage {
//			println("ws: closed as not a text message, mt:", mt)
//			break
//		}
//
//		if err != nil {
//			log.Println("ws: fail read message:", err)
//			break
//		}
//		log.Printf("ws: received: %s", messageString)
//
//		wsMessage, err := parseMessage(messageString)
//		if err != nil {
//			log.Println(errors.Wrap(err, "ws: fail parse message"))
//			continue
//		}
//
//		log.Println("parsed ws message", wsMessage)
//
//		c.onMessage(wsMessage)
//	}
//
//	c.connection.Close()
//	c.onClose()
//}
//
//func NewWSClient(conn *websocket.Conn, onClose func(), onMessage func(*wsMessage)) *WSClient {
//	client := &WSClient{
//		connection: conn,
//		onClose: onClose,
//		onMessage: onMessage,
//	}
//
//	return client
//}
