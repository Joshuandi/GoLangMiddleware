package handler

import (
	"GoLangMiddleware/database"
	user "GoLangMiddleware/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandlerInterface interface {
	UserLoginHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) UserHandlerInterface {
	return &UserHandler{db: db}
}

func (u *UserHandler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	switch r.Method {
	case http.MethodGet:
		u.getAllUser(w, r)
		//u.loginUserHandler(w, r)
	case http.MethodPost:
		u.orderUserHandler(w, r, id)
	}
}

func (u *UserHandler) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	// if !midlleware.Auth(w, r) {
	// 	return
	// }
	// if !midlleware.AllowOnlyGet(w, r) {
	// 	return
	// }
	if id := r.URL.Query().Get("Id"); id != "" {
		//OutputJson(w, getAllUser)
		return
	}
	//OutputJson(w, getAllUser)
}
func (u *UserHandler) orderUserHandler(w http.ResponseWriter, r *http.Request, id string) {}
func OutputJson(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (u *UserHandler) getAllUser(w http.ResponseWriter, r *http.Request) {
	var result = []user.User{}
	sqlGet := "Select * from users;"
	rows, err := database.Db.Query(sqlGet)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var userss = user.User{}
		if err = rows.Scan(
			&userss.Id,
			&userss.Username,
			&userss.Email,
			&userss.Password,
			&userss.Age,
			&userss.Division,
			&userss.Created_at,
			&userss.Updated_at,
		); err != nil {
			fmt.Println("No Data", err)
		}
		result = append(result, userss)
	}
	jsonData, _ := json.Marshal(&result)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
