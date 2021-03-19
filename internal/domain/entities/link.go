package entities

type Link struct {
	From string `json:"from" field:"linkFrom"`
	To   string `json:"to" field:"linkTo"`
}
