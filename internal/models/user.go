package models

type UserInfo struct {
	Uid int
	Name string
	password string 
	AllTags []string
	Modified int
}

type SubCanvas struct {
	Uid int 
	Cid  string
	Tags []string
}