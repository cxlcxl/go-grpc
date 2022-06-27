package services

import (
	"context"
)

type SearchService struct {
}

func (ss *SearchService) SearchUser(ctx context.Context, param *SearchParam) (*UserResponse, error) {
	return &UserResponse{RoleId: []int64{1, 2, 3}, Username: param.Username, Email: "666@163.com"}, nil
}

func (ss *SearchService) mustEmbedUnimplementedSearchServiceServer() {}
