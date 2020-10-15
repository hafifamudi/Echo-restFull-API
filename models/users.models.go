package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/hafif/echoFramework/db"
	"github.com/hafif/echoFramework/helpers"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func RegisterUser(nama string, password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT users (username,password) VALUES (?,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, password)
	if err != nil {
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"lastInsertedID": lastInsertedID,
	}

	return res, nil

}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username=?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.ID, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Println("password and username doesn't match")
		return false, err
	}

	return true, nil
}
