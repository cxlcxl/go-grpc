package services

import (
	"context"
	"io"
)

type UserService struct {
}

func (us *UserService) SearchUser(ctx context.Context, param *SearchParam) (*UserResponse, error) {
	return &UserResponse{Username: param.Username, Role: []*Role{{RoleName: "admin"}}}, nil
}
func (us *UserService) GetUserSteam(stream SearchService_GetUserSteamServer) error {
	var users = new(UsersResponse)
	users.Users = make([]*UserResponse, 0)
	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				_ = stream.SendAndClose(users)
				break
			}
			return err
		}
		users.Users = append(users.Users, &UserResponse{Username: recv.Username})
	}
	return nil
}
func (us *UserService) mustEmbedUnimplementedSearchServiceServer() {}
