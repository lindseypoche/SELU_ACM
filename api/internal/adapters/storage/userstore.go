package storage

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	"errors"

// 	// Driver
// 	_ "github.com/go-sql-driver/mysql"

// 	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
// 	"github.com/lindseypoche/SELU_ACM/api/internal/utils/errors/rest"
// )

// const (
// 	mysqlUsersUsername = "mysql_users_username"
// 	mysqlUsersPassword = "mysql_users_password"
// 	mysqlUsersHost     = "mysql_users_host"
// 	mysqlUsersScheme   = "mysql_users_schema"

// 	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, password) VALUES(?, ?, ?, ?, ?);"
// 	queryGetUser                = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
// 	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
// 	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
// 	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created FROM users WHERE email=? AND password=?;"
// )

// var (
// 	username = os.Getenv(mysqlUsersUsername)
// 	password = os.Getenv(mysqlUsersPassword)
// 	host     = os.Getenv(mysqlUsersHost)
// 	scheme   = os.Getenv(mysqlUsersScheme)
// )

// type mysqlRepo struct {
// 	db        *sql.DB
// 	tableName string
// 	host      string
// }

// func newMySQLClient(host, table string) (*sql.DB, error) {

// 	// only for development
// 	os.Setenv("root", mysqlUsersUsername)
// 	os.Setenv("root", mysqlUsersPassword)
// 	os.Setenv(host, mysqlUsersHost)    // 127.0.0.1:3306
// 	os.Setenv(table, mysqlUsersScheme) // users_db

// 	// define datasource name. // user:password@tcp(host)/schema?charset=utf8
// 	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
// 		username,
// 		password,
// 		host,
// 		scheme,
// 	)

// 	var err error
// 	db, err := sql.Open("mysql", dataSourceName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// ping test database
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	// mysql.SetLogger(logger.GetLogger())
// 	log.Println("database successfully configured")

// 	return db, nil
// }

// // NewMySQLRepository creates a new mysql repository
// func NewMySQLRepository(host, tableName string) domain.UserRepository {
// 	repo := &mysqlRepo{
// 		host:      host,
// 		tableName: tableName,
// 	}
// 	db, err := newMySQLClient(host, tableName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	repo.db = db
// 	return repo
// }

// // GetByID attempts to get the user from the database with specified id
// func (r *mysqlRepo) GetByID(id int64) (*domain.User, rest.Err) {
// 	stmt, err := r.db.Prepare(queryGetUser)
// 	if err != nil {
// 		// logger.Error("error when trying to get user statement", err)
// 		return nil, rest.NewInternalServerError("error when trying to get user", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	user := domain.User{}
// 	result := stmt.QueryRow(id)

// 	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
// 		// logger.Error("error when trying to get user by id", getErr)
// 		return nil, rest.NewInternalServerError("error when trying to get user", errors.New("database error")) // new
// 	}
// 	return &user, nil
// }

// // Save attempts to save the user into the database
// func (r *mysqlRepo) Save(user *domain.User) (*domain.User, rest.Err) {
// 	stmt, err := r.db.Prepare(queryInsertUser)
// 	if err != nil {
// 		// logger.Error("error when trying to prepare save user statement", err)
// 		return nil, rest.NewInternalServerError("error when trying to get user", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	// add user to db
// 	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password)
// 	if saveErr != nil {
// 		// logger.Error("error when trying to save user", saveErr)
// 		return nil, rest.NewInternalServerError("error when trying to get user", errors.New("database error"))
// 	}

// 	// get the last row (ie. userID) the user was inserted
// 	userID, err := insertResult.LastInsertId()
// 	if err != nil {
// 		// logger.Error("error when trying to get last insert id after creating a new user", err)
// 		return nil, rest.NewInternalServerError("error when trying to save user", errors.New("database error"))
// 	}
// 	user.ID = userID
// 	return user, nil
// }
