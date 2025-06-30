package models

type UserInfo struct {
	Uid int
	Name string
	password string 
	AllTags []string
}

type SubCanvas struct {
	Uid int 
	Modified int
	Cid  string
	Tags []string
}