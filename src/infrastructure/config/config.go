package config

// FilePath ...
var FilePath = ""

// Configuration ...
type Configuration struct {
	DBConfig   DBConfig  `json:"bd_config"`
	ApiRoute   string    `json:"api_route"`
	ServerPort int       `json:"server_port"`
	LogConfig  LogConfig `json:"logger_config"`
}

// DBConfig ...
type DBConfig struct {
	DBDialect  string `json:"db_dialect"`
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBName     string `json:"db_name"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBConn     int    `json:"db_conn"`
	DBIdleConn int    `json:"db_idle_conn"`
}

// LogConfig ...
type LogConfig struct {
	Prefix   string `json:"prefix"`
	FileName string `json:"file_name"`
}
