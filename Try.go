package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	ch := make(chan packet)
	bufferedDataCh := make(chan dataPacket, 3)

	go client(ch, bufferedDataCh)
	go server(ch, bufferedDataCh)

	for {

	}
}

func client(ch chan packet, dataCh chan dataPacket) {
	//step 1: the client want to make a connection. Sending first packet
	firstPacket := packet{rand.Intn(1000), 0, 1}
	ch <- firstPacket

	//step 2: the client recieves the syn-ack packet
	recievedSynAckPacket := <-ch
	toSendAckPacket := recievedSynAckPacket

	//check if the recieved syn-ack packets ackNumber is equal to the first seqNum + 1
	if recievedSynAckPacket.ackNumber == firstPacket.seqNumber+1 {
		//step 3: the client sends back ack-packet to the server
		toSendAckPacket.ackNumber = recievedSynAckPacket.seqNumber + 1
		toSendAckPacket.seqNumber = recievedSynAckPacket.ackNumber
		ch <- toSendAckPacket

		firstDataPacket := dataPacket{"I am a data string", 0}
		secondDataPacket := dataPacket{"I am a data string1", 1}
		//thirdDataPacket := dataPacket{"I am a data string2", 3}
		thirdDataPacket := dataPacket{"I am a data string2", 2}

		dataCh <- firstDataPacket
		dataCh <- thirdDataPacket
		dataCh <- secondDataPacket

	} else {
		fmt.Println("Something went wrong the recieved ack number is not correct")
	}
}

func server(ch chan packet, dataCh chan dataPacket) {
	//dataArray := []dataPacket{}
	dataMap := make(map[int]string)

	//step 1: the server recieves the first packet.
	recievedSynPacket := <-ch
	toSendSynAckPacket := recievedSynPacket

	//step 2: the sever sends back syn-ack packet
	toSendSynAckPacket.ackNumber = recievedSynPacket.seqNumber + 1
	toSendSynAckPacket.seqNumber = rand.Intn(1000)
	ch <- toSendSynAckPacket

	//step 3: the server recieves ack-packet
	recievedAckPacket := <-ch

	//checks if the sequence number is equal to the ackNumber it sent with the Syn-Ack packet
	if recievedAckPacket.seqNumber == recievedSynPacket.seqNumber+1 {
		recievedAckPacket.syn = 0
		fmt.Println("Connection established")

		for i := 0; i < 3; i++ {
			recievedData := <-dataCh
			dataMap[recievedData.metaData] = recievedData.data
		}

		var maxNumber int = findMax(dataMap)
		completeMessage := ""

		for i := 0; i <= maxNumber; i++ {
			var j = dataMap[i]
			if j == "" {
				fmt.Println("Missing message with number", i)
				os.Exit(i)
			}
			completeMessage += j + " "
		}
		fmt.Println(completeMessage)

		//we need a check to see if there are any messages lost
	} else {
		fmt.Println("Something went wrong the recieved seq is not correct")
	}
}

func findMax(m map[int]string) int {
	var max int
	for key := range m {
		if key > max {
			max = key
		}
	}
	return max
}

type packet struct {
	seqNumber int //sequence number
	ackNumber int //acknowledgement number
	syn       int //syncronization
}

// the data we send after making the connection
type dataPacket struct {
	data     string //the data to send
	metaData int    //if more than one packet is sent, or they are in the wrong order
}
