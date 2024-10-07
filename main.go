package main

import (
	"context"
	"errors"
	"foo/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	db *gorm.DB
}

func (s *Server) UsersPost(_ context.Context, req *api.User) error {
	err := s.db.Create(req).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) UsersUserIdGet(_ context.Context, params api.UsersUserIdGetParams) (*api.User, error) {
	var user api.User
	err := s.db.First(&user, params.UserId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Server) UsersGet(_ context.Context) ([]api.User, error) {
	var users []api.User
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Server) NewError(_ context.Context, err error) *api.ErrorStatusCode {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &api.ErrorStatusCode{StatusCode: 404, Response: "record not found"}
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return &api.ErrorStatusCode{StatusCode: 409, Response: "duplicated key"}
	}
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return &api.ErrorStatusCode{StatusCode: 409, Response: "foreign key violated"}
	}

	return &api.ErrorStatusCode{StatusCode: 500}
}

func run() error {
	db, err := gorm.Open(postgres.Open("postgres://postgres:@localhost:5432/postgres"), &gorm.Config{TranslateError: true})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&api.User{})
	if err != nil {
		return err
	}
	s := &Server{db: db}
	h, err := api.NewServer(s)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8080", h)
}
