package entity

type Person struct {
	Name  string `json:"name" binding:"required"`
	Age   uint8  `json:"age" binding:"gte=0,lte=130"`
	Email string `json:"email" validate:"required,email"`
}

type Video struct {
	GUID        string `json:"id"`
	Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Actors      int8   `json:"actors" binding:"gt=2,lt=20"`
	Author      Person `json:"author"`
}
