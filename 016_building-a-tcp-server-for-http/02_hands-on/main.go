package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// create a server that returns the URL of the GET request

func main() {
	li, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(con)
	}
}

//handle extracts the the URL of the GET request and returns it to the user-agent
func handle(con net.Conn){
	defer con.Close()
	con.SetDeadline(time.Now().Add(30 * time.Second))

	// extract the url
	s := bufio.NewScanner(con)
	i := 0
	var url string
	for s.Scan(){
		line := s.Text()
		
		if i == 0 {
			//headers
			fs := strings.Fields(line)
			url = fs[1]
			fmt.Fprint(con, url)
		}
		if line == "" {
			break
		}
		i++
	}

	// return to user-agent
	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Golang Server</title></head><body><strong>%v</strong></body></html>`, url)
	fmt.Fprint(con,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html,santiago/morales\r\n")
	fmt.Fprint(con, "\r\n")
	fmt.Fprint(con,body)

}