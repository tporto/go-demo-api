package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserReposutory(db *sql.DB) *Users {
	return &Users{db}
}

// Create cria um novo usuário
func (u Users) Create(user models.User) (uint64, error) {
	id := 0

	erro := u.db.QueryRow(
		`insert into users (name, nick, email, password) values ($1, $2, $3, $4) RETURNING id`,
		user.Nome, user.Nick, user.Email, user.Senha,
	).Scan(&id)

	if erro != nil {
		return 0, erro
	}

	return uint64(id), nil
}

// Find busca usuários por nome ou nick
func (u Users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, erro := u.db.Query("select id, name, nick, email, created_at from users where name LIKE $1 or nick LIKE $2", nameOrNick, nameOrNick)
	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if erro = rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByID busca usuário por id
func (u Users) FindByID(id uint64) (models.User, error) {
	rows, erro := u.db.Query("select id, name, nick, email, created_at from users where id = $1", id)
	if erro != nil {
		return models.User{}, erro
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if erro = rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// FindByEmail busca usuário por email
func (u Users) FindByEmail(email string) (models.User, error) {
	rows, erro := u.db.Query("select id, password from users where email = $1", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if erro = rows.Scan(&user.ID, &user.Senha); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// Update
func (u Users) Update(ID uint64, user models.User) error {
	statement, erro := u.db.Prepare(
		"update users set name = $1, nick = $2, email = $3 where id = $4",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(user.Nome, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Delete
func (u Users) Delete(ID uint64) error {
	statement, erro := u.db.Prepare("delete from users where id = $4")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Follow
func (u Users) Follow(userID, followID uint64) error {
	statement, erro := u.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values ($1, $2)",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userID, followID); erro != nil {
		return erro
	}

	return nil
}

// StopFollow
func (u Users) StopFollow(userID, followID uint64) error {
	statement, erro := u.db.Prepare(
		"delete from followers where user_id = $1 and follower_id = $2",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userID, followID); erro != nil {
		return erro
	}

	return nil
}
