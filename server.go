package main

import (
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import these packages for db connection

	"hello/models"
	"hello/users"
)

var db, _ = gorm.Open("mysql", "testuser:password@/FZ3A?charset=utf8&parseTime=True&loc=Local");

func health(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application.json")
	io.WriteString(res, `{"status":true, "message":"server is running"}`)
}

func main()  {
	defer db.Close()
	fmt.Println("starting golang server...")
	router := mux.NewRouter();

	// db.Debug().DropTableIfExists(&userModel{})
	// db.Debug().AutoMigrate(&users{})

	us := &models.User{}

	router.HandleFunc("/admin/addEmployee", users.AddEmployee).Methods("POST");
	router.HandleFunc("/employee/create", users.CreateEmployee).Methods("POST");

	router.HandleFunc("/", health).Methods("GET");
	router.HandleFunc("/user/create", users.UserCreate).Methods("POST");
	router.HandleFunc("/user/login", users.UserLogin).Methods("POST");
	router.HandleFunc("/user/update", users.UpdateUser).Methods("POST");
	http.ListenAndServe(":8000", router)
}
