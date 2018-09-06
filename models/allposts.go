package models

// AllPosts store all posts information. key: file name, value : TPost.
var AllPosts map[string]TPost

func LoadAllPosts() {
	//postDir := beego.AppConfig.String("PostDirectory")

}

func init() {

}
