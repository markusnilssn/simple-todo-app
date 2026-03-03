package internal

type Priority int64
const (
	Low = 0 
	Medium = 1
	High = 2
	Critical = 3
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Priority Priority `json:"priority"`
}