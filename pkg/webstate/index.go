package webstate

type Data struct {
	Cards Cards
    Title string
    Artist string
}

func NewData() Data {
	return Data{
        Title: "",
        Artist: "",
    }
}
