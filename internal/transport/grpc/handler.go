package grpc

import (
	"context"

	userpb "github.com/SaidGo/project-protos/proto/user"
	"github.com/SaidGo/users-service/internal/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	userpb.UnimplementedUserServiceServer
	svc *user.Service
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	if in.GetEmail() == "" || in.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "email and name are required")
	}
	u, err := h.svc.Create(in.GetEmail(), in.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create: %v", err)
	}
	return &userpb.CreateUserResponse{User: toPB(u)}, nil
}

func (h *Handler) GetUserById(ctx context.Context, in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	id := in.GetId() // uint64
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id required")
	}
	u, err := h.svc.Get(id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "get: %v", err)
	}
	return &userpb.GetUserResponse{User: toPB(u)}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, in *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	if in.GetUser() == nil || in.GetUser().GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user.id required")
	}
	u := in.GetUser()
	upd, err := h.svc.Update(u.GetId(), u.GetEmail(), u.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "update: %v", err)
	}
	return &userpb.UpdateUserResponse{User: toPB(upd)}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, in *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id := in.GetId() // uint64
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id required")
	}
	if err := h.svc.Delete(id); err != nil {
		return nil, status.Errorf(codes.Internal, "delete: %v", err)
	}
	return &userpb.DeleteUserResponse{Deleted: true}, nil
}

func (h *Handler) ListUsers(ctx context.Context, in *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	// Дефолты пагинации: page>=1, page_size=10
	page := int(in.GetPage())
	if page == 0 {
		page = 1
	}
	pageSize := int(in.GetPageSize())
	if pageSize == 0 {
		pageSize = 10
	}

	users, total, err := h.svc.List(page, pageSize)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list: %v", err)
	}

	resp := &userpb.ListUsersResponse{
		Users:    toPBList(users),
		Page:     uint32(page),
		PageSize: uint32(pageSize),
		Total:    uint64(total),
	}
	return resp, nil
}

// --- helpers ---

func toPB(u *user.User) *userpb.User {
	if u == nil {
		return nil
	}
	return &userpb.User{
		Id:    uint64(u.ID),
		Email: u.Email,
		Name:  u.Name,
	}
}

func toPBList(list []user.User) []*userpb.User {
	out := make([]*userpb.User, 0, len(list))
	for _, v := range list {
		uu := v
		out = append(out, toPB(&uu))
	}
	return out
}
