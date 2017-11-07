package main

import (
	"log"
	"sync"

	"bufio"
	"os"
	"syscall"

	"os/signal"

	"sync/atomic"

	"github.com/sryanyuan/tcpnetwork"
	"runtime"
)

var (
	kServerAddress  = "localhost:1444"
	serverConnected int32
	stopFlag        int32
)

//echo server routine
func echoserver() (*tcpnetwork.TCPNetwork, error) {
	//
	var err error
	server := tcpnetwork.NewTCPNetwork(1024, tcpnetwork.NewStreamProtocol4())
	err = server.Listen(kServerAddress)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func routineEchoServer(sever *tcpnetwork.TCPNetwork, wg *sync.WaitGroup, stopCh chan struct{}) {
	defer func() {
		log.Println("server done")
		wg.Done()
	}()

	for {
		select {
		case evt, ok := <-sever.GetEventQueue():
			{
				if !ok {
					return
				}

				switch evt.EventType {
				case tcpnetwork.KConnEvent_Connected:
					{
						log.Println("Client", evt.Conn.GetRemoteAddress(), " connected.")
					}
				case tcpnetwork.KConnEvent_Close:
					{
						log.Println("Client", evt.Conn.GetRemoteAddress(), " dissconnected.")
					}
				case tcpnetwork.KConnEvent_Data:
					{
						evt.Conn.Send(evt.Data, 0)
					}
				}

			}
		case <-stopCh:
			{
				return
			}
		}
	}
}

func echoClient(*tcpnetwork.TCPNetwork, *tcpnetwork.Connection, error) {

}

func routineEchoClient() {

}

func routineInput(wg *sync.WaitGroup, clientConn *tcpnetwork.Connection) {

}

func main() {
	//create server

}
