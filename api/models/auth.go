package models

// UserUid is the struct for the user uid.
type UserUid struct {
	Uid      string `json:"uid" firestore:"id"`
}

// UserToken is the struct for the user JWT token.
type UserToken struct {
	Token      string `json:"token" firestore:"id"`
}