package main

import (
	"fmt"
)

func main() {
	fmt.Println("hej");
	ch:= make(chan packet)

	go client(ch);
	go server(ch);

	for {

	};
} 

func client(ch chan packet){
	//step 1: the client want to make a connection. Sending first packet
	firstPacket := packet{9001, 0, 1}
	ch <- firstPacket

	//step 2: the client recieves the syn-ack packet
	recievedSynAckPacket := <- ch
	toSendAckPacket := recievedSynAckPacket

	//check if the recieved syn-ack packets ackNumber is equal to the first seqNum + 1
	if(recievedSynAckPacket.ackNumber == firstPacket.seqNumber+1 ){
		//step 3: the client sends back ack-packet to the server
		toSendAckPacket.ackNumber = recievedSynAckPacket.seqNumber +1; 
		toSendAckPacket.seqNumber = recievedSynAckPacket.ackNumber; 
		ch <- toSendAckPacket
	}
}


	func server(ch chan packet){

	//step 1: the server recieves the first packet.
	recievedSynPacket := <- ch
	toSendSynAckPacket := recievedSynPacket;

	//step 2: the sever sends back syn-ack packet
	toSendSynAckPacket.ackNumber = recievedSynPacket.seqNumber +1;
	toSendSynAckPacket.seqNumber = 5001;
	ch <- toSendSynAckPacket
	
		//step 3: the server recieves ack-packet
	recievedAckPacket := <- ch

	//checks if the sequence number is equal to the ackNumber it sent with the Syn-Ack packet
	if(recievedAckPacket.seqNumber == recievedSynPacket.seqNumber + 1){
		fmt.Println("YESSSSSS")
	} else {
		fmt.Println("Something went wrong the recieved seq is not correct")
	}
}


type packet struct{
	seqNumber int //sequence number
	ackNumber int //acknowledgement number
	syn int		//syncronization
}
