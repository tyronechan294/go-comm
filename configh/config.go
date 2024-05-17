package configh

type Transfer struct {
	FeeLimit int64
}

type Wallet struct {
	Path string
}

type Collect struct {
	Addr string
	Path string
}

type Tron struct {
	Host   string
	APIKey string

	Transfer Transfer

	Wallet Wallet

	Collect Collect
}

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
	App     string

	Salt string

	MySQL string

	Rds      Rds
	RocketMQ RocketMQ

	Tron Tron
}
