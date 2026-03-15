package rules

import (
	"go/token"
	"golang.org/x/tools/go/analysis"
	"unicode"
)

func checkInEnglish(msg string) (bool, string) {
	for _, r := range msg {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return false, "log message should use only English letters"
		}
	}
	return true, ""
}

func CheckInEnglish(pass *analysis.Pass, pos token.Pos, msg string) {
	if ok, errMsg := checkInEnglish(msg); !ok {
		pass.Report(analysis.Diagnostic{
			Pos:     pos,
			Message: errMsg,
		})
	}
}
