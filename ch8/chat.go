package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	Msg  chan<- string // an outgoing message channel
	Name string        // the name of the client
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
	timeout  = time.Minute * 5
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients

	sendToAllCh := func(msg string, name_flag bool) {
		for cli := range clients {
			m := msg
			if name_flag {
				m += cli.Name
			}
			select {
			case cli.Msg <- m:
			default:
				// Do nothing.
			}
		}
	}

	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			sendToAllCh("["+time.Now().Format(time.StampMilli)+"] "+msg, false)
		case cli := <-entering:
			clients[cli] = true
			sendToAllCh("Present:", true)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Msg)
		}
	}
}

func handleConn(conn net.Conn) {
	out := make(chan string, 1) // outgoing client messages
	go clientWriter(conn, out)
	in := make(chan string) // incoming client
	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			in <- input.Text()
		}
	}()

	who := strings.Split(conn.RemoteAddr().String(), ":")[1]
	out <- "Print your name: "
	select {
	case name := <-in:
		who = name
	case <-time.After(timeout):
		conn.Close()
		return
	}

	cli := client{out, who}
	out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

loop:
	for {
		select {
		case msgs := <-in:
			messages <- who + ": " + msgs
		case <-time.After(timeout):
			close(in)
			break loop
		}
	}

	// NOTE: ignoring potential errors from input.Err()
	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
