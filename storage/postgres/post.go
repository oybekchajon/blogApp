package postgres

import (
	"github.com/oybekchajon/blogApp/storage/repo"
)

func (m *DBManager) CreatePost(post *repo.Post) (*repo.Post, error) {
	query := `
		INSERT INTO posts(
			title,
			description,
			image_url,
			user_id,
			category_id
		) VALUES($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	

	row := m.db.QueryRow(
		query,
		post.Title,
		post.Description,
		post.ImageUrl,
		post.UserID,
		post.CategoryID,
	)

	err := row.Scan(
		&post.ID,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (m *DBManager) GetPost(id int) (*repo.Post, error){
	var result repo.Post

	query := `
		SELECT
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			views_count
		FROM posts
		WHERE id=$1
	`

	row := m.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.ImageUrl,
		&result.UserID,
		&result.CategoryID,
		&result.CreatedAt,
		&result.ViewsCount,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}