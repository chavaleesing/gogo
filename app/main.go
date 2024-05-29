package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Id        int    `db:"id"`
	Username  string `db:"username"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

var db *sqlx.DB

func main() {
	//test_pointer()
	println("-------------------------------------")

	var err error
	db, err = sqlx.Connect("postgres", "user=postgres_user password=password dbname=postgres_db sslmode=disable")

	if err != nil {
		println("connection has been failed")
		fmt.Println(err)
	}

	err = AddUser(User{20, "xx@gmail.com", "xx", "xxx"})
	if err != nil {
		println("Failed to add user")
	}

	UpdateUser(User{10, "sing@gmail.com", "UPDATEDDDDD", "song"})

	GetUsersX()
	u, _ := GetUserX(10)
	fmt.Println(*u)

	users, err := GetUsers()
	if err != nil {
		println("Failed to get user")
	}
	for _, u := range users {
		fmt.Println(u)
	}

	uu, _ := GetUserById(2)
	fmt.Println(*uu)

	println("-------------------------------------")
}

func GetUsersX() ([]User, error) {
	q := "select id, username, first_name, last_name from public.user"
	users := []User{}
	err := db.Select(&users, q)
	if err != nil {
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}
func GetUserX(id int) (*User, error) {
	q := "select id, username, first_name, last_name from public.user where id = $1"
	user := User{}
	err := db.Get(&user, q, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return &user, nil
}

func GetUsers() ([]User, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, username, first_name, last_name from public.user"
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	fmt.Printf("%#v", users)

	return users, nil
}

func GetUserById(id int) (*User, error) {
	row := db.QueryRowx("select id, username, first_name, last_name from public.user where id = $1", id)

	user := User{}
	row.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName)

	return &user, nil
}

func AddUser(user User) error {
	txn, _ := db.Begin()
	query := "insert into public.user (id, username, first_name, last_name)  values ($1, $2, $3, $4)"
	result, err := txn.Exec(query, user.Id, user.Username, user.FirstName, user.LastName)
	if err != nil {
		txn.Rollback()
		fmt.Println(err)
		return err
	}

	affected, err := result.RowsAffected()
	if affected <= 1 {
		txn.Rollback()
		return err
	}
	txn.Commit()
	return nil
}

func UpdateUser(user User) error {
	query := "update public.user set first_name = $1 where id = $2"
	result, err := db.Exec(query, user.FirstName, user.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	affected, err := result.RowsAffected()
	if affected <= 0 {
		return err
	}
	return nil
}

func test_pointer() {
	st1 := "test jaa"
	println(st1)
	set_value_format(&st1)
	println(st1)
	var pt *string
	pt = &st1
	pt2 := &st1
	println(pt, pt2)
	println(*pt)
}

func set_value_format(input *string) {
	*input = "I'm new format of " + *input
}
