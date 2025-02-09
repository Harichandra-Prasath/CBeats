// Dumper that dumps all the logs collected to sink

package main

import (
	"fmt"
	"log"
	"net"
)

type Dumper struct {
	tcpConn *net.TCPConn
}

// Creates a New Dumper
func NewDumper() (*Dumper, error) {

	raddr, err := net.ResolveTCPAddr("tcp", "localhost:"+SNK_PORT)
	if err != nil {
		return nil, fmt.Errorf("creating dumper: %s", err)

	}

	newConn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return nil, fmt.Errorf("creating dumper: %s", err)

	}

	dumper := &Dumper{
		tcpConn: newConn,
	}

	return dumper, nil
}

// Core function that dumps all the logs
func (d *Dumper) dumpLogs(logs *Logs) error {

	_, err := d.tcpConn.Write(*logs.data)
	if err != nil {
		return fmt.Errorf("dumping logs: %s", err)
	}

	log.Printf("Logs of batch %d dumped successfully", logs.batch)

	return nil
}
