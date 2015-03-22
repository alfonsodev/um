package usermanager

import (
	"encoding/json"

	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/google-api-go-client/plus/v1"
	"database/sql"
	"fmt"
	googleAuth "github.com/alfonsodev/googleauth"
	"github.com/alfonsodev/sqlh"
	_ "github.com/lib/pq"
	"os"
	"reflect"
	"strings"
)

var db *sql.DB

type User struct {
	Id          sql.NullInt64
	Username    sql.NullString
	Fullname    sql.NullString
	Email       sql.NullString
	Location    sql.NullString
	Googleid    sql.NullString
	Googletoken sql.NullString
	Person      sql.NullString
}

var (
	CLIENT_ID     = os.Getenv("GOOGLE_CLIENT_ID")
	CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
	SCOPE         = os.Getenv("GOOGLE_SCOPE")
	REDIRECT      = os.Getenv("GOOGLE_REDIRECT")
)
var config = &oauth.Config{
	ClientId:     CLIENT_ID,
	ClientSecret: CLIENT_SECRET,
	// Scope determines which API calls you are authorized to make
	Scope:    SCOPE,
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
	//Use "postmessage" for the code-flow for server side apps
	RedirectURL: REDIRECT,
}

func StrutForScan(u interface{}) []interface{} {
	val := reflect.ValueOf(u).Elem()
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		//		if valueField.Type().Name() == "string" && valueField.String() == "" {

		v[i] = valueField.Addr().Interface()
	}
	return v
}

func NewUserModel() User {
	um := User{}
	um.Person.String = "{}"
	return um
}

func (um *User) Save() bool {
	var query string

	fmt.Printf("\nvalues:%v\n", um.Id)
	if um.Id.Int64 > 0 {
		names, values := sqlh.StructToKeyValue(um)
		query = "UPDATE myschema.Users SET "
		query += names
		query += " WHERE myschema.Users = $1"
		fmt.Printf("\nQuery:%s\n", query)
		fmt.Printf("\nvalues:%s\n", values)
		//		_, err := db.Query(query, values)
		//		PanicIf(err)
	} else {
		keys, placeholders, values := sqlh.StructListKeys(um)
		query = "INSERT INTO usermanager.users (" + keys + ") "
		query += " VALUES(" + placeholders + ") "
		fmt.Printf("\nQuery:%s\n", query)
		fmt.Printf("\nValues:%s\n", values)
		_, err := db.Query(query, values...)
		PanicIf(err)
	}
	//	defer rows.Close()
	return true
}

func init() {
	var err error
	params := ""
	if os.Getenv("UM_USER") != "" {
		params = params + "user=" + os.Getenv("UM_USER")
	}
	if os.Getenv("UM_DBNAME") != "" {
		params = params + " dbname= " + os.Getenv("UM_DBNAME")
		params = params + " dbname=foo "
	} else {
	}
	params = params + " sslmode=disable"
	fmt.Printf("[INIT] %s\n", params)
	db, err = sql.Open("postgres", params)
	PanicIf(err)

	fmt.Printf("[INIT] usermanager\n")
}

func PanicIf(err error) {
	if err != nil {
		fmt.Println("[Err] " + err.Error())
		panic(err)
	}
}

func List() {
	rows, err := db.Query("SELECT * FROM usermanager.users")
	PanicIf(err)
	defer rows.Close()
	var user User
	fmt.Println("User list")
	fmt.Println("-----------------")
	fmt.Println("ID     | Username | googleid")
	fmt.Println("-----------------")

	for rows.Next() {
		err := rows.Scan(StrutForScan(&user)...)
		// err := rows.Scan(&user.username)
		fmt.Printf("%6d | %s | %s \n", user.Id.Int64, string(user.Username.String), string(user.Googleid.String))
		PanicIf(err)
	}

}

func GetUserByUsername(uname string) (User, error) {
	row, err := db.Query("SELECT * FROM usermanager.users WHERE usermanager.users.username = $1 LIMIT 1", uname)

	PanicIf(err)
	defer row.Close()

	var user User
	for row.Next() {
		err := row.Scan(StrutForScan(&user)...)

		//err := rows.Scan(&user.username)
		//		fmt.Printf("\nUsername: %v: %s", user.id, string(user.username.String))
		PanicIf(err)
	}

	return user, err
}

func GetByGoogleId(gid string) (User, error) {
	user := NewUserModel()
	q := "SELECT * FROM usermanager.users WHERE usermanager.users.googleid = $1 LIMIT 1"
	fmt.Println(q)
	fmt.Println(gid)
	row, err := db.Query(q, gid)
	PanicIf(err)
	row.Next()

	errRow := row.Scan(StrutForScan(&user))

	return user, errRow
}

func GoogleAuthLogic(code string) string {
	accessToken, idToken, err := googleAuth.Exchange(code)
	gplusID, err := googleAuth.DecodeIdToken(idToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n ADMINS:%s=%s\n", os.Getenv("ADMINS"), gplusID)
	if !strings.Contains(os.Getenv("ADMINS"), gplusID) {
		return ""
	} else {
		user, _ := GetByGoogleId(gplusID)
		if user.Id.Int64 == 0 {
			person := getPersonFromToken(accessToken)
			user := NewUserModel()
			jsonPerson, _ := json.Marshal(person)
			user.Person.String = string(jsonPerson)
			user.Save()
		}
		return gplusID
	}
}

func getPersonFromToken(token string) (person *plus.Person) {
	// Create a new authorized API client
	t := &oauth.Transport{Config: config}
	tok := new(oauth.Token)
	tok.AccessToken = token
	t.Token = tok
	service, err := plus.New(t.Client())
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
	// Get a list of people that this user has shared with this app
	people := service.People.Get("me")
	person, err = people.Do()
	//TODO:Handle all this posible errors
	if err != nil {
		if err.Error() == "AccessTokenRefreshError" {
			fmt.Printf("\n err: %s", err)
			return // &appError{errors.New(m), m, 500}
		}
		fmt.Printf("\n err: %s", err)
		return // &appError{err, m, 500}
	}

	return person
}
