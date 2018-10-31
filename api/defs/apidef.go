package defs

//requests
type UserCredential struct {
	Username string `json:"user_name"`//打tag
	Pwd string `json:"pwd"`
}

//Data model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}