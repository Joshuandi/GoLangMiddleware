package middleware

import (
	"GoLangMiddleware/config"
	"fmt"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
)

var cfg config.Config

func Auth(w http.ResponseWriter, r *http.Request) bool {
	uname, pass, ok := r.BasicAuth()
	_ = cleanenv.ReadConfig(".env", &cfg)
	if !ok {
		w.Write([]byte("Something Wrong "))
		w.WriteHeader(401)
		return false
	}
	isValid := (uname == cfg.Lusername) && (pass == cfg.Lpassword)
	fmt.Println(cfg.Lusername, cfg.Lpassword)
	if !isValid {
		w.Write([]byte("Username or Password Incorrect "))
		w.WriteHeader(401)
		return false
	}
	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only Get Allowed"))
		return false
	}
	return true
}
