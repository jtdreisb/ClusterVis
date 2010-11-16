package main

import (
	"os"
	"net"
	"fmt"
	"flag"
	"time"
)

const (
	millisecond = 1024 * 1024
	maxnode     = 196
	minnode     = 1
)

/* Show is only used by the timer task but might be useful in future if some
 * status is critical; a talkssh can force a redisplay
 */
type Status struct {
	Id   int
	Err  os.Error
	Show bool
}

var verbose bool
/* meaning: 
 * 0xff -- never connected
 * 0 -- connected, lost connection
 * # > 0 to fe -- count of successfull connects
 */
var status chan Status
var low, high int
var timeout int64

func talkssh(which int, timeout int64) os.Error {
	remote := fmt.Sprintf("10.0.0.%d:22", which)
	if verbose {
		fmt.Printf("Node %d starts with %v\n", which, remote)
	}
	for {
		conn, err := net.Dial("tcp4", "", remote)
		if verbose {
			fmt.Printf("Trying to talk to %v(%v): %v\n", which, remote, err)
		}
		if err == nil {
			conn.Close()
		}
		status <- Status{Id: which, Err: err}
		time.Sleep(timeout * millisecond)
	}
	return nil
}
func main() {

	flag.IntVar(&low, "l", 1, "low part of the range")
	flag.IntVar(&high, "h", 196, "high part of the range")
	flag.Int64Var(&timeout, "t", 5000, "time in seconds to wait for the connects")
	flag.BoolVar(&verbose, "v", false, "be noisy")
	flag.Parse()

	if high > maxnode {
		fmt.Printf("Max node is %d\n", maxnode)
	}

	if low > minnode {
		fmt.Printf("Min node is %d\n", minnode)
	}

	fmt.Fprintf(os.Stderr, "Starting webserv\n")
	webchan := StartWebServer()


	nodestatus := make([]byte, high+1)
	webchan <- nodestatus
	status = make(chan Status, high-low+1)
	for i := low; i <= high; i++ {
		go talkssh(i, timeout)
	}
	for  {
		select {
			case s := <-status:
				if s.Err == nil {
					/* but clamp it at 0xfe */
					i := nodestatus[s.Id]
					/* we'll miscount by one; not a big deal */
					if i != 0xfe {
						i++
					}
					nodestatus[s.Id] = i
				} else {
					if nodestatus[s.Id] != 0xff {
						nodestatus[s.Id] = 0
					}
				}
			case webchan <- nodestatus:
		}
	}
}
