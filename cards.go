package cards

type WordCard struct {
	Id         int    `json:"id"`
	Firstword  string `json:"firstword"`
	Secondword string `json:"secondword"`
	frequency  int    `json:"frequency`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}
