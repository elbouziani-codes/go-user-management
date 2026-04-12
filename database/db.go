package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB


func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "DataUsers.db")
	if err != nil {
		log.Fatalln(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	CreatTables()
	SeedRoles()
}

func CreatTables() {
	contentSql, err := os.ReadFile("database/shema.sql")
	if err != nil {
		log.Fatalln("sir1 ....")
	}
	_, err = DB.Exec(string(contentSql))
	if err != nil {
		log.Fatalln("sir2 ....")
	}
}

func SeedRoles() {
	Slice := []string{"admin", "user" , "moderator"}
	for _, v := range Slice {
		_, err := DB.Exec("INSERT OR IGNORE INTO roles (Name_Role) VALUES (?)", v)
		if err != nil {
			log.Fatalln(v, " : ", err.Error())
		}
	}
}

func GetAllRoles() []string{
	rows , err := DB.Query("SELEC Name_Role FROM roles")
	if err != nil {
		log.Fatalln( err.Error())
	}
	var Slice []string
	for rows.Next() {
		var Role string
		rows.Scan(&Role)
		Slice = append(Slice, Role)
	}
	return Slice
}

func GetIdRoles(NameRool string) int{
	err := DB.QueryRow("SELECT Id FROM R")
}
