package orm

import "time"

// Config mysql config.
type Config struct {
	Dialect     string        // db name, like "mysql", "mssql", "postgres" or "sqlite"
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}
