// Dumper that dumps all the logs collected to sink

package main

import (
	"fmt"
	"log"
	"net"
)

type Dumper struct {
	TcpConn *net.TCPConn
}

// Creates a New Dumper
func NewDumper(port string) (*Dumper, error) {

	raddr, err := net.ResolveTCPAddr("tcp", "localhost:"+port)
	if err != nil {
		return nil, fmt.Errorf("creating dumper: %s", err)

	}

	newConn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return nil, fmt.Errorf("creating dumper: %s", err)

	}

	dumper := &Dumper{
		TcpConn: newConn,
	}

	return dumper, nil
}

// Core function that dumps all the logs
func (d *Dumper) dumpLogs(logs *Logs) error {

	_, err := d.TcpConn.Write(*logs.data)
	if err != nil {
		return fmt.Errorf("dumping logs: %s", err)
	}

	log.Printf("Logs of batch %d dumped successfully for %s ", logs.batch, logs.file)

	return nil
}
