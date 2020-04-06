package log

// Config config for logger.
type Config struct {
	Filepath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}
