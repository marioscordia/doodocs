package config

// type Config struct {
// 	Email string `json:"email"`
// 	Password string `json:"password"`
// }

// func NewConfig() (*Config, error) {
// 	file, err := os.Open("./config/config.json")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var config *Config
// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&config); err != nil {
// 		return nil, err
// 	}

// 	return config, nil
// }

// func (c *Config) SetVars() error{
// 	if err := os.Setenv("EMAIL", c.Email); err != nil {
// 		return err
// 	}

// 	if err := os.Setenv("PASSWORD", c.Password); err != nil{
// 		return err
// 	}

// 	return nil
// }