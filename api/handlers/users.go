package handlers

import (
	"database/sql"
	"net/http"

	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
	"github.com/touch-some-grass-bro/vitamart/utils"
)

func SetHostel(queries *db.Queries) http.HandlerFunc {
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

    hostel := r.URL.Query().Get("hostel")

    if hostel == "" {
      resp["error"] = "Hostel name cannot be empty"
      utils.JSON(w, http.StatusBadRequest, resp)
      return
    }

    err = queries.UpdateHostel(r.Context(), db.UpdateHostelParams{
    	Email:  email,
    	Hostel: sql.NullString{
    		String: hostel,
    		Valid:  true,
    	},
    })

  }
}

