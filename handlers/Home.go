package handlers

import (
	"bytes"
	"net/http"
	"text/template"

	"Users/database"
)

type Data struct {
	// Errr             string
	LenUser          int
	LenActiveUser    int
	LenNotActiveUser int
	Users            []user
	Roles           []string
}

type user struct {
	Id     int
	Name   string
	Role   string
	Email  string
	Status bool
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := &Data{}
	data.Roles = database.Slice
	temp, err := template.ParseFiles("temp/html.html")
	if err != nil {
	}

	name := ""
	email := ""
	role := ""
	password := ""
	status := ""
	if r.Method == http.MethodPost {
		name = r.FormValue("Name")
		email = r.FormValue("Email")
		role = r.FormValue("Role")
		password = r.FormValue("Password")
		status = r.FormValue("Status")
		if name == "" || email == "" || role == "" || status == "" || password == "" {
			// data.Errr = "Empty Value"
		}
		//		var boole = 1
		//		if status == false {
		//			boole = 0
		//		}
		//		database.DB.Exec("INSERT INTO users (user_name ,passwords, role_id, email , active) values (? ,? ,? ,?,?)",name ,password,,email,boole)

	}

	data.LenActiveUser = 0
	data.LenNotActiveUser = 0
	data.LenUser = 0
	var b bytes.Buffer
	err = temp.Execute(&b, data)
	if err != nil {
		ErrorHandler(w, "error in Execute", 500)
	}
	w.Write(b.Bytes())
}
