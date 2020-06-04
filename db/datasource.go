package db

import "fmt"

type Datasource struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SSLMode  bool
}

func (ds Datasource) String() string {
	sslMode := "require"
	if !ds.SSLMode {
		sslMode = "disable"
	}

	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		ds.Host, ds.Port, ds.DBName, ds.User, ds.Password, sslMode,
	)
}
