package config

const configFileNam =".gatorconfig.json"

type Config struct {
  DBURL string `json:"db_url"`
  CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
  cfg.CurrentUserName = userName
  return write(*cfg)
}