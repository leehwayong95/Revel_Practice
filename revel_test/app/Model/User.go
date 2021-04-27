package Model

type User struct {
	Id string
	name string
}

var UserList []User

func GetName (id string) string {
	length := len(UserList)
	for i := 0 ; i < length ; i++ {
		if UserList[i].Id == id {
			return UserList[i].name
		}
	}
	return ""
}

func AddUser (id string, name string) {
	UserList = append(UserList, User{id, name})
}

func GetUserList (rang int) []User{
	if rang == 0 {
		return UserList
	} else {
		return UserList[:rang]
	}
}