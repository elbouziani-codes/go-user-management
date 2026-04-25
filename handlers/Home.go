package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"Users/database"
)

type Data struct {
	Errr             string
	LenUser          int
	LenActiveUser    int
	LenNotActiveUser int
	Users            []database.User
	Roles            []string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	data := &Data{}
	var err error
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

		if name == "" || email == "" || role == "" || password == "" || status == "" || (status != "true" || status != "false") {
			data.Errr = "Empty Value"
		}

		active := 0
		if status == "True" {
			active = 1
		}

		roleID, err := database.GetIdRoles(role)
		if err != nil {
			ErrorHandler(w, "role not found", 400)
			return
		}

		err = database.Adduser(name, email, password, active, roleID)
		if err != nil {
			fmt.Println("B")
			ErrorHandler(w, "cannot insert user", 500)
			return
		}
		http.Redirect(w, r, "/", 200)
	}

	data.Roles, err = database.GetAllRoles()
	if err != nil {
		log.Fatalln(err)
	}

	temp, err := template.ParseFiles("temp/html.html")
	if err != nil {
	}
	fmt.Println("A")

	data.Users, err = database.GetAllUsers()
	if err != nil {
		log.Fatalln(err)
	}
	data.LenActiveUser = 0
	data.LenNotActiveUser = 0
	data.LenUser = 0
	var b bytes.Buffer
	err = temp.Execute(&b, data)
	if err != nil {
					fmt.Println("100")

		ErrorHandler(w, "error in Execute", 500)
	}
	w.Write(b.Bytes())
}
