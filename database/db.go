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
		log.Fatalln("read schema error:", err)
	}
	_, err = DB.Exec(string(contentSql))
	if err != nil {
		log.Fatalln("exec schema error:", err)
	}
}

func SeedRoles() {
	Slice := []string{"admin", "user", "moderator"}
	for _, v := range Slice {
		_, err := DB.Exec("INSERT OR IGNORE INTO roles (Name_Role) VALUES (?)", v)
		if err != nil {
			log.Fatalln(v, " : ", err.Error())
		}
	}
}

func GetAllRoles() ([]string, error) {
	rows, err := DB.Query("SELECT Name_Role FROM roles")
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()
	var Slice []string
	for rows.Next() {
		var Role string
		err := rows.Scan(&Role)
		if err != nil {
			return []string{}, err
		}
		Slice = append(Slice, Role)
	}
	if err := rows.Err(); err != nil {
		return []string{}, err
	}
	return Slice, nil
}

func GetIdRoles(NameRool string) (int, error) {
	IdRole := 0
	row := DB.QueryRow("SELECT id FROM roles WHERE Name_role = ?", NameRool)
	err := row.Scan(&IdRole)
	if err != nil {
		return -1, err
	}
	return IdRole, nil
}
