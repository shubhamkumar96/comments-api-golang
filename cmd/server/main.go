package main

import (
	"fmt"
	"github.com/shubhamkumar96/comments-api-golang/internal/comment"
	"github.com/shubhamkumar96/comments-api-golang/internal/db"
	transportHttp "github.com/shubhamkumar96/comments-api-golang/internal/transport/http"
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

	cmtService := comment.NewService(db)
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Comment API in Golang")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
