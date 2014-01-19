package templates

import (
	"io"
)

type StrongoEnvironment struct {

}

type Template interface {
	Render(writer io.Writer, context interface{}) error
}
