package ftpserver

import (
	"net"
	"sync"
	"time"
)

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()

	select {
	case <-c:
		return false
	case <-time.After(timeout):
		return true
	}
}

type Passive struct {
	listenSuccessAt int64
	listenFailedAt  int64
	closeSuccessAt  int64
	listenAt        int64
	connectedAt     int64
	connection      *net.Conn
	cmd             string
	param           string
	cid             string
	port            int
	waiter          sync.WaitGroup
}

func getThatPassiveConnection(passiveListen *net.TCPListener, p *Passive){
	var perr error
	p.connection, perr = passiveListen.AcceptTCP()
	if perr != nil {
		p.listenFailedAt = time.Now().Unix()
		p.waiter.Done()
		return
	}
	passiveListen.Close()
	p.listenSuccessAt = time.Now().Unix()
	p.waiter.Done() 
}

