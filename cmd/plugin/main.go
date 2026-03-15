package main

import (
	"golang.org/x/tools/go/analysis"
	"github.com/ADM307/loglinter/pkg/loglinter"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{loglinter.Analyzer}, nil
}
