package rules

import (
	"go/token"
	"golang.org/x/tools/go/analysis"
	"unicode"
)

func checkSpecialSymbols(msg string) (bool, string) {
	for _, r := range msg {
		if isEmoji(r) {
			return false, "log message should not contain emoji"
		}
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			return false, "log message should not contain special characters"
		}
	}

	return true, ""
}

func isEmoji(r rune) bool {
	return (r >= 0x1F600 && r <= 0x1F64F) ||
		(r >= 0x1F300 && r <= 0x1F5FF) ||
		(r >= 0x1F680 && r <= 0x1F6FF) ||
		(r >= 0x1F700 && r <= 0x1F77F) ||
		(r >= 0x1F780 && r <= 0x1F7FF) ||
		(r >= 0x1F800 && r <= 0x1F8FF) ||
		(r >= 0x1F900 && r <= 0x1F9FF) ||
		(r >= 0x1FA00 && r <= 0x1FA6F) ||
		(r >= 0x1FA70 && r <= 0x1FAFF) ||
		(r >= 0x2600 && r <= 0x26FF) ||
		(r >= 0x2700 && r <= 0x27BF) ||
		(r >= 0xFE00 && r <= 0xFE0F) ||
		(r >= 0x1F1E6 && r <= 0x1F1FF)
}

func CheckSpecialSymbols(pass *analysis.Pass, pos token.Pos, msg string) {
	if ok, errMsg := checkSpecialSymbols(msg); !ok {
		pass.Report(analysis.Diagnostic{
			Pos:     pos,
			Message: errMsg,
		})
	}
}
