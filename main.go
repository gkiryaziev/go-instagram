package main

import (
	//"fmt"
	"log"

	"./libs"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//1934125366 makeup_manikur
//2940280567 gkiryaziev

func main() {

	instagram := libs.NewInstagram("gkiryaziev", "instagramAdm1n")

	err := instagram.Login()
	checkError(err)

	//users, err := instagram.SearchUsers("gkiryaziev")
	//checkError(err)
	//
	//fmt.Println("Status:", users.Status)
	//fmt.Println("Count:", users.NumResults)
	//for _, user := range users.Users {
	//	if user.Username == "gkiryaziev" {
	//		fmt.Println("\tPK:", user.Pk)
	//		fmt.Println("\tUser Name:", user.Username)
	//		fmt.Println("\tPicture:", user.ProfilePicURL)
	//		fmt.Println()
	//	}
	//}

	//info, err := instagram.GetUserNameInfo(1934125366)
	//fmt.Println(info.User.MediaCount)

	instagram.GetUserTags(1934125366)
	//fmt.Println(activ)
}
