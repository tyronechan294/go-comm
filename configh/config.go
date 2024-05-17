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

type Config struct {
	Project string
	Module  string

	Salt string

	MySQL string

	Rds      Rds
	RocketMQ RocketMQ
}
