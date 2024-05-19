package randh

var (
	Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52
	Symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                   // 32
	Digits  = "0123456789"                                           // 10

	AlphaNumeric = Letters + Digits       // 94
	Characters   = AlphaNumeric + Symbols // 94
)
