### Instagram's private API for [Go](www.golang.org) language

---

### Installation

`go get -u github.com/gkiryaziev/go-instagram`

### Functions and description

`NewInstagram(userName, password string) (*instagram, error)` Create a new object of Instagram. Constructor takes a `user name` and `password` as arguments and call `Login()` method after a successful initialization.

`(this *instagram) Login() error` Login to Instagram server.

`(this *instagram) GetMediaLikers(mediaId string) (*MediaLikers, error)` Getting information about likers. Method takes the `id` of media as argument.

`(this *instagram) GetMedia(mediaId string) (*Media, error)` Getting information about media and comments. Method takes the `id` of media as argument.

`(this *instagram) GetRecentActivity() (*RecentActivity, error)` Getting recent activity.

`(this *instagram) SearchUsers(query string) (*SearchUsers, error)` Search user. Method takes `query` string as argument.

`(this *instagram) GetUserNameInfo(userNameId int64) (*UserNameInfo, error)` Getting information about user. Method takes the `id` of user as argument.

`(this *instagram) GetUserTags(userNameId int64) (*UserTags, error)` Getting user tags. Method takes the `id` of user as argument.

`(this *instagram) SearchTags(query string) (*SearchTags, error)` Search tag. Method takes `query` string as argument.

`(this *instagram) TagFeed(tag, maxId string) (*TagFeed, error)` Getting media with tag. Method takes the `tag` and `id` of the next page.
