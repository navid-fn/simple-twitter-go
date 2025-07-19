package store

import "database/sql"

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Likes       int    `json:"likes"`
}

type PostgresPostStore struct {
	db *sql.DB
}

func NewPostgresPostStore(db *sql.DB) *PostgresPostStore {
	return &PostgresPostStore{db: db}
}

type PostStore interface {
	CreatePost(*Post) (*Post, error)
	GetPostByID(id int64) (*Post, error)
}

func (pDB *PostgresPostStore) CreatePost(post *Post) (*Post, error) {
	tx, err := pDB.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query :=
		`INSERT INTO posts (title, description)
		VALUES ($1, $2)
		RETURNING id
	`
	err = tx.QueryRow(query, post.Title, post.Description).Scan(&post.ID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (pDB *PostgresPostStore) GetPostByID(id int64) (*Post, error) {
	post := &Post{}
	return post, nil
}
