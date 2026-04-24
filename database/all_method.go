package database

type User struct {
	Id     int
	Name   string
	Role   string
	Email  string
	Status bool
}
func Adduser(name, email, password string, active, roleID int) error {
	_, err := DB.Exec(
		`INSERT INTO users (user_name, role_id, email,passwords, active)
		VALUES (?, ?, ?, ?,?)`,
		name, roleID, email, password, active,
	)
	if err != nil {
		return err
	}
	return nil
}
func GetAllUsers(){}
func GetIdRoles(NameRool string) (int, error) {
	IdRole := 0
	row := DB.QueryRow("SELECT id FROM roles WHERE Name_role = ?", NameRool)
	err := row.Scan(&IdRole)
	if err != nil {
		return -1, err
	}
	return IdRole, nil
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
