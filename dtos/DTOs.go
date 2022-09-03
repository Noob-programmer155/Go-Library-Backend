package dtos

type IdAndName struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type BookDTO struct {
	Id           uint64   `json:"id"`
	Title        string   `json:"title"`
	NewPublisher string   `json:"newPublisher"`
	Publisher    string   `json:"publisher"`
	Description  string   `json:"description"`
	Theme        []int    `json:"theme"`
	NewTheme     []string `json:"newTheme"`
	Favorite     bool     `json:"favorite"`
}

type BookDTOResp struct {
	Id          uint64      `json:"id"`
	Title       string      `json:"title"`
	Author      string      `json:"author"`
	Publisher   IdAndName   `json:"publisher"`
	PublishDate string      `json:"publishDate"`
	Description string      `json:"description"`
	Theme       []IdAndName `json:"theme"`
	File        string      `json:"file"`
	Image       string      `json:"image"`
	Favorite    bool        `json:"favorite"`
	Status      bool        `json:"status"`
}

type UserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserInfoDTO struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	ImageURL string `json:"image_url"`
	Status   bool   `json:"status"`
}

type BookResponse struct {
	Data        []BookDTOResp `json:"data"`
	SizeAllPage int           `json:"sizeAllPage"`
}

type TypeResponse struct {
	Data        []IdAndName `json:"data"`
	SizeAllPage int         `json:"sizeAllPage"`
}

type UserResponse struct {
	Data        []UserInfoDTO `json:"data"`
	SizeAllPage int           `json:"sizeAllPage"`
}
