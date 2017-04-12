package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	size      int
	dir       string
	cmd       string
	id        string
	keepAlive bool
	workers   int
)

func main() {
	flag.IntVar(&size, "size", 50, "Max size for log file")
	flag.StringVar(&dir, "dir", "./", "Location to save log file")
	flag.StringVar(&cmd, "cmd", "", "Command to run")
	flag.StringVar(&id, "id", "kj", "Identifer of the command to run")
	flag.BoolVar(&keepAlive, "keep-alive", true, "Should kj restart the process?")
	flag.IntVar(&workers, "workers", 1, "How many workers we should spawn")
	flag.Parse()

	// channel to catch nohup
	sigs := make(chan os.Signal)

	// notify channel
	signal.Notify(sigs, syscall.SIGHUP)

	// go catch the signals
	go nohup(sigs)

	// check defaults
	if cmd == "" && len(os.Args) >= 2 {
		cmd = strings.Join(os.Args[1:], " ")
	}

	// check for a valid command
	if cmd == "" {
		fmt.Println("No command to run.")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for worker := 1; worker <= workers; worker++ {
		wg.Add(1)
		go func(worker int) {
			Run(worker, dir, id, cmd)
			wg.Done()
		}(worker)
	}

	wg.Wait()
}

// Run will run the command
func Run(worker int, dir, id, cmd string) {
	for {
		// Make sure we can create the log directory
		if dirErr := os.MkdirAll(dir, 0755); dirErr != nil {
			fmt.Println(dirErr.Error())
			os.Exit(1)
		}

		// execute the command
		command := exec.Command("bash", "-c", cmd)

		log := ""
		if worker != 1 {
			log = "-" + strconv.Itoa(worker)
		}
		// open the out file for writing
		output, _ := os.Create(dir + id + log + ".log")
		defer output.Close()

		// capture stdin/stdout
		command.Stdout = output
		command.Stderr = output
		command.Start()
		command.Wait()

		// if keepalive, run again
		if keepAlive {
			// sleep a second, lets not kill the system
			<-time.After(1 * time.Second)
			continue
		}
		break
	}

}
