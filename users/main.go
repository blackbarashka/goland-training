package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"users/models"
	"users/sql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Connection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	s, ok := os.LookupEnv("NEW_USER")
	if !ok {
		fmt.Println(os.Stderr, "NEW_USER environment variable not set")
	}
	if s == "YES" {
		var fullname, phoneNumber string
		_, err := fmt.Scan(&fullname, &phoneNumber)
		if err != nil {
			fmt.Println(err)
		}
		user := models.NewUser(fullname, &phoneNumber)
		if err := sql.Insert(ctx, conn, user); err != nil {
			fmt.Println(err)
		}
		return
	} else if s == "NO" {
		if err := sql.Select(ctx, conn); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(os.Stderr, "NEW_USER environment variable not set")
	}
}
