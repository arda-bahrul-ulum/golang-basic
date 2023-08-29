package main

import "fmt"

type User struct {
	ID int
	FirstName string
	LastName string
	Email string
	isActive bool
}

type Group struct {
	Name string
	Admin User
	Users []User
	isAvailable bool
}

func main()  {
	// cara 1
	// user := User{}
	// user.ID = 1
	// user.FirstName = "Jhony"
	// user.LastName = "Jack"
	// user.Email = "jhony@mal.com"
	// user.isActive = true

	// cara 2
	// user2 := User{   
	// 	ID : 1,
	// 	FirstName : "Jacky",
	// 	LastName : "Chan",
	// 	Email : "chan@mal.com",
	// 	isActive : true,
	// }

	user3 := User{1, "Zelda", "Safitri", "zelda@mail.com", true}
	user4 := User{1, "Gogon", "melisa", "gogon@mail.com", true}

	// displayUser1 := displayUser(user3)
	// displayUser2 := displayUser(user4)

	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)

	users := []User{user3, user4}

	group := Group{"Gamer", user3, users, true}

	// call method group
	group.DisplayGroup()
	// displayGroup(group)

	// call method user
	// result := user3.display()
	// fmt.Println(result)

	// fmt.Println(user4.display())
}

func (group Group) DisplayGroup()  {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func displayGroup(group Group)  {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

// method
func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}