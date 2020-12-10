package users;

import (
	"io"
	"fmt"
	"net/http"
	"strconv"

	"hello/models"
	"hello/shared"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db, _ = gorm.Open("mysql", "testuser:password@/FZ3A?charset=utf8&parseTime=True&loc=Local");
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	User := &models.User{}
	firstName := req.FormValue("first_name")
	lastName := req.FormValue("last_name")
	email := req.FormValue("email")
	password := req.FormValue("password")
	isAdmin, _ := strconv.ParseBool(req.FormValue("is_admin"))
	hashedPwd := shared.CreatePasswords(password)
	nationalIDStr := req.FormValue("national_id") // need to convert it to int
	ssnStr := req.FormValue("ssn") // need to convert values to int as they are recieved in string
	userType := req.FormValue("user_type")
	service := req.FormValue("service")

	nationalID, err := strconv.Atoi(nationalIDStr);
	if err != nil {
		fmt.Println(err)
	}
	ssn, err := strconv.Atoi(ssnStr);

	User.FirstName = firstName;
	User.LastName = lastName;
	User.Email = email;
	User.Password = hashedPwd;
	User.IsAdmin = isAdmin;
	User.NationalID = nationalID;
	User.SSN = ssn;
	User.Service = service;
	User.UserType = userType;

	db.Where("Email = ?", email).First(User).Save(User)
	io.WriteString(res, `user saved`)
}
