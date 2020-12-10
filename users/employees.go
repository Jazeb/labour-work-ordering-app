package users;

import (
	"io"
	"net/smtp"
	"net/http"
	"strconv"

	"hello/models"
	log "github.com/sirupsen/logrus"
)

func AddEmployee(res http.ResponseWriter, req *http.Request) {
	fullName := req.FormValue("full_name")
	email := req.FormValue("email")
	phoneNumber := req.FormValue("phone_number")

	if (fullName == "" || email == "" || phoneNumber == "") {
		log.Info("Provide required information")
		io.WriteString(res, "provide required data\n")
	}else{
		io.WriteString(res, "everything is here\n")
		log.Info("Sending email to ", email)
		log.Info("email sent, click to signup")
	}
}

func CreateEmployee(res http.ResponseWriter, req *http.Request) {
	Employee := &models.Employees{}

	fullName := req.FormValue("full_name")
	email := req.FormValue("email")
	phoneNumber := req.FormValue("phone_number")
	nationalID := req.FormValue("national_id")
	service := req.FormValue("service")
	ssn, _ := strconv.Atoi("ssn");
	password := req.FormValue("password")
	confirmPassword := req.FormValue("confirm_password")

	if password != confirmPassword {
		log.Fatal("password must be same")
	}

	// ssn, _ := strconv.Atoi(ssnStr);
	Employee.FullName = fullName
	Employee.Email = email
	Employee.PhoneNumber, _ =  strconv.Atoi(phoneNumber)
	Employee.Service = service
	Employee.NationalID, _ = strconv.Atoi(nationalID)
	Employee.SSN = ssn
}
