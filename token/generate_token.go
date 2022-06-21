package token

import (
	"GoLangMiddleware/config"
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/api/idtoken"
)

var (
	cfg                  config.Config
	googleTokenValidator *idtoken.Validator
)



func getToken(next http.Handler) http.Handler {
	_ = cleanenv.ReadConfig(".env", &cfg)
	googleTokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authhandler := strings.Split(r.Header.Get("Authorization"), ""); len(authhandler) == 2 && authhandler[0] == "Bearer" {
			idToken := authhandler[1]
			idTokenPayload, err := googleTokenValidator.Validate(r.Context(), idToken, cfg.GoogleClientId)
			if err != nil {
				log.Println("Error when validate id token", err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorixed Request"))
				return
			}
			log.Println(idTokenPayload)
			next.ServeHTTP(w, r)
		} else {
			log.Println("Error when validate id token", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorixed Request"))
			return
		}
	})
}
