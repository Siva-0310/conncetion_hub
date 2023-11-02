package database

import (
	"connection_hub/auth_service/structs"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func ResgisterUser(user *structs.User, tx *sql.Tx) int64 {
	sql_statement := "INSERT INTO users(user_name,user_password,user_email) VALUES(?,?,?)"
	res, err := tx.Exec(sql_statement, user.User_name, user.User_password, user.User_email)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return -2
		}
		return -1
	}
	if ans, err := res.LastInsertId(); err == nil {
		return ans
	}
	return -1
}
