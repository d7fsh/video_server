package dbops

// 添加用户
func AddUserCredential(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?,?);")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

// 根据用户名获取用户信息
func GetUser(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?;")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var pwd string
	err = stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil {
		return "", err
	}
	return pwd, nil
}

// 删除用户
func DeleteUser(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

// 用户信息更新
