package models

import "time"

type User struct {
	Id             uint64
	Name           string
	Password       string
	Email          string
	Books          []Book
	BooksFavorites []Book
	Role           int
	Provider       int
	ClientId       string
	ImageURL       string
}

type Book struct {
	Id            string
	Title         string
	PublishDate   time.Time
	Description   string
	Rekomended    string
	File          string
	Image         string
	BookUser      User
	PublisherBook Publisher
	TypeBooks     []TypeBook
	BookFavorite  []User
}

type Publisher struct {
	Id    uint64
	Name  string
	Books []Book
}

type TypeBook struct {
	Id       uint64
	Name     string
	BookType []Book
}

// report

type BookReport struct {
	Id           uint64
	IdBook       string
	TitleBook    string
	Author       UserBookReport
	Publisher    Publisher
	User         UserBookReport
	StatusReport int
	DateReport   time.Time
}

type UserReport struct {
	Id           uint64
	User         UserBookReport
	Admin        UserBookReport
	StatusReport int
	DateReport   time.Time
}

type UserBookReport struct {
	Id    uint64
	Name  string
	Email string
}

const (
	ANON          int = 0
	USER          int = 1
	SELLER        int = 2
	ADMINISTRATIF int = 3
	MANAGER       int = 4
)

const (
	Google   int = 0
	Facebook int = 1
	Github   int = 2
)

const (
	DELETE          int = 0
	ADD             int = 1
	DOWNLOAD        int = 2
	IN_STANDARD     int = 3
	IN_OAUTH        int = 4
	VALIDATED       int = 5
	OUT             int = 6
	PROMOTED_SELLER int = 7
	PROMOTED_ADMIN  int = 8
	DEMOTED_USER    int = 9
)
