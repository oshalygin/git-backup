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
			expected := "The \"path\" flag must be set"

			actual := checkRequiredFlags(flag).Error()
			g.Assert(actual).Equal(expected)
		})

		g.It("should set error to nil if the path flag was set", func() {

			flag := "./projects"

			actual := checkRequiredFlags(flag)
			g.Assert(actual).Equal(nil)
		})

	})

}
