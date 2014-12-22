package templates

import (
	"io"
//	"log"
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
	GetData()
	Render(c RenderContext) error
}

type I18n interface {
	GetText(key string) string
}

type Block interface {

}

type RenderContext struct {
	Writer io.Writer  // current writer

}

func (c RenderContext) WriteString(s string) (n int, err error) {
	return io.WriteString(c.Writer, s)
}

type RenderFuture interface {
	Render(c RenderContext)
}

type StrongTask struct {

}

type IStrongoComponent interface {
	GetData()
}

type StrongoComponent struct {
	semaphore chan int
	components []IStrongoComponent
}

func (c *StrongoComponent) OnDataReady(){
//	log.Print("OnDataReady 1")
	c.semaphore <- 1
//	log.Print("OnDataReady 2")
}

func (c *StrongoComponent) WhenDataReady(){
//	log.Print("WhenDataReady 1")
	<-c.semaphore
//	log.Print("WhenDataReady 2")
}

func NewStrongoComponent(components []IStrongoComponent) *StrongoComponent {
	return &StrongoComponent{
		semaphore: make(chan int, 1),
		components: components,
	}
}

func (self StrongoComponent) Components() []IStrongoComponent {
	return self.components
}

func (c StrongoComponent) GetData() {
	if c.components != nil{
		for _, component := range c.components {
			component.GetData()
		}
	}
}
