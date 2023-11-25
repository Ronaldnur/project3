package user_pg

import (
	"database/sql"
	"project3/entity"
	"project3/pkg/errs"
	"project3/repository/user_repository"
)

const (
	CreateNewUsers = `
	INSERT INTO "user"
	(
		full_name,
		email,
		password,
		role
	)
	VALUES ($1, $2, $3, $4)
	RETURNING id,created_at
	`

	getUserByEmailQuery = `
    SELECT id, full_name, email, password, role, created_at, updated_at
    FROM "user"
    WHERE email = $1
`

	updateUserByIdQuery = `
    UPDATE "user"
    SET full_name = $2, email = $3
    WHERE id = $1
	RETURNING id,full_name,email,updated_at
`
	DeleteUserById = `
DELETE FROM "user"
WHERE id = $1
`
	GetUserById = `
SELECT id, full_name, email, password, role, created_at, updated_at
FROM "user"
WHERE id = $1
`
	PatchRole = `
	UPDATE "user"
	SET role=$2
	WHERE id=$1
	RETURNING id, full_name, email,updated_at
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.Repository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) CreateNewUser(newUser entity.User) (*entity.User, errs.MessageErr) {

	var user entity.User

	rows := u.db.QueryRow(CreateNewUsers, newUser.Full_name, newUser.Email, newUser.Password, newUser.Role)

	err := rows.Scan(&user.Id, &user.Created_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {

	var user entity.User
	rows := u.db.QueryRow(getUserByEmailQuery, userEmail)

	err := rows.Scan(&user.Id, &user.Full_name, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &user, nil
}

func (u *userPG) UpdateUser(userId int, newUpdate entity.User) (*entity.User, errs.MessageErr) {
	var userupdate entity.User

	rows := u.db.QueryRow(updateUserByIdQuery, userId, newUpdate.Full_name, newUpdate.Email)

	err := rows.Scan(&userupdate.Id, &userupdate.Full_name, &userupdate.Email, &userupdate.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &userupdate, nil
}

func (u *userPG) DeleteUser(userId int) errs.MessageErr {
	_, err := u.db.Exec(DeleteUserById, userId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")

	}
	return nil
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	rows := u.db.QueryRow(GetUserById, userId)

	err := rows.Scan(&user.Id, &user.Full_name, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &user, nil
}

func (u *userPG) PatchUserRole(userId int, role entity.User) (*entity.User, errs.MessageErr) {
	var userupdate entity.User

	rows := u.db.QueryRow(PatchRole, userId, role.Role)

	err := rows.Scan(&userupdate.Id, &userupdate.Full_name, &userupdate.Email, &userupdate.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}
	return &userupdate, nil
}
