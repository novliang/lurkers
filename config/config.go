package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Log     Log     `mapstructure:"log" json:"log" yaml:"log"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type System struct {
	Addr int `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type JWT struct {
	SigningKey    string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
	DefaultExpire int    `mapstructure:"default-expire" json:"defaultExpire" yaml:"default-expire"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

type Captcha struct {
	Sign   string `mapstructure:"sign" json:"sign" yaml:"sign"`
	Expire int    `mapstructure:"expire" json:"expire" yaml:"expire"`
}
