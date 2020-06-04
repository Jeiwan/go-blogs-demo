package db

import "time"

type Post struct {
	ID   int
	Date time.Time
	Post string
}

type CreatePostRequest struct {
	Date time.Time
	Post string
}

func (db DB) CreatePost(request CreatePostRequest) (int, error) {
	var postID int

	rows, err := db.db.NamedQuery(
		`INSERT INTO posts (date, post) VALUES (:date, :post) RETURNING id`,
		map[string]interface{}{
			"date": request.Date,
			"post": request.Post,
		},
	)
	if err != nil {
		return -1, err
	}

	for rows.Next() {
		if err := rows.Scan(&postID); err != nil {
			return -1, err
		}
	}

	return postID, nil
}

type GetPostsRequest struct{}

func (db DB) GetPosts(request GetPostsRequest) ([]Post, error) {
	var posts []Post
	if err := db.db.Select(&posts, "SELECT * FROM posts ORDER BY date DESC"); err != nil {
		return nil, err
	}

	return posts, nil
}
