# Hand-in02

## a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

By illustrating the three-way handshake we decided to use a struct that we call packet. A packet has a seqNumber, ackNumber and syn that is used to ensure the connection that is established is the same.

After ensuring the connection has been established, we use another struct, dataPacket, to represent the actual data string. a dataPacket has a data string and metaData int to represent the order of the messages.

## b) Does your implementation use threads or processes? Why is it not realistic to use threads?

Yes, we use two threads, go client(ch, bufferedDataCh) and go server(ch, bufferedDataCh). This is not realistic since the protocol runs across a network as mentioned in the hand-in description.

## c) How do you handle message re-ordering?

We add each message according to their metaData in a map[int]string. This links each message to a corresponding number. If succeeded, each message is concatenated to the completeMessage, which finally is printed in the order that the client inteded to send it.

## d) How do you handle message loss?

We range through the dataMap and check if there is any empty strings. If there is one, the connection is deemed unstable and the program is terminated. A message is printed in the terminal with what message number went lost.

## e) Why is the 3-way handshake important?

The three-way handshake has two important functions. 1) It makes sure that both sides know that they are ready to transfer data. 2) It also allows both sides to agree on the initial sequence numbers, which are sent and acknowledged during the handshake.
