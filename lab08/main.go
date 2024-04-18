package main

import (
	"log"
)
type User struct {
 id int
 name string
 init bool
}
func (u *User) define(id int, n string) {
 u.id = id
 u.name = n
 u.init = true
}
func define(u *User, id int, n string) {
 u.id = id
 u.name = n
 u.init = true
}
func NewUser(id int, n string) User {
 return User{
 id:id,
 name:n,
 init:true,
 }
}
func main() {
 user := User{}
 define(&user,2,"Mateusz")
 user2 := NewUser(2,"Dawid")
 user2.define(3,"Piotao")
 // _user2 := NewUser(1,"Dawid")

//  user3 := &NewUser(4,"Zupa")
//  user3.define(4,"Kasza")
//  log.Println(user3)
 log.Println(user2)
 if user.init {
 log.Println("user ok")
 }else{
 log.Fatal("user not initialized")
 }
}
