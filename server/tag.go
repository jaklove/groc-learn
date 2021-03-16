package server

import (
	"context"
	"encoding/json"
	"go-service/pkg/bapi"
	"go-service/proto"
)

type TagServer struct {}

func NewTagServer()*TagServer  {
	return &TagServer{}
}

func (t *TagServer)GetTagList(ctx context.Context,r *proto.GetTagListRequest)(*proto.GetTagListReply,error)  {
	api := bapi.NewAPI("http://127.0.0.1:8001")
	list, err := api.GetTagList(ctx, r.GetName())
	if err != nil{
		return nil, err
	}

	tagList := proto.GetTagListReply{}
	err = json.Unmarshal(list,&tagList)
	if err != nil{
		return nil, err
	}

	return &tagList,nil
}

