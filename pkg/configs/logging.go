package configs

type LoggingConfiguration struct {
	Level     string `default:"debug"`
	Timestamp bool   `default:"true"`
}
