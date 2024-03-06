package response

type BooksResponse struct {
	Id       int    `json:"id"`
	Judul    string `json:"judul"`
	Penerbit string `json:"penerbit"`
	Rating   int    `json:"rating"`
}
