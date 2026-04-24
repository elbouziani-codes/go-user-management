package database

import (
	"log"
	"os"
)

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
