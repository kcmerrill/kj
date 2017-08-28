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
	size    int
	dir     string
	cmd     string
	id      string
	runOnce bool
	bg      bool
	workers int
)

func main() {

	flag.IntVar(&size, "size", 5, "Max size in MB for log file")
	flag.StringVar(&dir, "dir", "./", "Location to save log file")
	flag.StringVar(&id, "id", "", "Identifer of the command to run")
	flag.BoolVar(&runOnce, "run-once", false, "Should kj restart the process if it dies?")
	flag.BoolVar(&bg, "bg", false, "Was kj started in background?")
	flag.IntVar(&workers, "workers", 1, "How many workers we should spawn")
	flag.Parse()

	// channel to catch nohup
	sigs := make(chan os.Signal)

	// notify channel
	signal.Notify(sigs, syscall.SIGHUP)

	// go catch the signals
	go nohup(sigs)

	// check defaults
	if len(flag.Args()) >= 1 && len(os.Args) >= 2 {
		cmd = strings.Join(flag.Args(), " ")
	}

	// check for a valid command
	if cmd == "" {
		fmt.Println("No command to run.")
		os.Exit(1)
	}

	// if id is empty ... lets set it to the command
	if id == "" {
		id = strings.Split(cmd, " ")[0]
	}

	// should we pop it into the background?
	if !bg {
		// we can just hijack the command to add our strings here
		// secret magic sauce
		if runOnce {
			exec.Command("kj", "--bg", "--run-once", "--dir", dir, "--id", id, "--size", strconv.Itoa(size), "--workers", strconv.Itoa(workers), cmd).Start()
		} else {
			exec.Command("kj", "--bg", "--dir", dir, "--id", id, "--size", strconv.Itoa(size), "--workers", strconv.Itoa(workers), cmd).Start()
		}
		// end 1k island dressing
		os.Exit(0)
	}

	// Run our commands ...
	var wg sync.WaitGroup
	for worker := 1; worker <= workers; worker++ {
		wg.Add(1)
		go func(worker, size int) {
			Run(worker, dir, id, cmd, runOnce, size)
			wg.Done()
		}(worker, size)
	}
	wg.Wait()
}

// Run will run the command
func Run(worker int, dir, id, cmd string, keepAlive bool, size int) {
	// setup our logger
	log := ""
	if worker != 1 {
		log = "-" + strconv.Itoa(worker)
	}
	outFile := dir + id + log + ".log"
	// go clean up after ourselves
	go Janitor(outFile, size)

	// giddy up!
	for {
		// Make sure we can create the log directory
		if dirErr := os.MkdirAll(dir, 0755); dirErr != nil {
			fmt.Println(dirErr.Error())
			os.Exit(1)
		}

		// execute the command
		command := exec.Command("bash", "-c", cmd)

		// open the out file for writing
		output, _ := os.OpenFile(outFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		defer output.Close()

		// capture stdin/stdout
		command.Stdout = output
		command.Stderr = output
		command.Start()
		// print out the process id to the log file
		output.WriteString("[kj pid] " + strconv.Itoa(command.Process.Pid) + "\n")

		command.Wait()

		// if run-once, stop ...
		if runOnce {
			break
		}

		// sleep a second, lets not kill the system
		<-time.After(1 * time.Second)
		continue

	}
}

// catch nohup goodness
func nohup(sigs chan os.Signal) {
	for {
		select {
		// catch the hangup signal for as long as kj is running
		case <-sigs:
			continue
		}
	}
}

// Janitor will watch a file and make sure it doesn't go over the specified size
func Janitor(file string, size int) {
	for {
		info, err := os.Stat(file)
		if err == nil {
			if int64(size) <= int64(float64(0.000001)*float64(info.Size())) {
				// we should do something about the file size!
				os.Truncate(file, 0)
			}
		}
		<-time.After(1 * time.Second)
	}
}
