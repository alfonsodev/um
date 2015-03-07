package usermanager

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	id       int64
	username []byte
	fullname []byte
	email    []byte
	location []byte
	person   []byte
}

func init() {
	var err error
	db, err = sql.Open("postgres", "dbname=foo sslmode=disable")
	PanicIf(err)

	fmt.Printf("[INIT] usermanager\n")
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func List() {
	rows, err := db.Query("SELECT * FROM usermanager.users")

	PanicIf(err)
	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(&user.id, &user.username, &user.fullname, &user.email, &user.location, &user.person)
		//err := rows.Scan(&user.username)
		fmt.Printf("\nUsername: %v: %s", user.id, string(user.username))
		PanicIf(err)
	}

}
