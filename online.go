package main

type online struct {
	Users []user `json:"users"`
}

func newOnline() *online {
	return &online{
		make([]user, 0),
	}
}
