package prototype

import (
	"time"
	"github.com/strongo/templates"
)

type AuthorCard_Payload struct {
	AuthorId int
}


type AuthorCard_Data struct {
	author *Author
}

type AuthorCard struct {
	component *templates.StrongoComponent
	payload AuthorCard_Payload
	dataProvider DataProvider

	data AuthorCard_Data
}

func NewAuthorCard(payload AuthorCard_Payload) *AuthorCard {
	return &AuthorCard{
		component: templates.NewStrongoComponent(nil),
		payload: payload,
		dataProvider: NewDataProvider(time.Millisecond*10),
	}
}

func (self *AuthorCard) GetData() {
	go func(){
		self.data = AuthorCard_Data{author: self.dataProvider.GetAuthor(self.payload.AuthorId)}
		self.component.OnDataReady()
	}()
}
//1,000,000,000
//0,011,382,440 - delay 10ms
//0,001,263,847 - delay 1ms

func (self *AuthorCard) Render(c templates.RenderContext) {
	c.WriteString("<div>\nAuthor: ")
	self.component.WhenDataReady()
//	log.Print("Render: " + self.Author.Name)
	c.WriteString(self.data.author.Name)
	c.WriteString("\n</div>")
}
