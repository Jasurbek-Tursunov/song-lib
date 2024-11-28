package postgres

type Config struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  bool
}
