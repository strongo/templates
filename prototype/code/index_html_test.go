package code

import (
	"testing"
	"bytes"
	"github.com/strongo/templates"
)


func GetIndexHtmlPayload() Payload_Index_html {
	return Payload_Index_html{
		BgColor: "white",
		AuthorIds: []int {101, 102, 103, 104, 105},
	}
}

func Test_Index_html(t *testing.T) {
	writer := new(bytes.Buffer)
	payload := GetIndexHtmlPayload()
	indexHtml := NewIndex_html("ru_RU", payload)
	indexHtml.GetData()
	indexHtml.Render(templates.RenderContext{Writer: writer})
	s := writer.String()
	t.Log(s)
//	if s != "Hello, stranger!" {
//		t.Errorf("Unexpected output")
//	}
}

func Benchmark_Index_html(b *testing.B) {
	writer := new(bytes.Buffer)
	payload := GetIndexHtmlPayload()
	for i := 0; i < b.N; i++ {
		indexHtml := NewIndex_html("ru_RU", payload)
		indexHtml.GetData()
		indexHtml.Render(templates.RenderContext{Writer: writer})
	}
}
