package main

import (
	"fmt"
	"git_truongvudinh/go_web/internal/common"
	"git_truongvudinh/go_web/internal/entity"
	"git_truongvudinh/go_web/internal/repository"
	"git_truongvudinh/go_web/internal/usecase"
	dao "git_truongvudinh/go_web/pkg"
	"log"
	"time"
)

func main() {
	dsn := "root:truong@tcp(localhost:3306)/go_web_example"
	db, err := dao.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database")

	userRepo := repository.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)

	newUser := &entity.User{
		FirstName: "tuyen",
		LastName:  "nguyen",
		Email:     "tuyenhihi2223@gmail.com",
		Password:  common.HashPassword("hihi"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser := userService.CreateNewUser(newUser)
	fmt.Printf("Created user: %+v\n", createdUser)

	defer db.Close()
}
