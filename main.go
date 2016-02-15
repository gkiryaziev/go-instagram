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

	//tags, err := instagram.GetUserTags(1934125366)
	//checkError(err)
	//fmt.Println(tags)

	//tags, err := instagram.SearchTags("trendever")
	//checkError(err)
	//for _, t := range tags.Results {
	//	if t.Name == "trendever" {
	//		fmt.Println("Name:", t.Name)
	//		fmt.Println("ID:", t.ID)
	//		fmt.Println("Media Count:", t.MediaCount)
	//	}
	//}

	// 1185555018460539520_298705361 70
	// 1185528481900579684_1734629785 106

	feed, err := instagram.TagFeed("trendever", "")
	checkError(err)

	fmt.Println("NumResults:", feed.NumResults)

	for _, v := range feed.Items {
		fmt.Println("ID:", v.ID, "LikeCount:", v.LikeCount, "Username:", v.User.Username)
	}

	fmt.Println("NextMaxID:", feed.NextMaxID)

	//comment, err := instagram.GetMediaComments("1185528481900579684_1734629785")
	//fmt.Println(comment.Caption.Text)
	//likers, err := instagram.GetMediaLikers("1185528481900579684_1734629785")
	//fmt.Println("UserCount", likers.UserCount)
	//for _, v := range likers.Users {
	//	fmt.Println(v.Username)
	//}

}
