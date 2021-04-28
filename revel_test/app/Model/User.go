package Model

import "fmt"

type User struct {
	Id string
	Name string
}

var UserList = []User {{"123","456"},{"456","789"}}

func (u *User) GetName (id string) string {
	length := len(UserList)
	for i := 0 ; i < length ; i++ {
		if UserList[i].Id == id {
			return UserList[i].Name
		}
	}
	return ""
}

func (u *User) AddUser (id string, name string) {
	fmt.Println("Model ? New User! > id : " + id + " / name : " + name)
	UserList = append(UserList, User{id, name})
}

func (u *User) GetUserList (rang int) []User{
	if rang == 0 {
		return UserList
	} else {
		return UserList[:rang]
	}
}