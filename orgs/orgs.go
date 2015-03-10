package orgs

import (
	"database/sql"
	"fmt"
	"os"

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
	params := ""
	if os.Getenv("UM_USER") != "" {
		params = params + "user=" + os.Getenv("UM_USER")
	}
	if os.Getenv("UM_DBNAME") != "" {
		params = params + " dbname= " + os.Getenv("UM_DBNAME")
	}
	params = params + " sslmode=disable"
	fmt.Printf("[INIT] %s\n", params)
	db, err = sql.Open("postgres", params)
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
