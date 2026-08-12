package config

type DataBase struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	UserName string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
}

func (d *DataBase) GetDsn() string {
	return d.UserName + ":" + d.Password + "@tcp(" + d.Host + ":" + string(d.Port) + ")/"
}
