package code

import (
	"io"
	"github.com/strongo/templates"
)

type Layout_html interface {
	RenderBlock_head_title(writer io.Writer) error
	RenderBlock_page_title(writer io.Writer) error
	RenderBlock_menu(writer io.Writer) error
	RenderBlock_content(writer io.Writer) error
}

type Payload_Layout_html struct {
	BgColor string
}

type layout_html_struct struct {
	i18n templates.I18n
	template Layout_html
	payload Payload_Layout_html
}

func New_layout_html_struct(i18n templates.I18n, template Layout_html, payload Payload_Layout_html) layout_html_struct {
	return layout_html_struct{
		i18n: i18n,
		template: template,
		payload: payload,
	}
}


func (layout layout_html_struct) Render(writer io.Writer, payload Payload_Layout_html) error {
	if _, err := writer.Write([]byte("<!DOCTYPE html>\n<html>\n<head>\n<title>")); err != nil {
		return err
	}
	if err := layout.template.RenderBlock_head_title(writer); err != nil {
		return err
	}
	writer.Write([]byte("</title>\n</head>\n<body style=\"background-color: "))
	writer.Write([]byte(payload.BgColor))
	writer.Write([]byte(";\">\n<h1>"))
	if _, err := writer.Write([]byte(layout.i18n.GetText("Welcome to page"))); err != nil {
		return err
	}
	writer.Write([]byte(" "))
	if err := layout.template.RenderBlock_page_title(writer); err != nil {
		return err
	}
	writer.Write([]byte("\n<h1>\n"))
	if err := layout.template.RenderBlock_menu(writer); err != nil {
		return err
	}
	writer.Write([]byte("\n<hr/>\n"))
	if err := layout.template.RenderBlock_content(writer); err != nil {
		return err;
	}
	writer.Write([]byte("\n</body>\n</html>")) // Lines 10:11
	return nil
}

func (t layout_html_struct) RenderBlock_head_title(writer io.Writer) error {
	_, err := writer.Write([]byte("{BLOCK head_title}"))
	return err
}

func (t layout_html_struct) RenderBlock_page_title(writer io.Writer) error {
	_, err := writer.Write([]byte("{BLOCK page_title}"))
	return err
}

func (t layout_html_struct) RenderBlock_menu(writer io.Writer) error {
	_, err := writer.Write([]byte("{BLOCK menu}"))
	return err
}

func (t layout_html_struct) RenderBlock_content(writer io.Writer) error {
	_, err := writer.Write([]byte("{BLOCK content}"))
	return err
}
