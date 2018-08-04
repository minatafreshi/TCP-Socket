package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"time"
	"flag"
	"net"
	"regexp"
)

var host = flag.string ("host" , "" , "the hostname or IP to connect to ; defualts to \"local host"\.)
var port = flag.Int ( "port" , 65534 , "the port to connect to defaults to 65534")

func main()  {
	flag.parse()
	dest := *host + ":" +strconv.Itoa(*port)
	fmt.Printf("connecting to %s ... \n" . dest)
	conn.err := net.dial("TCP" , dest)
	if err != nil {
		if _,t := err.(*net.OpError); t{
			fmt.Println("some problem connecting!")
		} else {
			fmt.Println("unknown error" + err.Error())
		}
		os.Exit(1)
	}
	go readconnection(conn)
	for {
		reader := bufio.NewReader(os.stdin)
		fmt.Print("> ")
		test, _ := reader.Readstring("\n")

		conn.setWriteDeanline(time.now().Add(1*time.second))
		_ , err := conn.write ([]byte(text))
		if err != nil{
			fmt.Println(" Error Writing to stream.")
			break
		}	
	}
}
func readconnection(conn net.conn){
	for {
	scanner := bufio.newscanner(conn)
		for{
			ok := scanner.scan()
			text := scanner.text()
			command := handlecommands(text)

			if !command {
				fmt.Printf("\b\b ** %s \n " , text)
			}
			if !ok {
				fmt.Println("Reached EOF on server connection.")
				break
			}
		}
	}
}
func handlecommands(text string) bool{
	r , err :=regexp.compile("^% . *%$")
	if err == nil {
		if r.matchstring(text){
			switch {
			case text == "%quit%":
				fmt.Println("\b\b server is leaving . hanging up!")
				os.Exit(0)
				
			}
			return true
		}
	}
	return false
}
