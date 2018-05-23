package prototype

import (
	"time"
	"strconv"

)
type Author struct {
	Id int
	Name string
}

type Book struct {
	Id int
	Name string
}

type DataProvider struct {
	Latency time.Duration
}

func NewDataProvider(latency time.Duration) DataProvider {
	return DataProvider{Latency: latency}
}

func (dp DataProvider) GetAuthors() map[int]*Author {
	result := make(map[int]*Author)
	for i := 0; i < 20; i++ {
		result[i] = dp.GetAuthor(i)
	}
	return result
}

func (dp DataProvider) GetAuthor(authorId int) *Author {
	if dp.Latency > 0 {
		time.Sleep(dp.Latency)
	}
	return &Author{
		Id: authorId,
		Name: "John Smith #" + strconv.Itoa(authorId),
	}
}
