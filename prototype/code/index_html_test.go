package code

import (
	"testing"
	"bytes"
	"github.com/strongo/templates/prototype/models"
)

func Test_Index_html(t *testing.T) {
	writer := new(bytes.Buffer)
	user := &models.User{Name: "John Smith"}
	payload := Payload_Index_html{
		BgColor: "white",
		User: user,
	}
	indexHtml := NewIndex_html("ru_RU", payload)
	indexHtml.Render(writer)
	s := writer.String()
	t.Log(s)
//	if s != "Hello, stranger!" {
//		t.Errorf("Unexpected output")
//	}
}

func Benchmark_Index_html(b *testing.B) {
	writer := new(bytes.Buffer)
	user := &models.User{Name: "John Smith"}
	payload := Payload_Index_html{
		BgColor: "white",
		User: user,
	}
	for i := 0; i < b.N; i++ {
		indexHtml := NewIndex_html("ru_RU", payload)
		indexHtml.Render(writer)
	}
}

