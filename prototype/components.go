package prototype

import (
//	"log"
	"time"
	"strconv"
	"github.com/strongo/templates"
)

type AuthorCard_Payload struct {
	AuthorId int
}


type AuthorCard struct {
	component *templates.StrongoComponent
	payload AuthorCard_Payload

	author *Author
	books []*Book
}

func NewAuthorCard(payload AuthorCard_Payload) *AuthorCard {
	return &AuthorCard{
		component: templates.NewStrongoComponent(nil),
		payload: payload,
	}
}

func (self *AuthorCard) GetData() {
//	log.Print("GetData1: " + strconv.Itoa(self.payload.AuthorId))
	go func(){
//		log.Print("GetData2: " + strconv.Itoa(self.payload.AuthorId))
		time.Sleep(time.Millisecond*1)
		self.author = &Author{
			Id: self.payload.AuthorId,
			Name: "John Smith #" + strconv.Itoa(self.payload.AuthorId),
		}
//		log.Print("GetData: " + self.Author.Name)
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
	c.WriteString(self.author.Name)
	c.WriteString("\n</div>")
}
