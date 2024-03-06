package request

type CreateBooksRequest struct {
	Judul    string `validate:"required,min=1,max=200" json:"judul"`
	Penerbit string `validate:"required,min=1,max=200" json:"penerbit"`
	Rating   int    `validate:"required,min=1,max=5" json:"rating"`
}
