package gormpractice

import (
	"testing"

	"gorm.io/gorm"
)

func testDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := InitDB("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("init db: %v", err)
	}
	return db
}

func TestSeedAndQueries(t *testing.T) {
	db := testDB(t)
	user, err := Seed(db)
	if err != nil {
		t.Fatalf("seed: %v", err)
	}
	var refreshed User
	if err := db.First(&refreshed, user.ID).Error; err != nil {
		t.Fatalf("get user: %v", err)
	}
	if refreshed.PostCount != 2 {
		t.Fatalf("post count got %d", refreshed.PostCount)
	}
	posts, err := GetUserPostsWithComments(db, user.ID)
	if err != nil {
		t.Fatalf("posts: %v", err)
	}
	if len(posts) != 2 || len(posts[0].Comments) != 2 {
		t.Fatalf("unexpected posts: %+v", posts)
	}
	post, count, err := GetPostWithMostComments(db)
	if err != nil {
		t.Fatalf("most comments: %v", err)
	}
	if post.Title != "GORM 入门" || count != 2 {
		t.Fatalf("got post=%+v count=%d", post, count)
	}
}

func TestCommentDeleteHook(t *testing.T) {
	db := testDB(t)
	user := User{Username: "bob", Email: "bob@example.com"}
	if err := db.Create(&user).Error; err != nil {
		t.Fatal(err)
	}
	post := Post{Title: "one", Content: "body", UserID: user.ID}
	if err := db.Create(&post).Error; err != nil {
		t.Fatal(err)
	}
	comment := Comment{Content: "first", UserID: user.ID, PostID: post.ID}
	if err := db.Create(&comment).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.First(&post, post.ID).Error; err != nil {
		t.Fatal(err)
	}
	if post.CommentStatus != CommentStatusHasComments {
		t.Fatalf("after create got %q", post.CommentStatus)
	}
	if err := db.Delete(&comment).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.First(&post, post.ID).Error; err != nil {
		t.Fatal(err)
	}
	if post.CommentStatus != CommentStatusNoComments {
		t.Fatalf("after delete got %q", post.CommentStatus)
	}
}
