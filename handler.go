package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"sync"
)

var (
	upgrader = websocket.Upgrader{}
	mutex    = sync.Mutex{}
	clients  = make(map[string]*websocket.Conn)
	users    = make(map[string]user)
)

func handleJoin(cli *websocket.Conn, join message) error {
	mutex.Lock()
	defer mutex.Unlock()
	online := *newOnline()
	for _, usr := range users {
		online.Users = append(online.Users, usr)
	}
	if err := cli.WriteJSON(&online); err != nil {
		return err
	}
	fmt.Printf("Join: %v\n", join.From)
	clients[join.From.Email] = cli
	users[join.From.Email] = join.From
	broadcast <- join
	return nil
}

func handleLeave(join message) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("Leave: %v\n", join.From)
	delete(clients, join.From.Email)
	delete(users, join.From.Email)
	broadcast <- message{
		statusLeave,
		join.From,
		"",
	}
}

func handleConn(c echo.Context) error {
	cli, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = cli.Close()
	}()
	var join message
	if err := cli.ReadJSON(&join); err != nil {
		return err
	} else if join.Status != statusJoin {
		return echo.NewHTTPError(403, "unexpected message")
	} else if _, ok := users[join.From.Email]; ok {
		return echo.NewHTTPError(403, "duplicate user")
	}
	if err := handleJoin(cli, join); err != nil {
		return err
	}
	defer handleLeave(join)
	for {
		var msg message
		if err := cli.ReadJSON(&msg); err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		} else if msg.Status == statusLeave && msg.From == join.From {
			break
		}
		broadcast <- msg
	}
	return nil
}
