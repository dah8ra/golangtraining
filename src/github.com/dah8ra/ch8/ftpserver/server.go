package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/dah8ra/ch8/system"
)

var CommandMap map[string]func(*Ftp)
var ConnectionMap map[string]*Ftp

var FileManager *system.FileManager
var AuthManager *system.AuthManager

func main() {
	start()
}

type Ftp struct {
	writer  *bufio.Writer
	reader  *bufio.Reader
	conn    net.Conn
	waiter  sync.WaitGroup
	user    string
	homeDir string
	path    string
	ip      string
	cmd     string
	param   string
}

func init() {
	CommandMap = make(map[string]func(*Ftp))

	CommandMap["cd"] = (*Ftp).handleCd

	CommandMap["USER"] = (*Ftp).handleUser
	CommandMap["PASS"] = (*Ftp).handlePass
	//	CommandMap["STOR"] = (*Ftp).handleStore
	//	CommandMap["APPE"] = (*Ftp).handleStore
	//	CommandMap["STAT"] = (*Ftp).handleStat
	//	CommandMap["SYST"] = (*Ftp).handleSyst
	CommandMap["PWD"] = (*Ftp).handlePwd
	//	CommandMap["TYPE"] = (*Ftp).handleType
	//	CommandMap["PASV"] = (*Ftp).handlePassive
	//	CommandMap["EPSV"] = (*Ftp).handlePassive
	//	CommandMap["NLST"] = (*Ftp).handleList
	//	CommandMap["LIST"] = (*Ftp).handleList
	CommandMap["QUIT"] = (*Ftp).handleQuit
	//	CommandMap["CWD"] = (*Ftp).handleCwd
	//	CommandMap["SIZE"] = (*Ftp).handleSize
	//	CommandMap["RETR"] = (*Ftp).handleRetr

	ConnectionMap = make(map[string]*Ftp)
}

func NewInstance(conn net.Conn) *Ftp {
	f := Ftp{}
	f.writer = bufio.NewWriter(conn)
	f.reader = bufio.NewReader(conn)
	f.path = "/"
	f.conn = conn
	f.ip = conn.RemoteAddr().String()
	return &f
}

func start() {
	fmt.Println("Starting...")
	url := fmt.Sprintf("localhost:%d", 2121)
	var listener net.Listener
	listener, err := net.Listen("tcp", url)

	if err != nil {
		fmt.Println("cannot listen on: ", url)
		return
	}
	fmt.Println("listening on: ", url)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listening err ", err)
			break
		}
		f := NewInstance(conn)
		go f.handleCommands()
	}
}

func (f *Ftp) handleCommands() {
	for {
		line, err := f.reader.ReadString('\n')
		if err != nil {
			break
		}
		cmd, param := parseLine(line)
		f.cmd = cmd
		f.param = param
		fmt.Println(cmd)
		fn := CommandMap[cmd]
		if fn == nil {
			f.writeMessage(550, "not allowed")
		} else {
			fn(f)
		}
	}
}

func (f *Ftp) writeMessage(code int, message string) {
	line := fmt.Sprintf("%d %s\r\n", code, message)
	f.writer.WriteString(line)
	f.writer.Flush()
}

func (f *Ftp) handleUser() {
	f.user = f.param
	f.writeMessage(331, "User name ok, password required")
}

func (f *Ftp) handlePwd() {
	f.writeMessage(257, "\""+f.path+"\" is the current directory")
}

func (f *Ftp) handlePass() {

}

func (f *Ftp) handleCd() {

}

func (f *Ftp) handleQuit() {
	f.writeMessage(221, "Goodbye")
	f.conn.Close()
	//	delete(ConnectionMap, p.cid)
}

func parseLine(line string) (string, string) {
	params := strings.SplitN(strings.Trim(line, "\r\n"), " ", 2)
	if len(params) == 1 {
		return params[0], ""
	}
	return params[0], strings.TrimSpace(params[1])
}
