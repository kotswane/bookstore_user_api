package users

import (
	"github.com/kotswane/bookstore_user_api/utils/date_utils"
	"github.com/kotswane/bookstore_user_api/utils/errors"
	"github.com/kotswane/bookstore_user_api/utils/mysql_utils"
	"github.com/kotswane/datasource/mysql/users_db"
)

var (
	userDB = make(map[int16]*User)
)

const (
	queryInsertUser = "INSERT INTO users (first_name,last_name,email,date_created) VALUES (?,?,?,?)"
	gueryGetUser    = "SELECT * FROM users WHERE Id = ?"
	queryUpdateUser = "UPDATE users SET first_name=?,last_name=?,email=? WHERE Id=?"
	queryDeleteUser = "DELETE FROM users WHERE Id = ?"
	guerySearchUser = "SELECT * FROM users WHERE Status = ?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(gueryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date_Created); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.Date_Created = date_utils.GetNowString()
	insertStatement, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Date_Created)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	insertId, err := insertStatement.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = insertId

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	updateStatement, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	_, err = updateStatement.RowsAffected()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(guerySearchUser)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, getErr := stmt.Query(status)
	if getErr != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date_Created, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}

	return result, nil
}
