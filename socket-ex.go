/*
* 一个socket的例子
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func connectionHandler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 128)
	var name, passwd string

	greeting := "Welcome to HiWorld.\n"
	conn.Write([]byte(greeting))

	for {

		conn.Write([]byte("Please enter your name : "))
		l, err := conn.Read(buf)
		//过滤掉回车换行符
		if err != nil || l <= 2 {
			continue
		}
		name = string(buf[0 : l-2])
		fmt.Println("Recieve Length: ", l)
		fmt.Println("buf: ", buf)
		fmt.Println("Name: ", name)
		for i := 0; i < l; i++ {
			buf[i] = 0x00
		}
		break
	}

	for {

		conn.Write([]byte("Please enter your password : "))
		l, err := conn.Read(buf)
		if err != nil || l <= 0 {
			continue
		}
		fmt.Println("Recieve Length: ", l)
		fmt.Println("buf: ", buf)
		passwd = string(buf[0 : l-2])
		fmt.Println("passwd: ", passwd)
		for i := 0; i < l; i++ {
			buf[i] = 0x00
		}
		break
	}

	if name == "hello" || passwd == "123456" {
		conn.Write([]byte("Successful."))
		return
	}

	conn.Write([]byte("Error."))
	return

}

func main() {
	fmt.Println("Starting Server...")
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	for {

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go connectionHandler(conn)
	}
}
