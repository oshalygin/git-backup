package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Required Flags", func() {

		g.It("should throw an error if the path flag was not set", func() {

			flag := ""
			pathArgument := ""
			expected := "The \"path\" flag must be set"

			actual := checkRequiredFlags(pathArgument, flag).Error()
			g.Assert(actual).Equal(expected)
		})

		g.It("should set error to nil if the path flag was set", func() {

			path := "./projects"
			pathArgument := ""

			actual := checkRequiredFlags(pathArgument, path)
			g.Assert(actual).Equal(nil)
		})

		g.It("should set error to nil if the path argument is passed in", func() {

			path := ""
			pathArgument := "./projects"

			actual := checkRequiredFlags(pathArgument, path)
			g.Assert(actual).Equal(nil)
		})

	})

}
