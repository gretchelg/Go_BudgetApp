package models

// User defines an app user
// the `bson` tag is useful for mapping to bson-encoded DB systems
type User struct {
	Email       string `json:"email" bson:"email"`
	AccessToken string `json:"access_token" bson:"access_token"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	// TODO add more
}
