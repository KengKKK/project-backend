package config

// ModelConnection ;
type ModelConnection struct {
	User   string
	Pass   string
	Host   string
	DBname string
}

// GetConfig ;
func GetConfig() ModelConnection {
	return ModelConnection{"root", "123456", "35.240.137.204:3306", "nextcorp"}
}
