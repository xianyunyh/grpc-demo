package handler

import (
	"context"
	"fmt"
	pb "gRPC-study/models"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostHandler struct {
	pb.UnimplementedPostServer
}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}
func (p *PostHandler) ListPost(ctx context.Context, req *pb.ListPostReuquest) (*pb.ListPostResponse, error) {
	page := req.Page
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	fmt.Println(req)
	rows := make([]*pb.PostItem, 0, pageSize)
	for i := 0; i < int(pageSize); i++ {
		rows = append(rows, &pb.PostItem{
			Id:       int32(i),
			Title:    "",
			Pic:      "sfsf",
			Tags:     []string{},
			Content:  "sfdsfs",
			PostTime: timestamppb.Now(),
		})
	}
	resp := &pb.ListPostResponse{
		Code:    0,
		Message: "",
		Data: &pb.ListPostResponse_ListPostData{
			Page:     page,
			PageSize: pageSize,
			Total:    100,
			Rows:     rows,
		},
	}
	return resp, nil
}

func (p *PostHandler) PostDetail(ctx context.Context, req *pb.GetPostOne) (*pb.PostItem, error) {

	return &pb.PostItem{
		Id:       req.Id,
		Title:    "",
		Pic:      "",
		Tags:     []string{},
		Content:  "",
		PostTime: timestamppb.Now(),
	}, nil
}

func (p *PostHandler) AddPost(ctx context.Context, req *pb.PostItem) (*pb.ResponseData, error) {
	return &pb.ResponseData{}, nil
}

func (p *PostHandler) UpdatePost(ctx context.Context, req *pb.PostItem) (*pb.ResponseData, error) {
	return &pb.ResponseData{}, nil
}

func (p *PostHandler) DeletePost(ctx context.Context, req *pb.GetPostOne) (*pb.ResponseData, error) {
	return &pb.ResponseData{}, nil
}
