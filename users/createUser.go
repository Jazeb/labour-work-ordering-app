package users;

import  (
"fmt"
"io"
"net/http"
"strconv"

"hello/models"
"hello/shared"
_ "github.com/go-sql-driver/mysql"

_ "github.com/jinzhu/gorm/dialects/mysql" // import these packages for db connection
)

func UserCreate(res http.ResponseWriter, req *http.Request)  {
	fmt.Println("creating user");
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

	fmt.Println(firstName, lastName, email, password, isAdmin, nationalIDStr, ssnStr, userType, service)

	nationalID, err := strconv.Atoi(nationalIDStr);
	ssn, err := strconv.Atoi(ssnStr);

	if err != nil {
		fmt.Println(err)
	}
	user := &models.User{}

	user.FirstName = firstName;
	user.LastName = lastName;
	user.Email = email;
	user.Password = hashedPwd;
	user.IsAdmin = isAdmin;
	user.NationalID = nationalID;
	user.SSN = ssn;
	user.Service = service;

	newerr := db.Where("email = ?", email).First(user).Error
	if newerr == nil {
		fmt.Println("user already exist, use diff email")
		io.WriteString(res, `{status:false, message:"email already exist, use diff email"}`)
		return
	}
	db.Save(&user);
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, `{status:true, message:"user created successfully}`)
}
