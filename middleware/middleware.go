package middleware

import (
	"GoLangMiddleware/config"
	"log"
	"net/http"
	"strings"

	"google.golang.org/api/idtoken"
)

type AuthMiddlewareInterface interface {
	GetToken(next http.Handler) http.Handler
	LoginMiddleware(next http.Handler) http.Handler
}

type AuthMiddleware struct {
	cfg                  config.Config
	googleTokenValidator *idtoken.Validator
}

func NewAuthMiddleware(cfg *config.Config, googleTokenValidator *idtoken.Validator) AuthMiddlewareInterface {
	return &AuthMiddleware{cfg: *cfg, googleTokenValidator: googleTokenValidator}
}

func (m *AuthMiddleware) LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uname, pass, ok := r.BasicAuth(); !ok || uname != m.cfg.Lusername || pass != m.cfg.Lpassword {
			w.WriteHeader(401)
			w.Write([]byte("ERROR DATA"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) GetToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authhandler := strings.Split(r.Header.Get("Authorization"), " "); len(authhandler) == 2 && authhandler[0] == "Bearer" {
			idToken := authhandler[1]
			idTokenPayload, err := idtoken.Validate(r.Context(), idToken, m.cfg.GoogleClientId)
			if err != nil {
				log.Println("Error when validate id token", err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorixed Request"))
				return
			}
			log.Println(idTokenPayload)
			next.ServeHTTP(w, r)
		} else {
			log.Println("Error when validate id token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorixed Request"))
			return
		}
	})
}

// func Auth(w http.ResponseWriter, r *http.Request) bool {
// 	uname, pass, ok := r.BasicAuth()
// 	_ = cleanenv.ReadConfig(".env", &cfg)
// 	if !ok {
// 		w.Write([]byte("Something Wrong "))
// 		w.WriteHeader(401)
// 		return false
// 	}
// 	isValid := (uname == cfg.Lusername) && (pass == cfg.Lpassword)
// 	if !isValid {
// 		w.Write([]byte("Username or Password Incorrect "))
// 		w.WriteHeader(401)
// 		return false
// 	}
// 	return true
// }

// func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
// 	if r.Method != "GET" {
// 		w.Write([]byte("Only Get Allowed"))
// 		return false
// 	}
// 	return true
// }
