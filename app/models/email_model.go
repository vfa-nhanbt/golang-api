package models

type EmailModel struct {
	ReceiverEmail string `json:"receiver_email" validate:"email,lte=255,required"`
}

type CreateBookEmailModel struct {
	AuthorName string `json:"author_name" validate:"required"`
	BookTitle  string `json:"book_title" validate:"required"`
	BookUrl    string `json:"book_url"`
	EmailModel
}
