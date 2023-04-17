package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var port = 4711
var channel = make(chan string, 2)

func WriteFortune() {
	fortune := []string{"You will have good fortune!", "This is not a fortune cookie you just ate", "Maybe you should just do it", "Mold :D"}

	for {
		log.Println("That's a good fortune! Lets publish it!")
		channel <- fortune[rand.Int()%len(fortune)]
		log.Println("That was stressful! Lets sleep for a while...")
		time.Sleep(5 * time.Second)
	}
}

func Fortune(w http.ResponseWriter, r *http.Request) {
	fortune := <- channel
	w.Write([]byte(fortune))
}

func main() {
	address := fmt.Sprint("localhost:", port)

	go WriteFortune()

	http.HandleFunc("/", Fortune)
	log.Printf("Starting webserver on http://%v\n", address)
	err := http.ListenAndServe(address, nil)
	log.Panic(err)
}
