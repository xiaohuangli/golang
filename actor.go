package main

import (
	"sync"
	"time"
	"fmt"
)

type Message struct {
	key		string
	value	string
}

type Actor struct {
	msgQ		chan Message
	waitGroup	sync.WaitGroup
}

func NewActor() *Actor {
	return &Actor{
		msgQ:make(chan Message, 1024),
	}
}

func (this *Actor) Run() {
	for {
		select {
		case <- this.msgQ:
			this.Process()
		}
	}
}

func (this *Actor) Process() {

}

func (this *Actor) AddMessage(msg Message) {
	this.msgQ <- msg
}

type Client struct {
	*Actor
	frame				uint64
	lastupdateTime		int64
}

func NewClient() *Client {
	return &Client{
		Actor:NewActor(),
		frame:0,
		lastupdateTime:time.Now().Unix(),
	}
}

type Server struct {
	*Actor
	frame 			uint64
	updateTime		int64
}

func NewServer() *Server {
	return &Server{
		Actor:NewActor(),
		frame:0,
		updateTime:time.Now().Unix(),
	}
}

func main()  {
	timer := time.NewTimer(time.Second * 2)
	ticker := time.NewTicker(time.Second * 1)

	for {
		time.Now().Unix()
		select {
		case now1 := <- timer.C:
			fmt.Println("timer: ", now1.Unix())
		case now2 := <- ticker.C:
			fmt.Println("ticker: ", now2.Unix())
		}
	}
}

type TimeOuter interface {
	TimeOut(int64)
}