package users

import (
	"errors"
	"strings"

	"github.com/lindseypoche/SELU_ACM/api/clients/mysql"
	"github.com/lindseypoche/SELU_ACM/api/utils/errors/rest"
	"github.com/lindseypoche/SELU_ACM/api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, password) VALUES(?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created FROM users WHERE email=? AND password=?;"
)

// GetByID attempts to get the user from the database with specified id
func (user *User) GetByID() rest.Err {
	stmt, err := mysql.Client.Prepare(queryGetUser)
	if err != nil {
		// logger.Error("error when trying to get user statement", err)
		return rest.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		// user not found
		// logger.Error("error when trying to get user by id", getErr)
		return rest.NewInternalServerError("error when trying to get user", errors.New("database error")) // new
	}
	return nil
}

// Save attempts to save the user into the database
func (user *User) Save() rest.Err {
	stmt, err := mysql.Client.Prepare(queryInsertUser)
	if err != nil {
		// logger.Error("error when trying to prepare save user statement", err)
		return rest.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	// add user to db
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password)
	if saveErr != nil {
		// logger.Error("error when trying to save user", saveErr)
		return rest.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}

	// get the last row (ie. userID) the user was inserted
	userID, err := insertResult.LastInsertId()
	if err != nil {
		// logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}
	user.ID = userID
	return nil
}

// Update updates an existing user's fields in the db
func (user *User) Update() rest.Err {
	stmt, err := mysql.Client.Prepare(queryUpdateUser)
	if err != nil {
		// logger.Error("error when trying to prepare update user statement", err)
		return rest.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}
	defer stmt.Close()

	// attempt to update user
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		// logger.Error("error when trying to update user", err)
		return rest.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}
	return nil
}

// Delete attempts to delete an existing user from the db
func (user *User) Delete() rest.Err {
	stmt, err := mysql.Client.Prepare(queryDeleteUser)
	if err != nil {
		// logger.Error("error when trying to prepare delete user statement", err)
		return rest.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		// logger.Error("error when trying to delete user", err)
		return rest.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	return nil
}

// FindByEmailAndPassword finds user by their email and password
func (user *User) FindByEmailAndPassword() rest.Err {

	stmt, err := mysql.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		// logger.Error("error when trying to prepare get user by email and password statement", err)
		return rest.NewInternalServerError("error when trying to find user", errors.New("database error"))
	}
	defer stmt.Close()

	// find user by email and password and set their status to active
	result := stmt.QueryRow(user.Email, user.Password)
	// populate user fields with the incoming data from the row
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		// user not found
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest.NewNotFoundError("invalid user credentials")
		}
		// logger.Error("error when trying to get user by email and password", getErr)
		return rest.NewInternalServerError("error when trying to find user", errors.New("database error")) // new
	}
	return nil
}
