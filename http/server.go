package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	host := "0.0.0.0"
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9999"
	}
	err := start(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}
}

func start(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("can't listen %s: %w", addr, err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		log.Print("accept connection")
		if err != nil {
			log.Printf("can't accept: %v", err)
			continue
		}
		log.Print("handle connection")
		handleConn(conn)
	}
}


func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	log.Print(requestLine)
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) != 3 {
		return
	}

	method, request, protocol := parts[0], parts[1], parts[2]
	if method == "GET" && request == "/" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		writer.WriteString("HTTP/1.1 200 Ok\r\n")
		writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(html)))
		writer.WriteString("Content-Type: text/html\r\n")
		writer.WriteString("Connection: close\r\n")
		writer.WriteString("\r\n")
		writer.WriteString(html)
		writer.Flush()
	}

	if method == "GET" && request == "/http/index.html" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/index.html")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: text/html\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/1.png" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/1.png")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: image/png\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/2.jpg" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/2.jpg")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: image/png\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/123.txt" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/123.txt")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: text/html\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/1.pdf" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/1.pdf")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: application/pdf\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}
}

