package dbops

//添加用户
func AddUserCredential(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	//3，判断插入结果
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		//3.2 插入失败，返回err
		return err
	}
	//3,1 插入成功，返回nil
	return nil
}

//根据用户名获取用户信息
func GetUserCredential(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?;")
	defer stmt.Close()
	if err != nil {
		return "", err
	}
	var pwd string
	err = stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil {
		return "", err
	}
	return pwd, nil
}

//删除用户
func DeleteUser(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?;")
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

//用户信息更新
