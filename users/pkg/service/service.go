package service

import (
	"context"
	"log"

	"github.com/wotmshuaisi/gokitexample/notificator/pkg/grpc/pb"
	"google.golang.org/grpc"
)

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Create(ctx context.Context, email string) (err error)
}

type basicUsersService struct {
	notificatorSerticeClient pb.NotificatorClient
}

func (b *basicUsersService) Create(ctx context.Context, email string) (err error) {
	// TODO: create user in Database
	reply, err := b.notificatorSerticeClient.SendEmail(context.Background(), &pb.SendEmailRequest{
		Email:   email,
		Content: "Hi! Thank you for registration!",
	})

	if reply != nil {
		log.Printf("Email ID: %s", reply.Id)
	}

	return err
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	conn, err := grpc.Dial("notificator:8082", grpc.WithInsecure())

	if err != nil {
		log.Printf("unable to connect to nofificator: %s", err.Error())
		return new(basicUsersService)
	}

	return &basicUsersService{
		notificatorSerticeClient: pb.NewNotificatorClient(conn),
	}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
