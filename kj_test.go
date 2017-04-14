package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func runHelloWorld() {
	os.Chdir("/tmp")
	command := exec.Command("kj", "echo", "hello", "world")
	command.Stdout = os.Stdout
	command.Start()
	command.Wait()
	// give it a second to spawn the process ...
	<-time.After(1 * time.Second)
}

func TestKJBasic(t *testing.T) {
	// first lets test the basics
	runHelloWorld()
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
