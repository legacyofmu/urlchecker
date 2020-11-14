package main

import (
	"fmt"
	"time"
)

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func testSexy() {
	go sexyCount("dhpark")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + "is sexy"
}

func checkPeopleSexy() {
	c := make(chan string)
	people := [2]string{"dhpark", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func main() {
	//testSexy()
	checkPeopleSexy()
}
