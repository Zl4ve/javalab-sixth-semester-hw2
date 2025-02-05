package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hw2/internal/core"
	"hw2/proto"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*core.User, error)
}

type UserService struct {
	proto.UserServiceServer
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) GetUser(ctx context.Context, request *proto.UserRequest) (response *proto.UserResponse, err error) {
	user, err := service.userRepository.GetById(ctx, request.GetId())

	if err != nil {
		return nil, err
	}

	if user == nil {
		user = &core.User{
			ID:        primitive.ObjectID{},
			FirstName: "",
			LastName:  "",
		}
	}

	return &proto.UserResponse{Name: user.FirstName}, nil
}
