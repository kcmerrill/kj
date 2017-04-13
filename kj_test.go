package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func run(cmd string, timeout bool) {
	os.Chdir("/tmp")
	command := exec.Command(cmd)
	command.Start()
	if timeout {
		<-time.After(4 * time.Second)
		command.Process.Kill()
	} else {
		command.Wait()
	}
	// give it a second to spawn the process ...
	<-time.After(1 * time.Second)
}

func TestKJBasic(t *testing.T) {
	// first lets test the basics
	run("echo hello world", false)
	// should cause 3 hello worlds to show up
	_, err := os.Stat("/tmp/echo.log")
	if err != nil {
		t.Fatalf("Expecting there to be /tmp/echo.log")
	} else {
		contents, _ := ioutil.ReadFile("/tmp/echo.log")
		lines := strings.Split(string(contents), "\n")
		if len(lines) <= 3 {
			t.Fatalf("Expecting _at least_ 3 lines")
		}
	}
}
