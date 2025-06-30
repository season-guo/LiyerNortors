package models

import (
	"LiyerNortorsAIpart/internal/db"

	"context"
	"errors"
)

func Login(req *LoginReq) (int, error) {
	pg, err := db.GetDB()
	if err != nil {
		return 0, err
	}

	hashPwd := HashPwd(req.Password)

	userInfo := UserInfo{};
	row := pg.QueryRow(context.TODO(), "Select * from \"user\" where name = $1", req.Name)  
	if err := row.Scan(&userInfo); err != nil {
		return 0, err
	}

	if hashPwd != userInfo.password {
		return 0, errors.New("password incorrect")
	}

	return userInfo.Uid, nil
}