package service

import (
	"context"
	"fmt"
	"log"

	"github.com/Bek0sh/online-market/main-page/internal/models"
	"github.com/Bek0sh/online-market/main-page/pkg/proto"
	"google.golang.org/grpc"
)

type grpcClient struct {
	client proto.UserInfoClient
}

func newGrpcClient() *grpcClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewUserInfoClient(conn)

	return &grpcClient{client: client}
}

func (g *grpcClient) GetCurrentUser() (*models.UserInfo, error) {
	request := proto.Empty{}
	response, err := g.client.GetCurrentUser(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp := models.UserInfo{
		Id:          int(response.GetId()),
		FullName:    fmt.Sprintf("%s %s", response.GetName(), response.GetSurname()),
		UserType:    response.GetUserRole(),
		PhoneNumber: response.GetPhoneNumber(),
	}

	return &resp, nil
}

func (g *grpcClient) GetUserById(id int) (*models.UserInfo, error) {
	request := proto.GetUserByIdRequest{Id: int32(id)}
	response, err := g.client.GetUserById(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp := models.UserInfo{
		Id:          id,
		FullName:    fmt.Sprintf("%s %s", response.GetName(), response.GetSurname()),
		UserType:    response.GetUserRole(),
		PhoneNumber: response.GetPhoneNumber(),
	}

	return &resp, nil
}

func (g *grpcClient) CheckUserType() error {
	request := proto.Empty{}
	_, err := g.client.CheckRole(context.Background(), &request)
	return err
}

func (g *grpcClient) CheckToken() error {
	request := proto.Empty{}
	_, err := g.client.CheckToken(context.Background(), &request)
	return err
}
