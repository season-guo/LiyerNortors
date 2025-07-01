package models

import (
	"LiyerNortorsAIpart/internal/db"

	"context"
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func GetNewTagsFromText(c *gin.Context,text []string) ([]string, error){
	var tags []string
	for _, v := range(text){
		tmp, err := Extract(c, v)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tmp...)
	}

	return tags, nil
}

func GetTagsFromImg(){}

func Save(c *gin.Context, uid int, img *multipart.FileHeader, req *SaveReq) error{
	pg, err := db.GetDB()
	if err != nil {
		return err
	}

	row := pg.QueryRow(context.TODO(), `SELECT uid from Canvas where pid = $1`, req.Pid)
	
	var realUid int
	if err := row.Scan(&realUid); err != nil {
		return err
	}

	if realUid != uid {
		return errors.New("not the right user")
	}

	var alltags []string
	row = pg.QueryRow(context.TODO(), `SELECT tags from "user" where uid = $1`, uid)
	if err := row.Scan(alltags); err != nil {
		return err
	}

	tags, err  := GetNewTagsFromText(c, req.Text) 
	if err != nil {
		return err
	}

	for _, tag := range(tags){
		if _, err := pg.Exec(context.TODO(),
		`INSERT INTO Canvas(tags, modified, pid, uid)
		VALUES($1, $2, $3, $4)
		ON CONFLICT (pid, uid)
		UPDATE Canvas 
		SET tags = array_append(tags, $1),
		modified = $2 
		where pid = $3 and uid = $4`,
		tag, req.Modified, req.Pid, uid); err != nil {
			return err
		}
	}

	return nil
}

func Analyze(c *gin.Context, uid int, input *AnalyzeReq) ([]string, error){
	pg, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var alltags []string
	row := pg.QueryRow(context.TODO(), `SELECT tags from "user" where uid = $1`, uid)
	if err := row.Scan(alltags); err != nil {
		return nil, err
	}

	row = pg.QueryRow(context.TODO(),`SELECT * from Canvas where uid = $1`, uid)
	var Canvas []SubCanvas
	if err := row.Scan(&Canvas); err != nil {
		return nil, err
	}

	row = pg.QueryRow(context.TODO(),`SELECT modified from "user" where uid == $1`, uid)
	var modified int
	if err := row.Scan(&modified); err != nil {
		return nil, errors.New("please save the latest change")	
	}
	
	tags, err := Contact(c, alltags, input.Input)
	if err != nil {
		return nil, err
	}

	var targets []string
	row = pg.QueryRow(context.TODO(),`SELECT cid FROM SubCanvas where uid = $1 AND tags && $2`, uid, tags)
	if err = row.Scan(targets); err != nil{
		return nil, err
	}

	return targets, nil
}