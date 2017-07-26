package main

import (
	"os"
	"os/exec"
	"testing"

	. "github.com/franela/goblin"
)

// Command spys
var passedInCommand string
var passedInArgs []string

// From Go's source code
// https://golang.org/src/os/exec/exec_test.go
func fakeExecCommand(command string, args ...string) *exec.Cmd {

	passedInCommand = command
	passedInArgs = args

	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func Test_GitService(t *testing.T) {
	g := Goblin(t)
	g.Describe("Git Service", func() {

		g.It("should succeed when calling FetchLatest on the current directory", func() {

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."

			actual := FetchLatest(directory)
			g.Assert(actual).Equal(true)
		})

		g.It("should fail when calling FetchLatest on a directory that does not exist", func() {

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "./foobar"

			actual := FetchLatest(directory)
			g.Assert(actual).Equal(false)
		})

		g.It("should call the git command when calling FetchLatest", func() {

			expected := "git"
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."
			FetchLatest(directory)

			actual := passedInCommand
			g.Assert(actual).Equal(expected)
		})

		g.It("should call the FetchLatest with the args 'fetch' and '-all'", func() {

			expected := []string{"fetch", "--all"}
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."
			FetchLatest(directory)

			actual := passedInArgs
			g.Assert(actual).Equal(expected)
		})

		g.It("should succeed when calling PullLatest on the current directory", func() {

			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			directory := "."
			actual := PullLatest(directory)

			g.Assert(actual).Equal(true)
		})

		g.It("should fail when calling PullLatest on a directory that does not exist", func() {

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "./foobar"

			actual := PullLatest(directory)
			g.Assert(actual).Equal(false)
		})

		g.It("should call the git command when calling PullLatest", func() {

			expected := "git"
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."
			PullLatest(directory)

			actual := passedInCommand
			g.Assert(actual).Equal(expected)
		})

		g.It("should call the PullLatest with the args 'fetch' and '-all'", func() {

			expected := []string{"pull", "--all"}
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."
			PullLatest(directory)

			actual := passedInArgs
			g.Assert(actual).Equal(expected)
		})

		g.It("should succeed when calling PushLatest on the current directory", func() {

			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			directory := "."
			actual := PushLatest(directory)

			g.Assert(actual).Equal(true)
		})

		g.It("should fail when calling PushLatest on a directory that does not exist", func() {

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "./foobar"

			actual := PushLatest(directory)
			g.Assert(actual).Equal(false)
		})

		g.It("should call the git command when calling PushLatest", func() {

			expected := "git"
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."
			PushLatest(directory)

			actual := passedInCommand
			g.Assert(actual).Equal(expected)
		})

		g.It("should call the PushLatest with the args 'push' 'bitbucket' and '-all'", func() {

			expected := []string{"push", "bitbucket", "--all"}
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."
			PushLatest(directory)

			actual := passedInArgs
			g.Assert(actual).Equal(expected)
		})

	})
}
