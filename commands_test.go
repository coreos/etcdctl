package main

import (
	"os/exec"
	"regexp"
	"testing"
)

const (
	CMD = "etcdctl"
)

//any text in stderr is bad , and hence no regX required to capture that.
type testFormat struct {
	stdoutRegX  *regexp.Regexp
	commandLine *exec.Cmd
}

//Testing story so far:
/*
-- set a key
-- get a key.
-- ls and ls --recursive.
-- test update
-- test get --consistent
-- test mkdir.
-- test rmdir.
*/

var keyPattern string = "^[a-z]([A-Za-z0-9]*)"

// key rules assumed only for testing: starts wiht lower case, after that can contain letters and numbers.
var lsRecursivePattern string = "^(/([a-z][A-Za-z0-9]*)*\n?)+"

// recursiveDirPattern  is of the form (/keypattern)+
var lsPattern string = "^(/([a-z]([A-Za-z0-9]*))\n)+(/[a-z]([A-Za-z0-9]*))"

var testCases []testFormat = []testFormat{
	testFormat{
		stdoutRegX:  regexp.MustCompile(keyPattern),
		commandLine: exec.Command(CMD, "set", "/foo", "bar"),
	},

	testFormat{
		stdoutRegX:  regexp.MustCompile(keyPattern),
		commandLine: exec.Command(CMD, "set", "/coreOS/keys/dog", "coreo"),
	},

	testFormat{

		stdoutRegX:  regexp.MustCompile(keyPattern),
		commandLine: exec.Command(CMD, "get", "/foo"),
	},

	testFormat{

		stdoutRegX:  regexp.MustCompile(keyPattern),
		commandLine: exec.Command(CMD, "get", "/coreOS/keys/dog"),
	},

	testFormat{

		stdoutRegX:  regexp.MustCompile(lsRecursivePattern),
		commandLine: exec.Command(CMD, "ls", "--recursive"),
	},

	testFormat{

		stdoutRegX:  regexp.MustCompile(lsPattern),
		commandLine: exec.Command(CMD, "ls"),
	},

	testFormat{
		stdoutRegX:  regexp.MustCompile(keyPattern),
		commandLine: exec.Command(CMD, "update", "/coreOS/keys/dog", "woof"),
	},

	testFormat{
		stdoutRegX:  regexp.MustCompile("^woof"),
		commandLine: exec.Command(CMD, "get", "/coreOS/keys/dog", "--consistent"),
	},

	testFormat{
		stdoutRegX:  regexp.MustCompile(""), // no output on successfull  mkdir
		commandLine: exec.Command(CMD, "mkdir", "/core/dir1"),
	},

	testFormat{
		stdoutRegX:  regexp.MustCompile(""), // no output on successfull  mkdir
		commandLine: exec.Command(CMD, "rmdir", "/core/dir1"),
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
			t.Log(ByteSliceToString(buf[:n]))
		}

		if err != nil && n > 0 {
			t.Fatal(err)
		}

		bufO := make([]byte, 100, 100)
		stdout.Read(bufO)
		// m, _ := stdout.Read(bufO)
		// fmt.Println(ByteSliceToString(bufO[:m]))
		if tst.stdoutRegX.Match(bufO) == false {
			t.Fail()
			t.Logf("Stdout pattern does not match for test number %d", index)
		}

		if err := cmd.Wait(); err != nil {
			t.Fatal(err)
		}

	}
}

// from stackoverflow:
// http://stackoverflow.com/questions/14230145/what-is-the-best-way-to-convert-byte-array-to-string
func ByteSliceToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}
