package postgresql

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (c *connectDB) CreateUser(id string, password string, name string, phone string, email string) bool {

	sql_statement := "INSERT INTO users (sessionid, id, password, name, phone, email, createdAt) VALUES ('', $1, $2, $3, $4, $5, now());"
	_, err := c.db.Exec(sql_statement, id, password, name, phone, email)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Created User")

	return err == nil
}

func (c *connectDB) ReadUser(id string, password string) (int, UserInfo) {
	var user UserInfo
	sql_statement := "SELECT * FROM users WHERE id=$1"

	err := c.db.QueryRow(sql_statement, id).Scan(&user.Index, &user.Sessionid, &user.Id, &user.Password, &user.Name, &user.Phone, &user.Email, &user.CreateAt)
	if err != nil {
		return http.StatusUnprocessableEntity, user
	}

	if user.Password != password {
		return http.StatusUnprocessableEntity, user
	}
	fmt.Println(user)

	uuid := uuid.New().String()
	sql_statement = "UPDATE users SET sessionid=$1 WHERE id=$2"

	_, err = c.db.Exec(sql_statement, uuid, id)
	if err != nil {
		panic(err)
	}

	return http.StatusOK, user
}

func (c *connectDB) UpdateUser(id string, password string) bool {
	sql_statement := "UPDATE users SET password=$1 WHERE id=$2"

	_, err := c.db.Exec(sql_statement, password, id)

	return err == nil
}

func (c *connectDB) DeleteUser(id string) bool {
	sql_statement := "DELETE FROM users WHERE id=$1"

	_, err := c.db.Exec(sql_statement, id)

	return err == nil
}

func (c *connectDB) Close() {
	c.db.Close()
}
