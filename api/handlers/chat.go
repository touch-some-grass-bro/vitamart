package handlers

import (
	"net/http"
	"strconv"

	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
	"github.com/touch-some-grass-bro/vitamart/utils"
	"github.com/touch-some-grass-bro/vitamart/models"
)

func JoinRoomHandler(queries *db.Queries, hub *models.ChatHub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{} = make(map[string]interface{})

		jwtTokenCookie, err := r.Cookie("token")

		if err != nil {
			resp["error"] = "No token found"
			utils.JSON(w, http.StatusUnauthorized, resp)
			return
		}

    isSellerQ := r.URL.Query().Get("isSeller")
    isBuyerQ := r.URL.Query().Get("isBuyer")

    if isSellerQ == "" && isBuyerQ == "" {
      resp["error"] = "isSeller or isBuyer must be true"
      utils.JSON(w, http.StatusBadRequest, resp)
      return
    }

    if isSellerQ == "true" && isBuyerQ == "true" {
      resp["error"] = "both isSeller and isBuyer cannot be true"
      utils.JSON(w, http.StatusBadRequest, resp)
      return
    }


    isSeller := isSellerQ == "true"
    isBuyer := isBuyerQ == "true"

    transactionID, err := strconv.Atoi(r.URL.Query().Get("transactionID"))

    if err != nil {
      resp["error"] = "transactionID must be an integer"
      utils.JSON(w, http.StatusBadRequest, resp)
      return
    }


		email, errMsg, status := utils.GetUserEmailFromJWT(jwtTokenCookie.Value)

		if errMsg != "" {
			resp["error"] = errMsg
			utils.JSON(w, status, resp)
			return
		}

    user, err := queries.GetUser(r.Context(), email)

		err = utils.JoinRoom(w, r, hub, queries, user.Email , user.Name, isBuyer, isSeller, transactionID)

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}
	}
}
