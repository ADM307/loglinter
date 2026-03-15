package rules

import (
	"go/token"
	"golang.org/x/tools/go/analysis"
	"unicode"
)

func checkLowercase(msg string) (bool, string) {
	if len(msg) == 0 {
		return true, ""
	}
	firstSymb := rune(msg[0])
	if unicode.IsLetter(firstSymb) && !unicode.IsLower(firstSymb) {
		return false, "log message should start with a lowercase letter"
	}
	return true, ""
}

func CheckLowercase(pass *analysis.Pass, pos token.Pos, msg string) {
	if ok, errMsg := checkLowercase(msg); !ok {
		pass.Report(analysis.Diagnostic{
			Pos:     pos,
			Message: errMsg,
		})
	}
}
