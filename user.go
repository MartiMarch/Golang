package main

import (
	"fmt"
)

type User struct {
	id    int
	name  string
	alias string
	print func(User)
}

func user_crear(id int, name string, alias string) User {
	var user = User{
		id:    id,
		name:  name,
		alias: alias,
		print: user_print,
	}
	return user
}

func user_print(user User) {
	fmt.Println("Datos de usuario", user.name)
	fmt.Print("  id: '", user.id, "',")
	fmt.Print(" name: '", user.name, "',")
	fmt.Print(" alias: '", user.alias, "'\n")
}
