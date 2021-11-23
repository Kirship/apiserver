package apiserver

type config struct {
	BindAddr string `toml:"bindAddr"`
	loglevel string `toml:"loglevel"`
}

func NewConfig() *config {
	return &config{
		BindAddr: ":8080",
		loglevel: "debug",
	}
}
