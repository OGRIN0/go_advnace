package main

// // import (
// // 	"fmt"
// // 	"log"
// // 	"time"
// // 	"sync"
// // )

// // type Message struct {
// // 	chats []string 
// // 	friends []string 
// // }

// // func main(){
// // 	now := time.Now()
// // 	id := getUserByName("jhon")
// // 	println(id)

// // 	wg := &sync.WaitGroup{}
// // 	ch := make(chan *Message, 2)

// // 	wg.Add(2)

// // 	go getUserChats(id, ch, wg)
// // 	go getUserFriends(id, ch, wg)

// // 	wg.Wait()
// // 	close(ch)

// // 	for msg := range ch {
// // 		log.Println(msg)
// // 	}

// // 	log.Println(time.Since(now))
// // }

// // func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
// // 	time.Sleep(time.Second*1)

// // 	ch <- &Message{
// // 		friends: []string{
// // 		"jhon",
// // 		"doe",
// // 		"dove",
// // 		"paarot",
// // 		"lab",
// // 		},
// // 	}

// // 	wg.Done()
// // }

// // func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
// // 	time.Sleep(time.Second*2)
// // 	ch <- &Message{
// // 		chats: []string{
// // 			"jhon",
// // 			"joe",
// // 			"dove",
// // 		},
// // 	}

// // 	wg.Done()
// // }

// // func getUserByName(name string) string{
// // 	time.Sleep(time.Second*1)
// // 	return fmt.Sprintf("%s-2", name)
// // }

// func main(){
// 	userId := 43

// 	println(userId)

// 	println(&userId)

// 	a := &userId

// 	println(a)

// 	var age *int 

// 	age = &userId 

// 	update(age, 42)

// 	println(age)

// 	updateButNotReally(*age, 67)
// 	println(*age)
// }

// func update(val *int, to int){
// 	*val = to 
// }

// func updateButNotReally(val int, to int){
// 	val = to
// }