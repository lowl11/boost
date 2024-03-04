package swagger

var (
	_cfg *Config
)

type Config struct {
	Title                    string
	URL                      string
	Oauth2RedirectURL        string
	Oauth2DefaultClientID    string
	PersistAuthorization     string
	DocExpansion             string
	DeepLinking              bool
	DefaultModelsExpandDepth int
}

func SetConfig(cfg Config) {
	_cfg = &cfg
}

func defaultConfig() Config {
	return Config{
		Title: "Swagger Boost API",
		URL:   "docs/swagger.json",
	}
}

func getConfig() *Config {
	if _cfg == nil {
		tmp := defaultConfig()
		_cfg = &tmp
	}
	return _cfg
}
