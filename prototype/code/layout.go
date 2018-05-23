package code

import (
	"github.com/strongo/templates"
)

type Layout_html interface {
	RenderBlock_head_title(c templates.RenderContext) error
	RenderBlock_page_title(c templates.RenderContext) error
	RenderBlock_menu(c templates.RenderContext) error
	RenderBlock_content(c templates.RenderContext) error
}

type Payload_Layout_html struct {
	BgColor string
}

type layout_html struct {
	i18n templates.I18n
	template Layout_html
	payload Payload_Layout_html
}

func New_layout_html(i18n templates.I18n, template Layout_html, payload Payload_Layout_html) layout_html {
	return layout_html{
		i18n: i18n,
		template: template,
		payload: payload,
	}
}


func (layout layout_html) Render(c templates.RenderContext, payload Payload_Layout_html) error {
	if _, err := c.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<title>"); err != nil {
		return err
	}
	if err := layout.template.RenderBlock_head_title(c); err != nil {
		return err
	}
	c.WriteString("</title>\n</head>\n<body style=\"background-color: ")
	c.WriteString(payload.BgColor)
	c.WriteString(";\">\n<h1>")
	if _, err := c.WriteString(layout.i18n.GetText("Welcome to page")); err != nil {
		return err
	}
	c.WriteString(" ")
	if err := layout.template.RenderBlock_page_title(c); err != nil {
		return err
	}
	c.WriteString("\n<h1>\n")
	if err := layout.template.RenderBlock_menu(c); err != nil {
		return err
	}
	c.WriteString("\n<hr/>\n")
	if err := layout.template.RenderBlock_content(c); err != nil {
		return err;
	}
	c.WriteString("\n</body>\n</html>") // Lines 10:11
	return nil
}

func (t layout_html) RenderBlock_head_title(c templates.RenderContext) error {
	_, err := c.WriteString("{BLOCK head_title}")
	return err
}

func (t layout_html) RenderBlock_page_title(c templates.RenderContext) error {
	_, err := c.WriteString("{BLOCK page_title}")
	return err
}

func (t layout_html) RenderBlock_menu(c templates.RenderContext) error {
	_, err := c.WriteString("{BLOCK menu}")
	return err
}

func (t layout_html) RenderBlock_content(c templates.RenderContext) error {
	_, err := c.WriteString("{BLOCK content}")
	return err
}
