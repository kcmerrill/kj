package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

var (
	size      int
	dir       string
	cmd       string
	id        string
	keepAlive bool
)

func main() {
	flag.IntVar(&size, "size", 50, "Max size for log file")
	flag.StringVar(&dir, "dir", "./", "Location to save log file")
	flag.StringVar(&cmd, "cmd", "", "Command to run")
	flag.StringVar(&id, "id", "kj", "Identifer of the command to run")
	flag.BoolVar(&keepAlive, "keep-alive", false, "Should kj restart the process?")
	flag.Parse()

	// channel to catch nohup
	sigs := make(chan os.Signal)

	// notify channel
	signal.Notify(sigs, syscall.SIGHUP)

	// go catch the signals
	go catchSigs(sigs)

	// check defaults
	if cmd == "" && len(os.Args) >= 2 {
		cmd = strings.Join(os.Args[1:], " ")
	}

	// check for a valid command
	if cmd == "" {
		fmt.Println("No command to run.")
		os.Exit(1)
	}

	for {
	// Make sure we can create the log directory
		if dirErr := os.MkdirAll(dir, 0755); dirErr != nil {
			fmt.Println(dirErr.Error())
			os.Exit(1)
		}

		// execute the command
		command := exec.Command("bash", "-c", cmd)

		// open the out file for writing
		output, _ := os.Create(dir + id + ".log")
		defer output.Close()

		// capture stdin/stdout
		command.Stdout = output
		command.Stderr = output
		command.Start()
		command.Wait()

		// if keepalive, run again
		if keepAlive {
			continue
		}
		break
	}
}

func catchSigs(sigs chan os.Signal) {
	for {
		select {
		// catch the hangup signal for as long as kj is running
		case <-sigs:
			continue
		}
	}
}
