package entities

type Todo struct {
	Id         int64  `json:'id'`
	Content    string `json:'content'`
	IsComplete bool   `json:'isComplete'`
}
