package rules

import (
	"go/token"
	"golang.org/x/tools/go/analysis"
)

type Rule func(pass *analysis.Pass, pos token.Pos, msg string)

func All() []Rule {
	return []Rule{
		CheckInEnglish,
		CheckLowercase,
		CheckSensitiveData,
		CheckSpecialSymbols,
	}
}
