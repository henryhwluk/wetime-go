package models

type User struct {
	UserID    string `json:"userID" bson:"userID"`
	Age       string `json:"age" bson:"age"`
	Sex       string `json:"sex" bson:"sex"`
	Password  string `json:"password" bson:"password"`
	AvatorURL string `json:"avatorURL" bson:"avatorURL"`
}

type Post struct {
	UserID    string `json:"userID" bson:"userID"`
	AvatorURL string `json:"avatorURL" bson:"avatorURL"`
	PostURL   string `json:"postURL" bson:"postURL"`
	PostTime  string `json:"postTime" bson:"postTime"`
}

type Follow struct {
	FollowerID string `json:"followerID" bson:"followerID"`
	UserID     string `json:"userID" bson:"userID"`
}
