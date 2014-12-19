package code

import (
	"io"
	"github.com/strongo/templates"
	"github.com/strongo/templates/prototype/models"
//	"github.com/strongo/templates/prototype/code/i18n/ru_ru"
)

type Payload_Index_html struct {
	User *models.User
	BgColor string
}

type Index_html struct {
	i18n templates.I18n
	extends layout_html_struct
	payload Payload_Index_html
}

func NewIndex_html(locale string, payload Payload_Index_html) templates.Template {
	template := Index_html{i18n: I18Storage[locale], payload: payload}
	template.extends = New_layout_html_struct(
		template.i18n,
		template,
		Payload_Layout_html{BgColor: payload.BgColor},
	)
	return template
}

func (t Index_html) Render(writer io.Writer) error {
	layoutPayload := Payload_Layout_html{BgColor: t.payload.BgColor}
	return t.extends.Render(writer, layoutPayload)
}

func (t Index_html) RenderBlock_head_title(writer io.Writer) error {
	_, err := writer.Write([]byte("Index.html"))
	return err
}

func (t Index_html) RenderBlock_page_title(writer io.Writer) error {
	_, err := writer.Write([]byte("Index.html!"))
	return err
}

func (t Index_html) RenderBlock_menu(writer io.Writer) error {
	return t.extends.RenderBlock_menu(writer)
}

func (t Index_html) RenderBlock_content(writer io.Writer) error {
	if _, err := writer.Write([]byte("<p>")); err != nil {
		return err
	}
	if _, err := writer.Write([]byte(t.i18n.GetText("Put your content here."))); err != nil {
		return err
	}
	if _, err := writer.Write([]byte("</p>")); err != nil {
		return err
	}
	return nil
}
