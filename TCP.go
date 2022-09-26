package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("hej")

// 	ch0 := make(chan data)
// 	var counter int
// 	s := make([]string, 0)

// 	go sender(ch0, counter, "lol")
// 	go reciever(ch0, s)

// }

// func sender(ch chan data, c int, m string) {
// 	step_one(ch, c)

// 	step_three(ch, m, c)

// }

// func reciever(ch chan data, s []string, c int) {
// 	step_two(ch, c)
// 	var save = <-ch
// 	s[save.sequenceNr] = save.message
// }

// func step_one(ch chan data, c int) {
// 	first := data{sequenceNr: c, message: "Can I make a connection?"}
// 	ch <- first
// }

// func step_two(ch chan data, c int) {
// 	var save = <-ch
// 	second := data{sequenceNr: save.sequenceNr + 1, message: "Yes you can!"}
// 	ch <- second
// }

// func step_three(ch chan data, m string, c int) {
// 	var save = <-ch

// 	third := data{sequenceNr: c, message: m}
// 	ch <- third
// }

// type data struct {
// 	sequenceNr int
// 	message    string
// }
