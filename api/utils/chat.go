package utils

import (
	"net/http"

	"github.com/gorilla/websocket"
	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
	"github.com/touch-some-grass-bro/vitamart/models"
)

func CreateRoom(hub *models.ChatHub, sellerEmail string, buyerEmail string, transcationID int) {
	room := &models.Room{
		ID: transcationID,
	}
	hub.Rooms[transcationID] = room
}

func NewHub() *models.ChatHub {
	return &models.ChatHub{
		Rooms:      map[int]*models.Room{},
		Register:   make(chan *models.ChatClient),
		Unregister: make(chan *models.ChatClient),
		Broadcast:  make(chan *models.Message, 5),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func JoinRoom(w http.ResponseWriter, r *http.Request, hub *models.ChatHub, queries *db.Queries, email, name string, isBuyer, isSeller bool, transactionID int) error {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return err
	}

	// Check if room exists
	if _, ok := hub.Rooms[transactionID]; !ok {
		CreateRoom(hub, "", "", transactionID)
	}

	user, err := queries.GetUser(r.Context(), email)

	if err != nil {
		return err
	}

	cl := &models.ChatClient{
		Conn:    conn,
		Email:   email,
		Name:    user.Name,
    RoomId: transactionID,
		Message: make(chan *models.Message, 10),
	}

	var messageContent string

	if isBuyer {
		cl.IsBuyer = true
		messageContent = "Buyer has joined the room"
	}

	if isSeller {
		cl.IsSeller = true
		messageContent = "Seller has joined the room"
	}

	message := &models.Message{
		Email:   "hub@vitamart.example.com",
		Name:    "Vitamart",
		RoomId:  transactionID,
		Content: messageContent,
	}

	hub.Register <- cl
	hub.Broadcast <- message

	go cl.WriteMessage()
	cl.ReadMessage(hub)

	return nil
}
