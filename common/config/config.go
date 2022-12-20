package config

type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
}

type Server struct {
	Post string `yaml:"post"`
}

type Mysql struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}
