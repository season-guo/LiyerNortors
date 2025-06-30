package models

import (
	"LiyerNortorsAIpart/internal/db"


	"context"
	"errors"
	"mime/multipart"
)

func GetTagsFromText(text []string) ([]string, error){
	var tags []string

	for _, sentence  := range(text){
		tags = append(tags, sentence)
	}

	return tags, nil
}

func GetTagsFromImg(){}

func Save(uid int, img *multipart.FileHeader, req *SaveReq) error{
	pg, err := db.GetDB()
	if err != nil {
		return err
	}

	tags, err  := GetTagsFromText(req.Text) 
	if err != nil {
		return err
	}

	row := pg.QueryRow(context.TODO(), "SELECT uid from Canvas where pid = $1", req.Pid)
	
	var realUid int
	if err := row.Scan(&realUid); err != nil {
		return err
	}

	if realUid != uid {
		return errors.New("not the right user")
	}

	for _, tag := range(tags){
		if _, err := pg.Exec(context.TODO(),"UPDATE Canvas SET (tags,modified) VALUES(array_append(tags, $1), $2) where pid = $3", tag, req.Modified, req.Pid); err != nil {
			return err
		}
	}

	return nil
}

func Analyze(uid int, input *AnalyzeReq) ([]string, error){
	pg, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	row := pg.QueryRow(context.TODO(),"SELECT * from Canvas where uid = $1", uid)
	
	var Canvas SubCanvas
	if err := row.Scan(&Canvas); err != nil {
		return nil, err
	}

	sentence := []string{input.Input}
	if Canvas.Modified != input.Modified {
		if err := Save(uid, nil, &SaveReq{Modified : input.Modified, Text : sentence, Pid : Canvas.Cid}); err != nil {
			return nil, err
		}
	}

	_, err  = GetTagsFromText(sentence)
	if err != nil {
		return nil, err
	}

	row = pg.QueryRow(context.TODO(),"SELECT tags FROM \"user\" where uid = $1 AND ", uid)
	return nil, nil
}