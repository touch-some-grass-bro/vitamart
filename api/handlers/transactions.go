package handlers

import (
	"net/http"
	"strconv"

	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
	"github.com/touch-some-grass-bro/vitamart/utils"
)

func BuyItemHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{} = make(map[string]interface{})

		jwtTokenCookie, err := r.Cookie("token")

		if err != nil {
			resp["error"] = "No token found"
			utils.JSON(w, http.StatusUnauthorized, resp)
			return
		}

		email, errMsg, status := utils.GetUserEmailFromJWT(jwtTokenCookie.Value)
		if errMsg != "" {
			resp["error"] = errMsg
			utils.JSON(w, status, resp)
			return
		}

		itemId, err := strconv.Atoi(r.URL.Query().Get("item_id"))

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}

		item, err := queries.GetItem(r.Context(), int64(itemId))
		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}

		err = queries.BuyItem(r.Context(), db.BuyItemParams{
			ItemID:     int64(itemId),
			BuyerEmail: email,
		})

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}

		resp["msg"] = "Purchased " + item.Name + " successfully"

		utils.JSON(w, http.StatusOK, resp)
	}
}

func SetProductToSoldHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var resp map[string]interface{} = make(map[string]interface{})

		jwtTokenCookie, err := r.Cookie("token")

		if err != nil {
			utils.JSON(w, http.StatusUnauthorized, map[string]string{"error": "No token found"})
			return
		}

		email, errMsg, status := utils.GetUserEmailFromJWT(jwtTokenCookie.Value)
		if errMsg != "" {
			utils.JSON(w, status, map[string]string{"error": errMsg})
			return
		}

		itemId, err := strconv.Atoi(r.URL.Query().Get("item_id"))

		if err != nil {
			utils.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		item, err := queries.GetItem(r.Context(), int64(itemId))
		if err != nil {
			utils.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if item.SellerEmail != email {
			resp["error"] = "You are not the seller of this item"
			utils.JSON(w, http.StatusUnauthorized, resp)
			return
		}

		err = queries.SetToSold(r.Context(), item.ID)

		if err != nil {
			resp["error"] = err.Error()
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}

		resp["msg"] = "Set " + item.Name + " to sold successfully"
		utils.JSON(w, http.StatusOK, resp)
	}
}
