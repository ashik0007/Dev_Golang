package main

import "fmt"

func main() {
	ub := &UserBuilder{}
	user := ub.
		Number(15).     // Giving a number
		Add(7).	       // Returns after addition operation on the number
		Remove(2).    // Returns after total operation on the number
		Build()	
	
	fmt.Println(user.Number)     // Returns the number
	fmt.Println(user.Add)       // Returns the result after the add operation
	fmt.Println(user.Remove)     // Returns the result after total operation
}

type User struct {
	Number int
	Add int
	Remove int
}

type UserBuilder struct {
	User
}

func (ub *UserBuilder) Build() User {
	return ub.User
}

// Returns the number
func (ub *UserBuilder) Number(number int) *UserBuilder {
	ub.User.Number = number
	return ub
}

//Returns the addition output
func (ub *UserBuilder) Add(add int) *UserBuilder {
	ub.User.Add = ub.User.Number + add
	return ub
}

//Returns the subtraction output
func (ub *UserBuilder) Remove(remove int) *UserBuilder {
	ub.User.Remove = ub.User.Add - remove
	return ub
}