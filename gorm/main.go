package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

func main() {
	db, err := gorm.Open(
		mysql.Open("root:root@tcp(127.0.0.1:3306)/testeGolang?parseTime=true"),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Username: "mathesukk", Email: "mathesu@mail.com", Password: "mypassword"})
	db.Create(
		&User{
			Username: "mathesukkkkkkkkkkkkk",
			Email:    "mathesukkkk@mail.com",
			Password: "mypassword",
		},
	)

	var user User
	var user2 User
	db.First(&user, "username = ?", "mathesukk")
	db.First(&user2, "email = ?", "mathesukkkk@mail.com")

	db.Model(&user).Update("Password", "mypasswordisgreatthankyou")
	db.Model(&user2).
		Updates(map[string]string{"Email": "mathesukk@gmail.com", "username": "maathesukkj"})

	fmt.Println(user)
	fmt.Println(user2)

	db.Delete(&user2) // as the model was created with gorm.Model, it has the DeletedAt field.
	// this field causes the delete to trigger a soft-delete, only changing its deleted_at date, and disabling it from being findable with normal queries

	// to perma delete it, we can do this:
	db.Unscoped().Delete(&user2)

	// raw queries
	var result User
	db.Raw("select * from users where id = ?", 2).Scan(&result)

	fmt.Println(result)

	db.Exec("delete from users")
}
