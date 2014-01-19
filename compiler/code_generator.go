package compiler

import (
	"io"
	"github.com/strongo/templates"
)

type CodeGenerator struct {
	writer io.Writer
	environment *templates.StrongoEnvironment

	fileName string
	codeLineNumber int
	//the current indentation
	indentation int
}

func New()
