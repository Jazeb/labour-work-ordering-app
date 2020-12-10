package users;

import (
	"hello/models"
	"encoding/json"
	"io"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(res http.ResponseWriter, req *http.Request)  {
	user := &models.User{}
	email := req.FormValue("email")
	password := req.FormValue("password")

	err := db.Where("Email = ?", email).First(user).Error;
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		io.WriteString(res, `{"status":false, "message":"User does not exist}`)
		return
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil {
		res.Header().Set("Content-Type", "application/json")
		io.WriteString(res, `{"status":false, "message":"Incorrect password}`)
	}else{
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(user)
	}
}
