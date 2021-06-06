package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	id       int
	username string
	password string
}

func main() {
	id := 2

	// get user with id = 2
	user, err := GetUser(id)
	if err != nil {
		fmt.Printf("Original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("Stack trace:\n%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("id: %d, username: %s, password: %s", user.id, user.username, user.password)
}

func GetUser(id int) (Users, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.18.3:10166)/go_db?charset=utf8&parseTime=True")
	if err != nil {
		panic(err.Error())
	}

	// important settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	//Use the DB normally, execute the query etc
	//Prepare statement for writing data
	stmtOut, err := db.Prepare("SELECT id,username,password FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	//query the id
	var u Users
	err = stmtOut.QueryRow(id).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		if err == sql.ErrNoRows {
			//fmt.Println("Err No Rows")
			return u, errors.Wrap(err, "data not found.")
		}
		return Users{}, err
	}

	return u, nil
}
