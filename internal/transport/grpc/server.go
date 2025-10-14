package grpc

import (
	"log"
	"net"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	userpb "github.com/SaidGo/project-protos/proto/user"
	"github.com/SaidGo/users-service/internal/user"
)

func RunGRPC(svc *user.Service, addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := ggrpc.NewServer()
	userpb.RegisterUserServiceServer(s, NewHandler(svc))
	reflection.Register(s)
	log.Printf("[users-service] gRPC listen on %s", addr)
	return s.Serve(l)
}
