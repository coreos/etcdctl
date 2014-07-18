package main

import (
	"os/exec"
	"regexp"
	"testing"
)

const (
	CMD = "etcdctl"
)

type testFormat struct {
	stdoutRegX *regexp.Regexp
	// stderrRegX  *regexp.Regexp
	commandLine *exec.Cmd
}

var testCases []testFormat = []testFormat{
	testFormat{
		stdoutRegX: regexp.MustCompile("([a-z][a-z0-9]*)*"),
		// stderrRegX:  regexp.MustCompile(""),
		commandLine: exec.Command(CMD, "set", "/foo", "bar"),
	},
}

func TestAll(t *testing.T) {
	for index, tst := range testCases {
		cmd := tst.commandLine
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			t.Fatal(err)
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			t.Fatal(err)
		}

		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}

		buf := make([]byte, 100, 100)
		n, err := stderr.Read(buf)
		if n > 0 {
			t.Fail()
			t.Log(CToGoString(buf[:n]))
		}

		if err != nil && n > 0 {
			t.Fatal(err)
		}

		bufO := make([]byte, 100, 100)
		stdout.Read(bufO)

		if tst.stdoutRegX.Match(bufO) == false {
			t.Fail()
			t.Logf("Stdout pattern does not match for test number %d", index)
		}

		if err := cmd.Wait(); err != nil {
			t.Fatal(err)
		}

	}
}

// functional tests: requires you to have etcd running on your machine locally.
// also requires a etcdctl executable which is located in your $GOBIN ( which should be a part of your #PATH)
// further assumption is that etcd starts of without carrying over any information from previous runs.

//Will test the simple: etcdctl ls and etcdctl ls --recursive versions of the flag.
// func TestLscommandEmpty(t *testing.T) {
// 	output, err := exec.Command(CMD, "ls").Output()
// 	if len(output) != 0 {
// 		fmt.Println(output)
// 		t.Error("Error! reported number of keys > 0 but no keys have been added so far")
// 	}

// 	if err != nil {
// 		t.Error("etcdctl ls encountered an error ! ")
// 	}
// }

// func TestGetcommand(t *testing.T) {

// 	cmd := exec.Command("etcdctl", "get", "/coreOS/keys/barz")
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	stderr, err := cmd.StderrPipe()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if err := cmd.Start(); err != nil {
// 		t.Fatal(err)
// 	}

// 	buf := make([]byte, 100, 100)
// 	n, err := stderr.Read(buf)
// 	if n > 0 {
// 		t.Error(CToGoString(buf[:n]))
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	n, _ = stdout.Read(buf)

// 	if n > 0 {
// 		fmt.Println("success: ")
// 		fmt.Println(CToGoString(buf[:n]))

// 	}
// 	if err := cmd.Wait(); err != nil {
// 		t.Fatal(err)
// 	}

// }

// from stackoverflow:
// http://stackoverflow.com/questions/14230145/what-is-the-best-way-to-convert-byte-array-to-string
func CToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

//get prints its errors to Stderr and output to Stdout.
