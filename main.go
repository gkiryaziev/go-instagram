package main

import (
	"fmt"
	"log"

	"./libs"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	instagram := libs.NewInstagram("gkiryaziev", "instagramAdm1n")

	err := instagram.Login()
	checkError(err)

	users, err := instagram.SearchUsers("sexybeatch")
	checkError(err)

	fmt.Println("Status:", users.Status)
	fmt.Println("Count:", users.NumResults)
	for _, user := range users.Users {
		fmt.Println("\tPK:", user.Pk)
		fmt.Println("\tUser Name:", user.Username)
		fmt.Println("\tPicture:", user.ProfilePicURL)
		fmt.Println()
	}
}
