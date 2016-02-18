package main

import (
	"fmt"
	"log"
	"time"

	"./instagram_api"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// login_required
	instagram, err := instagram_api.NewInstagram("gkiryaziev", "InstagramAdm1n")
	checkError(err)

	start := time.Now()

	NextMaxID := ""
	mediaSum := 0
	likersSum := 0

	//loop := true

	for i := 0; i < 1; i++ {
		//for loop {

		feed, err := instagram.TagFeed("trendever", NextMaxID)
		checkError(err)

		for _, v := range feed.Items {
			fmt.Println("ID:", v.ID, "LikeCount:", v.LikeCount, "Username:", v.User.Username)
			for _, v := range v.ImageVersions2.Candidates {
				if v.Width == 240 && v.Height == 240 {
					fmt.Println("<img src=\"" + v.URL + "\">")
				}
			}

			for _, v := range v.Comments {
				fmt.Println("<img src=\""+v.User.ProfilePicURL+"\">", v.User.Username, v.Text)
			}

			likers, err := instagram.GetMediaLikers(v.ID)
			checkError(err)
			fmt.Println("\tUserCount:", likers.UserCount)
			for _, v := range likers.Users {
				fmt.Println("\t", v.Username)
			}
			likersSum += likers.UserCount
		}

		mediaSum += feed.NumResults
		//loop = feed.MoreAvailable
		NextMaxID = feed.NextMaxID
	}

	fmt.Println("------------------------")
	fmt.Println("Media Sum:", mediaSum)
	fmt.Println("Likers Sum:", likersSum)
	fmt.Println("Elapsed:", time.Since(start))

}
