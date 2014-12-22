package code

import (
	"github.com/strongo/templates"
	"github.com/strongo/templates/prototype"
)

type Payload_Index_html struct {
	AuthorIds []int
	BgColor string
}

type Index_html struct {
	i18n templates.I18n
	extends layout_html_struct
	payload Payload_Index_html
	component *templates.StrongoComponent
	authorCards []*prototype.AuthorCard
}

func NewIndex_html(locale string, payload Payload_Index_html) templates.Template {

	authorCards := make([]*prototype.AuthorCard, len(payload.AuthorIds))
	components := make([]templates.IStrongoComponent, len(authorCards))
	for i, authorId := range payload.AuthorIds {
		authorCardPayload := prototype.AuthorCard_Payload{AuthorId: authorId}
		authorCard := prototype.NewAuthorCard(authorCardPayload)
		authorCards[i] = authorCard
		components[i] = authorCard
	}
	template := Index_html{
		i18n: I18Storage[locale],
		payload: payload,
		authorCards: authorCards,
	}

	template.component = templates.NewStrongoComponent(components)

	template.extends = New_layout_html_struct(
		template.i18n,
		template,
		Payload_Layout_html{BgColor: payload.BgColor},
	)
	return template
}

func (t Index_html) GetData() {
	t.component.GetData()
}



func (t Index_html) Render(c templates.RenderContext) error {
	layoutPayload := Payload_Layout_html{BgColor: t.payload.BgColor}
	return t.extends.Render(c, layoutPayload)
}

func (t Index_html) RenderBlock_head_title(c templates.RenderContext) error {
	_, err := c.WriteString("Index.html")
	return err
}

func (t Index_html) RenderBlock_page_title(c templates.RenderContext) error {
	_, err := c.WriteString("Index.html!")
	return err
}

func (t Index_html) RenderBlock_menu(c templates.RenderContext) error {
	return t.extends.RenderBlock_menu(c)
}

func (t Index_html) RenderBlock_content(c templates.RenderContext) error {
	if _, err := c.WriteString("<p>"); err != nil {
		return err
	}

	if _, err := c.WriteString(t.i18n.GetText("Put your content here.")); err != nil {
		return err
	}

	for _, authorCard := range t.authorCards{
		c.WriteString("\n<li>")
		authorCard.Render(c)
		c.WriteString("</li>")
	}

	if _, err := c.WriteString("</p>"); err != nil {
		return err
	}
	return nil
}
