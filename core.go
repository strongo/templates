package templates

import (
	"io"
	"github.com/strongo/templates/text"
)

type StrongoEnvironment struct {
	Funcs text.FuncMap
}

// Defines mapping from template function name to real name.
// Usage: environment.Funcs(map[string]StrongoFunc)
type StrongoFunc struct {
	Package string
	Name string
}

type Template interface {
	Render(writer io.Writer) error
}

type I18n interface {
	GetText(key string) string
}

type Block interface {

}


type RenderTaskInterface interface {
	Write(writer io.Writer)
}

type RenderTask struct {
	Id int
}

type RenderContext struct {
	Writer io.Writer  // current writer
}

type Renderer struct {
	FinalWriter io.Writer
	CurrentContext RenderContext
	renderTasks map[int] RenderTaskInterface
}
