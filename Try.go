package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch:= make(chan packet)
	//ch:= make(chan packet, 2) make the channel af buffered channel, to time the number of elements in the channel????
	now := time.Now()
	fmt.Println(now)

	go client(ch);
	go server(ch);

	for {

	};
} 

func client(ch chan packet){
	//step 1: the client want to make a connection. Sending first packet
	firstPacket := packet{rand.Intn(1000), 0, 1, ""}
	ch <- firstPacket

	//step 2: the client recieves the syn-ack packet
	
	recievedSynAckPacket := <- ch
	toSendAckPacket := recievedSynAckPacket

	//check if the recieved syn-ack packets ackNumber is equal to the first seqNum + 1
	if(recievedSynAckPacket.ackNumber == firstPacket.seqNumber+1 ){
		//step 3: the client sends back ack-packet to the server
		toSendAckPacket.ackNumber = recievedSynAckPacket.seqNumber +1; 
		toSendAckPacket.seqNumber = recievedSynAckPacket.ackNumber; 
		toSendAckPacket.data = "Hello I'm data"
		ch <- toSendAckPacket
	}
}

	func server(ch chan packet){
	//step 1: the server recieves the first packet.
	recievedSynPacket := <- ch
	toSendSynAckPacket := recievedSynPacket;

	//step 2: the sever sends back syn-ack packet
	toSendSynAckPacket.ackNumber = recievedSynPacket.seqNumber +1;
	toSendSynAckPacket.seqNumber = rand.Intn(1000);
	ch <- toSendSynAckPacket
	
	//step 3: the server recieves ack-packet
	recievedAckPacket := <- ch

	//checks if the sequence number is equal to the ackNumber it sent with the Syn-Ack packet
	if(recievedAckPacket.seqNumber == recievedSynPacket.seqNumber + 1){
		recievedAckPacket.syn = 0;
		fmt.Println("Connection established", recievedAckPacket)
		fmt.Println("data recieved was", recievedAckPacket.data)
		
	} else {
		fmt.Println("Something went wrong the recieved seq is not correct")
	}
}


type packet struct{
	seqNumber int //sequence number
	ackNumber int //acknowledgement number
	syn int		//syncronization
	data string //the data to send
}
