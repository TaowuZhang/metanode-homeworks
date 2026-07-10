package main

import (
	"fmt"
	"log"

	"metanode-go-backend-homeworks/internal/gormpractice"
)

func main() {
	db, err := gormpractice.InitDB("gorm_practice.db")
	if err != nil {
		log.Fatal(err)
	}
	user, err := gormpractice.Seed(db)
	if err != nil {
		log.Fatal(err)
	}
	posts, err := gormpractice.GetUserPostsWithComments(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user %s has %d posts\n", user.Username, len(posts))
	for _, post := range posts {
		fmt.Printf("- %s comments=%d status=%s\n", post.Title, len(post.Comments), post.CommentStatus)
	}
	post, count, err := gormpractice.GetPostWithMostComments(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("most commented post: %s (%d comments)\n", post.Title, count)
}
