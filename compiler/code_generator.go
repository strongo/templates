package compiler2

import (
	"io"
//	"bytes"
	"os"
	"path/filepath"
	"github.com/strongo/templates"
	"github.com/strongo/templates/text"
	"fmt"
)

type CodeGenerator struct {
//	writer io.Writer
	environment templates.StrongoEnvironment

	fileName string
	codeLineNumber int
	//the current indentation
	indentation int
}

func (CodeGenerator) WriteLine(s string){

}

func NewCodeGenerator(e templates.StrongoEnvironment) *CodeGenerator {
	g := new(CodeGenerator);
	g.environment = e
//	g.writer = new(bytes.Buffer)
	return g
}

func (g *CodeGenerator) CompileDir(path string) error {
	return filepath.Walk(path, g.FileSystemWalker)
}

func (g *CodeGenerator) FileSystemWalker(path string, info os.FileInfo, err error) error {
	if err == nil {
		if info.IsDir() {
			fmt.Println("")
			fmt.Printf("%s/%s", path, info.Name())
			fmt.Println("")
		} else {
			fmt.Printf("\t%s", info.Name())
			t1 := text.New("t1")
			t1.ParseName
			fmt.Println("")
		}
	} else {
		fmt.Printf("ERROR: %s", err)
		fmt.Println("")
	}
	return nil
}

type TemplateToCodeCompiler interface {

}

type TemplateToGoCodeCompiler struct {
	g *CodeGenerator
}

func (c *TemplateToGoCodeCompiler) Compile(t text.Template, writer io.Writer) error {
	return nil
}
