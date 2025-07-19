package store

import "database/sql"

type User struct {
	ID             int
	Username       string
	Email          string
	HashedPassword string
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{db: db}
}

type UserStore interface {
	CreateUser(user *User) (*User, error)
	GetUserByID(id int64) (*User, error)
}

func (pgDB *PostgresUserStore) CreateUser(user *User) (*User, error) {
	tx, err := pgDB.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	query :=
		`
	INSERT INTO users (username, hashed_password, email)
	VALUES ($1, $2, $email)
	RETURNING id
	`

	err = tx.QueryRow(query, user.Username, user.HashedPassword, user.Email).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (pDB *PostgresUserStore) GetUserByID(id int64) (*User, error) {
	user := &User{}
	return user, nil
}
