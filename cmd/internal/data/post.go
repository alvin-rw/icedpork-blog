package data

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNotFound error = fmt.Errorf("requested resource could not be found")
)

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Published   bool      `json:"published"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
}

type PostModel struct {
	DB *sql.DB
}

func (p *PostModel) GetPost(id int) (*Post, error) {
	query := `
		SELECT title, content, published, created_at, last_updated FROM posts
		WHERE id = $1`

	post := Post{}

	err := p.DB.QueryRow(query, id).Scan(
		&post.Title,
		&post.Content,
		&post.Published,
		&post.CreatedAt,
		&post.LastUpdated,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return &post, nil
}

func (p *PostModel) GetPosts(pageSize int) ([]Post, error) {
	query := `
		SELECT id, title, content, published, created_at, last_updated FROM posts
		ORDER BY id DESC
		LIMIT $1`

	var posts []Post

	rows, err := p.DB.Query(query, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		post := Post{}

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Published,
			&post.CreatedAt,
			&post.LastUpdated,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (p *PostModel) Create(post *Post) error {
	query := `
		INSERT INTO posts (title, content, published)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, last_updated`

	args := []interface{}{
		post.Title,
		post.Content,
		post.Published,
	}

	err := p.DB.QueryRow(query, args...).Scan(&post.ID, &post.CreatedAt, &post.LastUpdated)
	if err != nil {
		return err
	}

	return nil
}
