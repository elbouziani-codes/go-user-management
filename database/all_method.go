package database

import "log"

type User struct {
	Id       int
	Name     string
	Role     string
	Email    string
	Password string
	Status   bool
}

func Adduser(name, email, password string, active, roleID int) error {
	_, err := DB.Exec(
		`INSERT INTO users (user_name, role_id, email,passwords, active)
		VALUES (?, ?, ?, ?,?)`,
		name, roleID, email, password, active,
	)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	var us []User
	q, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return []User{}, err
	}

	status := 0
	role := 1
	for q.Next() {
		var u User
		err = q.Scan(&u.Id, &u.Name, &role, &u.Email, &u.Password, &status)
		if err != nil {
			return []User{}, err
		}
		if status == 1 {
			u.Status = true
		} else {
			u.Status = false
		}
		nameRole, err := GetNameRoles(role)
		if err != nil {
			return []User{}, err
		}
		u.Role = nameRole
		us = append(us, u)
	}
	return us, nil
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

func GetNameRoles(IdRool int) (string, error) {
	NameRole := ""
	row := DB.QueryRow("SELECT Name_role FROM roles WHERE id = ?", IdRool)
	err := row.Scan(&NameRole)
	if err != nil {
		return "", err
	}
	return NameRole, nil
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
