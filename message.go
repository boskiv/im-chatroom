package main

type online struct {
	Users []user `json:"users"`
}

const (
	statusJoin  = 1
	statusLeave = -1
)

type message struct {
	Status  int    `json:"status"`
	From    user   `json:"from"`
	Content string `json:"content"`
}
