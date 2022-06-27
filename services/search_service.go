package services

import (
	"context"
)

type SearchService struct {
}

func (ss *SearchService) SearchUser(ctx context.Context, param *SearchParam) (*UserResponse, error) {
	roles := []*Role{{RoleName: "用户", Permissions: []string{"login", "insert"}}}
	return &UserResponse{RoleId: []int64{1, 2, 3}, Username: param.Username, Email: "666@163.com", Role: roles}, nil
}

func (ss *SearchService) mustEmbedUnimplementedSearchServiceServer() {}
