// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type GoogleToken struct {
	ID           int64     `json:"id"`
	UserEmail    string    `json:"userEmail"`
	CreatedAt    time.Time `json:"createdAt"`
	RefreshToken string    `json:"refreshToken"`
	AccessToken  string    `json:"accessToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	TokenType    string    `json:"tokenType"`
}

type Item struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageBinary []byte    `json:"imageBinary"`
	Price       int32     `json:"price"`
	SellerEmail string    `json:"sellerEmail"`
	Issold      bool      `json:"issold"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Transaction struct {
	ID         int64  `json:"id"`
	ItemID     int64  `json:"itemID"`
	BuyerEmail string `json:"buyerEmail"`
}

type User struct {
	Email             string         `json:"email"`
	Name              string         `json:"name"`
	JoinYear          int32          `json:"joinYear"`
	ProfilePictureUrl string         `json:"profilePictureUrl"`
	Hostel            sql.NullString `json:"hostel"`
}
