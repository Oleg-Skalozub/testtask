package config

// FilePath ...
var FilePath = ""

// Configuration ...
type Configuration struct {
	DBDialect  string `json:"db_dialect"`
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBName     string `json:"db_name"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	ApiRoute   string `json:"api_route"`
}
