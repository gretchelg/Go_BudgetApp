package service

type Config struct {
	MongoURI string
	Plaid    struct {
		Secret string
		Env    string
	}
}
