package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var portMapping = map[string]string{
	"8010": "US/Eastern",
	"8020": "Asia/Tokyo",
	"8030": "Europe/London",
}

var myPortMapping = map[string]int{
	"8010": 55555,
	"8020": 55556,
	"8030": 55557,
}

var indexMapping = map[int]string{
	0: "8010",
	1: "8020",
	2: "8030",
}

var count int
var lfindex int

func main() {
	myIP := "localhost"
	myAddr := new(net.TCPAddr)
	myAddr.IP = net.ParseIP(myIP)

	//	var tab string = ""
	for port, _ := range portMapping {
		syncConnect(myPortMapping[port], port)

		//		for i := 0; i < count; i++ {
		//			tab += " "
		//		}
		//		go connect(myPortMapping[port], port, count, tab)
		//		count += 3
	}
	time.Sleep(time.Minute)
}

func syncConnect(myPort int, serverPort string) {
	myIP := "localhost"
	serverIP := "localhost"

	usMyAddr := new(net.TCPAddr)
	usMyAddr.IP = net.ParseIP(myIP)
	usMyAddr.Port = 55555

	jpMyAddr := new(net.TCPAddr)
	jpMyAddr.IP = net.ParseIP(myIP)
	jpMyAddr.Port = 55556

	euMyAddr := new(net.TCPAddr)
	euMyAddr.IP = net.ParseIP(myIP)
	euMyAddr.Port = 55557

	usAddr := serverIP + ":" + "8010"
	usTcpAddr, _ := net.ResolveTCPAddr("tcp", usAddr)

	jpAddr := serverIP + ":" + "8020"
	jpTcpAddr, _ := net.ResolveTCPAddr("tcp", jpAddr)

	euAddr := serverIP + ":" + "8030"
	euTcpAddr, _ := net.ResolveTCPAddr("tcp", euAddr)

	conn1, err1 := net.DialTCP("tcp", usMyAddr, usTcpAddr)
	conn2, err2 := net.DialTCP("tcp", jpMyAddr, jpTcpAddr)
	conn3, err3 := net.DialTCP("tcp", euMyAddr, euTcpAddr)
	checkError(err1)
	checkError(err2)
	checkError(err3)
	fmt.Printf("%s\t%s\t%s\n", portMapping["8010"], portMapping["8020"], portMapping["8030"])
	for {

		tmp1 := make([]byte, 256)
		n1, err1 := conn1.Read(tmp1)
		if err1 != nil {
			if err1 != io.EOF {
				fmt.Println("read error: ", err1)
			}
			break
		}
		fmt.Printf("%s\t", tmp1[:n1])
		defer conn1.Close()

		tmp2 := make([]byte, 256)
		n2, err2 := conn2.Read(tmp2)
		if err2 != nil {
			if err2 != io.EOF {
				fmt.Println("read error: ", err2)
			}
			break
		}
		fmt.Printf("%s\t", tmp2[:n2])
		defer conn2.Close()

		tmp3 := make([]byte, 256)
		n3, err3 := conn3.Read(tmp3)
		if err3 != nil {
			if err3 != io.EOF {
				fmt.Println("read error: ", err3)
			}
			break
		}
		fmt.Printf("%s\n", tmp3[:n3])
		defer conn3.Close()
	}
}

// For goroutine but cannot use because cannot reuse net.Conn
func connect(myPort int, serverPort string, count int, tab string) {
	myAddr := new(net.TCPAddr)
	myIP := "localhost"
	myAddr.IP = net.ParseIP(myIP)
	myAddr.Port = myPort

	serverIP := "localhost"
	s := serverIP + ":" + serverPort
	tcpAddr, err := net.ResolveTCPAddr("tcp", s)
	checkError(err)

	//	fmt.Printf("m: %s\tt: %s\n", myAddr, tcpAddr)

	conn, err := net.DialTCP("tcp", myAddr, tcpAddr)
	checkError(err)

	tmp := make([]byte, 256)
	//	index := 0
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error: ", err)
			}
			break
		}
		//		fmt.Printf("%s ", tmp[:n])
		fmt.Printf("%s%s", tab, tmp[:n])
		//		index++
		if lfindex%3 == 0 {
			//			if index%3 == 0 {
			fmt.Println()
		}
		lfindex++
		//		time.Sleep(time.Second)
	}
	defer conn.Close()

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
