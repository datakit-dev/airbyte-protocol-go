package airbyte

// log level
type Level string

const (
	Debug      Level = "DEBUG"
	Fatal      Level = "FATAL"
	Info       Level = "INFO"
	LevelERROR Level = "ERROR"
	LevelTRACE Level = "TRACE"
	Warn       Level = "WARN"
)
