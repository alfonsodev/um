package usermanager

import (
	"fmt"
	"github.com/alfonsodev/sqlh"
	"testing"
)

func TestSave(t *testing.T) {

	// server := &Server{
	// 	Name:    "gopher",
	// 	ID:      123456,
	// 	Enabled: true,
	// }

	// var user User
	// user.Username = "pepe"
	// fmt.Printf("%s\n", user.Username)
	// user.Save()

	var user User
	user.Googleid.String = "alskfjalsdf"
	user.Person.String = "{}"

	k, p, v := sqlh.StructListKeys(&user)
	s := "INSERT INTO usermanager.users(" + k + ") VALUES (" + p + ")"
	fmt.Println(s)
	_, err := db.Query(s, v...)
	if err != nil {
		fmt.Println(">>>>>>>>" + err.Error())
	}
}
