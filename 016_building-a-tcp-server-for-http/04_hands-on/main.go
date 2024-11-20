package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// create a multiplexer (mux, servemux, router, server, http mux, ...) so that your server responds to different URIs and METHODS.


func main(){
	li, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln(err)
	}	
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err  != nil {
			fmt.Fprintln(conn, err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn){
	s := bufio.NewScanner(conn)
	i := 0

	for s.Scan(){
		ln := s.Text()
		fmt.Println(ln)

		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}



func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]

	fmt.Printf("METHOD: %s\nURL: %s\n",m, u)

	var body string
	if m == "GET" && u == "/" {
		body = home()
		
	} else if m == "GET" && u == "/about" {
		body = about()
	} else if m == "GET" && u == "/contact-me" {
		body = contactMe()
	}
	
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n",len(body))
	fmt.Fprint(conn,"Content-Type: text/html; charset=utf-8\r\n")
	fmt.Fprint(conn,"\r\n")
	fmt.Fprint(conn,body)
}
func home() string {
	return `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Home</title>
  </head>
  <body>
    <h1>Home</h1>
    <ul>
      <li><a href="/about">About</a></li>
      <li><a href="/contact-me">Contact Me</a></li>
      <li> --> <a href="/">Home</a></li>
    </ul>
  </body>
</html>`
}
func about() string {
	return `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>About</title>
  </head>
  <body>
    <h1>About</h1>
    <ul>
      <li> --> <a href="/about">About</a></li>
      <li><a href="/contact-me">Contact Me</a></li>
      <li><a href="/">Home</a></li>
    </ul>
  </body>
</html>`
}

func contactMe() string {
	return `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Contact Me</title>
  </head>
  <body>
    <h1>Contact Me</h1>
    <ul>
      <li><a href="/about">About</a></li>
      <li> --> <a href="/contact-me">Contact Me</a></li>
      <li><a href="/">Home</a></li>
    </ul>
  </body>
</html>`
}