package regexph

import "regexp"

var (
	Regexp_ACCT   = "^[a-zA-Z0-9]{6,12}$"
	Regexp_PWD    = "^[a-zA-Z0-9]{32}$"
	Regexp_INV_CD = "^([a-zA-Z0-9]{6})?$"
	Regexp_TOKEN  = "^[a-zA-Z0-9]{64}$"
	Regexp_Eml    = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
)

func Match(regexStr, content string) bool {
	return regexp.MustCompile(regexStr).MatchString(content)
}

func Acct(str string) bool {
	return Match(Regexp_ACCT, str)
}

func Pwd(str string) bool {
	return Match(Regexp_PWD, str)
}

func InvCd(str string) bool {
	return Match(Regexp_INV_CD, str)
}

func Token(str string) bool {
	return Match(Regexp_TOKEN, str)
}

func Eml(str string) bool {
	return Match(Regexp_Eml, str)
}
