package configh

type Rds struct {
	Host     string
	Password string
}

type RocketMQ struct {
	Hosts     []string
	AccessKey string
	SecretKey string
}

type Nsq struct {
	Host string
}

type Smtp struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Config struct {
	Project string
	Module  string

	Salt string

	MySQL string

	Rds Rds
	//RocketMQ RocketMQ

	Nsq Nsq

	Smtp Smtp
}
