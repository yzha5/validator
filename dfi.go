package validator

type Result struct {
	Err error
	Msg map[string][]Rule
}

type Rule string

const (
	NOTNULL Rule = "NOTNULL"
	EMAIL   Rule = "EMAIL"
	PHONE   Rule = "PHONE"
	GT      Rule = "GT"
	GTE     Rule = "GTE"
	LT      Rule = "LT"
	LTE     Rule = "LTE"
	EQ      Rule = "EQ"
	NEQ     Rule = "NEQ"
	LIKE    Rule = "LIKE"
	BETW    Rule = "BETW"
	NBETW   Rule = "NBETW"
	REG     Rule = "REG"
	LOW     Rule = "LOW"
	CAP     Rule = "CAP"
	LETTER  Rule = "LETTER"
	NUMERIC Rule = "NUMERIC"
	AN      Rule = "AN"
	ANP     Rule = "ANP"
	URL     Rule = "URL"
	LEN     Rule = "LEN"
	MINLEN  Rule = "MINLEN"
	MAXLEN  Rule = "MAXLEN"
	MIN     Rule = "MIN"
	MAX     Rule = "MAX"
	ONEOF   Rule = "ONEOF"
)
