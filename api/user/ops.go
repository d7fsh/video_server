package user

import "video_server/api/dbops"

func DeleteUser(loginName, pwd string) error {
	err := dbops.DeleteUser(loginName, pwd)
	return err
}

func AddUser(loginName, pwd string) error {
	err := dbops.AddUserCredential(loginName, pwd)
	return err
}
