package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	userpb "github.com/SaidGo/project-protos/proto/user"
	"github.com/SaidGo/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler { return &Handler{svc: svc} }

func toPB(u *user.User) *userpb.User {
	if u == nil {
		return nil
	}
	return &userpb.User{Id: u.ID, Email: u.Email, Name: u.Name}
}

func (h *Handler) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	if in.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	u, err := h.svc.Create(in.GetEmail(), in.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create: %v", err)
	}
	return &userpb.CreateUserResponse{User: toPB(u)}, nil
}

func (h *Handler) GetUserById(ctx context.Context, in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	if in.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	u, err := h.svc.Get(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "get: %v", err)
	}
	return &userpb.GetUserResponse{User: toPB(u)}, nil
}

func (h *Handler) ListUsers(ctx context.Context, in *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, total, err := h.svc.List(int(in.GetPage()), int(in.GetPageSize()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list: %v", err)
	}
	out := make([]*userpb.User, 0, len(users))
	for i := range users {
		out = append(out, toPB(&users[i]))
	}
	return &userpb.ListUsersResponse{
		Users:    out,
		Page:     in.GetPage(),
		PageSize: in.GetPageSize(),
		Total:    uint64(total),
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, in *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	if in.GetUser() == nil || in.GetUser().GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user with id required")
	}
	u, err := h.svc.Update(in.GetUser().GetId(), in.GetUser().GetEmail(), in.GetUser().GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "update: %v", err)
	}
	return &userpb.UpdateUserResponse{User: toPB(u)}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, in *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if in.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if err := h.svc.Delete(in.GetId()); err != nil {
		return nil, status.Errorf(codes.Internal, "delete: %v", err)
	}
	return &userpb.DeleteUserResponse{Deleted: true}, nil
}
