package models

import(
	//"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type LoginReq = RegisterReq

type SaveReq struct {
	Modified int `json:"Modified"`
	Text []string `json:"text"`
	Pid string `json:"parentId"`
}

type AnalyzeReq struct {
	Modified int `json:"Modified"`
	Input string	`json:"input:"`
}