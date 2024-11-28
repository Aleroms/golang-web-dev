package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

/*
# Building upon the code from the previous problem:

Add code to respond to the following METHODS & ROUTES:
	GET /
	GET /apply
	POST /apply

*/
func main(){
	li, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}
func handle(conn net.Conn){
	defer conn.Close()

	s := bufio.NewScanner(conn)
	i := 0
	for s.Scan() {
		ln := s.Text()

		if i == 0 {
			//Headers
			fields := strings.Fields(ln)
			err := mux(conn,fields)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		if ln == "" {
			//check if has body
			break
		}
		i++

		fmt.Println(ln)
	}

	fmt.Println("Code got here.")
}
func mux(conn net.Conn, fields []string) (err error){

	if len(fields) != 3 {
		return errors.New("Not Headers. Length incorrect")
	}
	method := fields[0]
	resource := fields[1]


	if method == "GET" {

		if resource == "/"{
			home(conn)
		} else if resource == "/apply" {
			applyGet(conn)
		}
	} else if method == "POST" {
		if resource == "/apply" {
			applyPost(conn)
		}
	}

	return err
}
func applyGet(conn net.Conn){
	body := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Apply Test</title>
  </head>
  <body>
    <h1>GET /Apply Test</h1>
    <a href="/">Home</a>
    <form action="/apply" method="POST">
      <input type="tel" name="telphone" id="telphone" />
	  <input type="submit" value="submit">
    </form>
  </body>
</html>
`
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	io.WriteString(conn,"Content-Type:text/html\r\n")
	fmt.Fprintf(conn,"Content-Length:%d\r\n",len(body))
	io.WriteString(conn,"\r\n")
	io.WriteString(conn,body)

	//logging purposes only
	io.WriteString(os.Stdout,"==================")
	io.WriteString(os.Stdout,"HTTP/1.1 200 OK\r\n")
	io.WriteString(os.Stdout,"Content-Type:text/html\r\n")
	fmt.Fprintf(os.Stdout,"Content-Length:%d\r\n",len(body))
	io.WriteString(os.Stdout,"\r\n")
	io.WriteString(os.Stdout,body)
}
func applyPost(conn net.Conn){
	body := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Apply Test</title>
  </head>
  <body>
    <h1>GET /Apply Test</h1>
    <a href="/">Home</a>
    POST
  </body>
</html>
`
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	io.WriteString(conn,"Content-Type:text/html\r\n")
	fmt.Fprintf(conn,"Content-Length:%d\r\n",len(body))
	io.WriteString(conn,"\r\n")
	io.WriteString(conn,body)

	//logging purposes only
	io.WriteString(os.Stdout,"==================")
	io.WriteString(os.Stdout,"HTTP/1.1 200 OK\r\n")
	io.WriteString(os.Stdout,"Content-Type:text/html\r\n")
	fmt.Fprintf(os.Stdout,"Content-Length:%d\r\n",len(body))
	io.WriteString(os.Stdout,"\r\n")
	io.WriteString(os.Stdout,body)
}
func home(conn net.Conn){
	
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Apply Test</title>
	</head>
	<body>
    <h1>Apply Test</h1>
    <a href="/apply">CLICK ME</a>
	</body>
	</html>`
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	io.WriteString(conn,"Content-Type:text/html\r\n")
	fmt.Fprintf(conn,"Content-Length:%d\r\n",len(body))
	io.WriteString(conn,"\r\n")
	io.WriteString(conn,body)
	
	//logging purposes only
	io.WriteString(os.Stdout,"==================")
	io.WriteString(os.Stdout,"HTTP/1.1 200 OK\r\n")
	io.WriteString(os.Stdout,"Content-Type:text/html\r\n")
	fmt.Fprintf(os.Stdout,"Content-Length:%d\r\n",len(body))
	io.WriteString(os.Stdout,"\r\n")
	io.WriteString(os.Stdout,body)
}