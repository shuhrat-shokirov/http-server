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
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Can't close conn: %v", err)
		}
	}()
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
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Can't close conn: %v", err)
		}
	}()
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
		contentType := ""
		fileName := ""
		if method == "GET" && protocol == "HTTP/1.1" {
			if strings.Contains(request, "?download") {
				contentType = "application/octet-stream"
				request = strings.Replace(request, "?download", "", -1)
			}
			switch request {
			case "/":
				fileName = "cmd/http/files/index.html"
				if(contentType == ""){
					contentType = "text/html"}
			case "/index-html.html":
				fileName = "cmd/http/files/index-html.html"
				if(contentType == ""){
					contentType = "text/html"}
			case "/1.png":
				fileName = "cmd/http/files/1.png"
				if(contentType == ""){
					contentType = "image/png"}
			case "/2.jpg":
				fileName = "cmd/http/files/2.jpg"
				if(contentType == ""){
					contentType = "image/jpg"}
			case "/123.txt":
				fileName = "cmd/http/files/123.txt"
				if(contentType == ""){
					contentType = "text/html"}
			case "/1.pdf":
				fileName = "cmd/http/files/1.pdf"
				if(contentType == ""){
					contentType = "application/pdf"}
			case "/html5.png":
				fileName = "cmd/http/files/html5.png"
				if(contentType == ""){
					contentType = "application/pdf"}
			default:
				fileName = "cmd/http/files/404-NotFound.png"
				contentType = "image/png"
			}
				writeHeader(conn, fileName, contentType, request)
			return
	}
}

func writeHeader(conn net.Conn, fileName, contentType, request string) {
	writer := bufio.NewWriter(conn)
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("can't read file: %v", err)
	}
	_, err = writer.WriteString("HTTP/1.1 200 OK\r\n")
	if err != nil {
		log.Printf("can't write: %v", err)
	}
	_, err = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
	if err != nil {
		log.Printf("can't write: %v", err)
	}
	_, err = writer.WriteString("Content-Type:" + " " + contentType + "\r\n")
	if err != nil {
		log.Printf("can't write: %v", err)
	}
	_, err = writer.WriteString("Connection: Close\r\n")
	if err != nil {
		log.Printf("can't write: %v", err)
	}
	_, err = writer.WriteString("\r\n")
	if err != nil {
		log.Printf("can't write: %v", err)
	}
	_, err = writer.Write(bytes)
	if err != nil {
		log.Printf("can't write: %v", err)
	}
	err = writer.Flush()
	if err != nil {
		log.Printf("can't sent response: %v", err)
	}
	log.Printf("response on: %s", request)
}
