package rules

import (
	"go/token"
	"golang.org/x/tools/go/analysis"
	"strings"
)

var sensitiveKeywords = []string{
	"password",
	"token",
	"secret",
	"key",
}

func checkSensitiveData(msg string) (bool, string) {
	lowered := strings.ToLower(msg)
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lowered, keyword) {
			return false, "log message may contain sensitive data: \"" + keyword + "\""
		}
	}
	return true, ""
}

func CheckSensitiveData(pass *analysis.Pass, pos token.Pos, msg string) {
	if ok, errMsg := checkSensitiveData(msg); !ok {
		pass.Report(analysis.Diagnostic{
			Pos:     pos,
			Message: errMsg,
		})
	}
}
