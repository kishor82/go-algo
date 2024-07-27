package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserByName("kishor")
	println(id)

	wg := &sync.WaitGroup{}

	ch := make(chan *Message, 2) // buffer size is important

	wg.Add(2)

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}

	log.Println(time.Since(now))
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)

	ch <- &Message{
		chats: []string{
			"kishor",
			"janr",
			"joe",
		},
	}
	wg.Done()
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)

	ch <- &Message{
		friends: []string{
			"kishor",
			"janr",
			"joe",
			"donald",
		},
	}
	wg.Done()
}

// Avoid this , thi will not throw an error but will cause memory leak, use defer to close when function return
func leaky() {
	ch := make(chan int)

	go func() {
		msg := <-ch
		log.Println(msg)
	}()
}
