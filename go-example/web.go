/*package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string
var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <- messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <- entering:
			clients[cli] = true
		
		case cli := <- leaving:
			clients[cli] = false
			close(cli)//关闭chan
		}
	}
}


func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "you are " + who
	messages <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}


func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	//无限循环
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func test_ex_chat()  {
    listener, err := net.Listen("tcp", "localhost:8000")
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
        go handleConn_ex813(conn)
    }
}
type client1 chan<- string // an outgoing message channel

var (
    entering1 = make(chan client)
    leaving1  = make(chan client)
    messages1 = make(chan string) // all incoming client messages
    //练习 8.12： 使broadcaster能够将arrival事件通知当前所有的客户端。
    // 为了达成这个目的，你需要有一个客户端的集合，
    // 并且在entering和leaving的channel中记录客户端的名字。
    clientMap1 = make(map[string]string)
    oldKey1 = int64(10001)
)
func broadcaster1() {
    clients := make(map[client]bool) // all connected clients
    for {
        select {
        case msg := <-messages1:
            // Broadcast incoming message to all
            // clients outgoing message channels.
        for cli := range clients {
            //练习 8.15： 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。
            // 修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好写
            if !clients[cli] {
                continue
            }
            cli <- msg
        }
        case cli := <-entering1:
            clients[cli] = true

        case cli := <-leaving1:
            delete(clients, cli)
            close(cli)
        }
    }
}
func handleConn_ex(conn net.Conn) {
    ch := make(chan string) // outgoing client messages
    go clientWriter1(conn, ch)

    who := conn.RemoteAddr().String()
    //练习 8.14：
    // 修改聊天服务器的网络协议这样每一个客户端就可以在entering时可以提供它们的名字。
    // 将消息前缀由之前的网络地址改为这个名字。
    if len(clientMap1[who]) <= 0 {
        str := "num" + fmt.Sprint(oldKey1)
        clientMap1[who] =  str
        oldKey1++
    }
    ch <- "You are " + clientMap1[who]
    messages1 <- clientMap1[who] + " has arrived"
    entering1 <- ch

    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages1 <- clientMap1[who] + ": " + input.Text()
    }
    // NOTE: ignoring potential errors from input.Err()

    leaving1 <- ch
    messages1 <- clientMap1[who] + " has left"
    conn.Close()
}
//练习 8.13： 使聊天服务器能够断开空闲的客户端连接，
// 比如最近五分钟之后没有发送任何消息的那些客户端。
// 提示：可以在其它goroutine中调用conn.Close()来解除Read调用，
// 就像input.Scanner()所做的那样
func handleConn_ex813(conn net.Conn) {
    ch := make(chan string) // outgoing client messages
    go clientWriter(conn, ch)

    who := conn.RemoteAddr().String()
    ch <- "You are " + who
    messages1 <- who + " has arrived"
    entering1 <- ch

    input := bufio.NewScanner(conn)

    abort := make(chan string)

    go func() {
        for  {
            select {
            case <-time.After(5 * time.Minute):
                conn.Close()
            case  str := <-abort:
                messages1 <- str
            }
        }
    }()

    for input.Scan() {
        str := input.Text()
        if str == "exit" {
            break
        }
        if len(str) > 0 {
            abort <- who + ": " + input.Text()
        }
    }
    // NOTE: ignoring potential errors from input.Err()

    leaving1 <- ch
    messages1 <- who + " has left"
    conn.Close()
}

func clientWriter1(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
    }
}
