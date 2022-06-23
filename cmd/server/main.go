package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shubhamkumar96/comments-api-golang/internal/comment"
	"github.com/shubhamkumar96/comments-api-golang/internal/db"
)

//	Run - is going to be responsible for the instantiation and start
//	of our Go Application
func Run() error {
	fmt.Println("Starting Up Our Application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to Connect to the Database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to Migrate Database")
		return err
	}

	fmt.Println("Successfully Connected and Pinged Database")

	id := uuid.New().String()

	fmt.Println(id)

	cmtService := comment.NewService(db)
	cmtService.PostComment(context.Background(),
		comment.Comment{
			ID:     id,
			Slug:   "Manual-test",
			Author: "Tester",
			Body:   "Testing",
		})

	fmt.Println(id)
	fmt.Println(cmtService.GetComment(context.Background(), id))

	return nil
}

func main() {
	fmt.Println("Comment API in Golang")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
