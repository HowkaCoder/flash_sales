package handler

import (
	"auth_service/kafka"
	"auth_service/model"
	pb "auth_service/protoc/auth"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB        *gorm.DB
	JwtSecret []byte
	pb.UnimplementedAuthServiceServer
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := model.User{
		Name:    req.Name,
		Address: req.Address,
		Email:   req.Email,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.RegisterResponse{Msg: "Error with password hashing", Status: 500}, err
	}
	user.Password = string(hashedPassword)

	if err := s.DB.Create(&user).Error; err != nil {
		return &pb.RegisterResponse{Msg: "Failed to register user", Status: 500}, err
	}
	return &pb.RegisterResponse{Msg: "User successfully registered", Status: 200}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user model.User
	if err := s.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return &pb.LoginResponse{Status: 404}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return &pb.LoginResponse{Status: 401}, errors.New("invalid password")
	}

	token, err := s.generateJWT(user.ID)
	if err != nil {
		return &pb.LoginResponse{Status: 500}, err
	}
	kafka.NotifyKafka("users", "Register - Request", "New User Lol!")
	return &pb.LoginResponse{Token: token, Status: 200}, nil
}

func (s *AuthService) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	var users []model.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      uint32(u.ID),
			Name:    u.Name,
			Email:   u.Email,
			Address: u.Address,
		})
	}

	return &pb.GetAllUsersResponse{Users: pbUsers}, nil
}

func (s *AuthService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	var user model.User
	if err := s.DB.First(&user, req.UserId).Error; err != nil {
		return &pb.GetUserByIdResponse{Status: 404}, err
	}

	return &pb.GetUserByIdResponse{
		Status: 200,
		User: &pb.User{
			Id:      uint32(user.ID),
			Name:    user.Name,
			Email:   user.Email,
			Address: user.Address,
		},
	}, nil
}

func (s *AuthService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var user model.User
	if err := s.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return &pb.UpdateUserResponse{Msg: "User not found", Status: 404}, err
	}

	user.Name = req.Name
	user.Address = req.Address

	if err := s.DB.Save(&user).Error; err != nil {
		return &pb.UpdateUserResponse{Msg: "Failed to update user", Status: 500}, err
	}

	return &pb.UpdateUserResponse{Msg: "User successfully updated", Status: 200}, nil
}

func (s *AuthService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if err := s.DB.Delete(&model.User{}, req.Id).Error; err != nil {
		return &pb.DeleteUserResponse{Msg: "Failed to delete user", Status: 500}, err
	}

	return &pb.DeleteUserResponse{Msg: "User successfully deleted", Status: 200}, nil
}

func (s *AuthService) ValidateJWTToken(ctx context.Context, req *pb.ValidateJWTTokenRequest) (*pb.ValidateJWTTokenResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return &pb.ValidateJWTTokenResponse{Valid: false, UserId: 0}, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return &pb.ValidateJWTTokenResponse{Valid: false, UserId: 0}, nil
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return &pb.ValidateJWTTokenResponse{Valid: false, UserId: 0}, nil
	}

	return &pb.ValidateJWTTokenResponse{Valid: true, UserId: uint32(userId)}, nil
}

func (s *AuthService) generateJWT(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.JwtSecret)
}
