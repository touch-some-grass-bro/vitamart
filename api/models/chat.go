package models

import (
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type ChatClient struct {
	Conn     *websocket.Conn
	Email    string
	Name     string
	IsSeller bool
	IsBuyer  bool
	Message  chan *Message
	RoomId   int
}

type Message struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	RoomId  int    `json:"room_id"`
	Content string `json:"content"`
}

type Room struct {
	ID           int `json:"id"`
	BuyerClient  *ChatClient
	SellerClient *ChatClient
}

type ChatHub struct {
	Rooms      map[int]*Room
	Register   chan *ChatClient
	Unregister chan *ChatClient
	Broadcast  chan *Message
}

func (h *ChatHub) Run() {
	for {
		select {
		case client := <-h.Register:
			room := h.Rooms[client.RoomId]
			if room == nil {
				continue
			}
			if client.IsBuyer {
				room.BuyerClient = client
			}
			if client.IsSeller {
				room.SellerClient = client
			}
			if room.BuyerClient != nil && room.SellerClient != nil {
				go h.BroadcastMessage(room, &Message{
					Email:   "hub@vitamart.example.com",
					Name:    "Vitamart",
					RoomId:  room.ID,
					Content: "Seller and Buyer are both online!",
				})
			}
		case client := <-h.Unregister:
			room := h.Rooms[client.RoomId]
			if room == nil {
				continue
			}
			if client.IsBuyer {
				room.BuyerClient = nil
			}
			if client.IsSeller {
				room.SellerClient = nil
			}
			if room.BuyerClient == nil && room.SellerClient == nil {
				delete(h.Rooms, room.ID)
			}

			msg := &Message{
				Email:  "hub@vitamart.example.com",
				Name:   "Vitamart",
				RoomId: room.ID,
			}

			if client.IsBuyer {
				msg.Content = "Buyer has left the room"
			}

			if client.IsSeller {
				msg.Content = "Seller has left the room"
			}

			go h.BroadcastMessage(room, msg)

		case message := <-h.Broadcast:
			room := h.Rooms[message.RoomId]
			if room == nil {
				continue
			}
			go h.BroadcastMessage(room, message)
		}
	}
}

func (h *ChatHub) BroadcastMessage(room *Room, message *Message) {
	if room.BuyerClient != nil && room.BuyerClient.Email != message.Email {
		room.BuyerClient.Message <- message
	}
	if room.SellerClient != nil && room.SellerClient.Email != message.Email {
		room.SellerClient.Message <- message
	}
}

func (c *ChatClient) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *ChatClient) ReadMessage(hub *ChatHub) {
	defer func() {
		if err := recover(); err != nil {
			log.Info().Err(err.(error)).Msg("Recovered from panic")
		}
	}()

	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Info().Msg(err.Error())
			}
		}

		msg := &Message{
			Email:   c.Email,
			Name:    c.Name,
			RoomId:  c.RoomId,
			Content: string(m),
		}

		hub.Broadcast <- msg
	}
}
