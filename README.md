### Instagram's private API for [Go](www.golang.org) language
---

[![Go Report Card](https://goreportcard.com/badge/github.com/gkiryaziev/go-instagram)](https://goreportcard.com/report/github.com/gkiryaziev/go-instagram)

### Installation

`go get -u github.com/gkiryaziev/go-instagram`

### Functions and description

`NewInstagram(userName, password string) (*instagram, error)` Create a new object of Instagram. Constructor takes a `user name` and `password` as arguments and call `Login()` method after a successful initialization.

`Login() error` Login to Instagram server.

`GetMediaLikers(mediaId string) (*MediaLikers, error)` Getting information about likers. Method takes the `id` of media as argument.

`GetMedia(mediaId string) (*Media, error)` Getting information about media and comments. Method takes the `id` of media as argument.

`GetRecentActivity() (*RecentActivity, error)` Getting recent activity.

`SearchUsers(query string) (*SearchUsers, error)` Search user. Method takes `query` string as argument.

`GetUserNameInfo(userNameId int64) (*UserNameInfo, error)` Getting information about user. Method takes the `id` of user as argument.

`GetUserTags(userNameId int64) (*UserTags, error)` Getting user tags. Method takes the `id` of user as argument.

`SearchTags(query string) (*SearchTags, error)` Search tag. Method takes `query` string as argument.

`TagFeed(tag, maxId string) (*TagFeed, error)` Getting media with tag. Method takes the `tag` and `id` of the next page.
