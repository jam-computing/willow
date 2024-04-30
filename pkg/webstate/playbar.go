package webstate

type PlaybarData struct {
	Title  string
	Artist string
}

func NewPlaybarData(title, artist string) *PlaybarData {
    return &PlaybarData{
        Title: title,
        Artist: artist,
    }
}
