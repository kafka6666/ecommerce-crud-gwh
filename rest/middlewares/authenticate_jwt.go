package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"

	"github.com/kafka6666/ecommerce-crud-gwh/config"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse jwt to get header and payload
		header := r.Header.Get("Authorization")
		if header == "" {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		headerSlice := strings.Split(header, " ")
		if len(headerSlice) != 2 {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		bearerToken := headerSlice[1]
		fmt.Println("==============")
		fmt.Println("Access Token:", bearerToken)

		tokenParts := strings.Split(bearerToken, ".")
		if len(tokenParts) != 3 {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		// decode jwt token
		message := jwtHeader + "." + jwtPayload
		cnf := config.GetConfig()

		byteArrMessage := []byte(message)
		byteArrSignature := []byte(cnf.JwtSecretKey)

		h := hmac.New(sha256.New, byteArrSignature)
		h.Write(byteArrMessage)
		hashedArr := h.Sum(nil)
		newJwtSignature := utils.Base64Encode(hashedArr)

		if jwtSignature != newJwtSignature {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorized. Halar po, tui hacker.")
			return
		}

		next.ServeHTTP(w, r)
	})
}
