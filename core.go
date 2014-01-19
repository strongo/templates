package templates

import (
	"io"
)

type Template interface {
	Render(writer io.Writer, context interface{}) error
}
