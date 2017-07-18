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

			_, err := checkRequiredFlags(pathArgument, flag)
			actual := err.Error()

			g.Assert(actual).Equal(expected)
		})

		g.It("should set error to nil if the path flag was set", func() {

			path := "./projects"
			pathArgument := ""

			_, actual := checkRequiredFlags(pathArgument, path)
			g.Assert(actual).Equal(nil)
		})

		g.It("should set error to nil if the path argument is passed in", func() {

			path := ""
			pathArgument := "./projects"

			_, actual := checkRequiredFlags(pathArgument, path)
			g.Assert(actual).Equal(nil)
		})

		g.It("should set the pathDirectory to the pathArgument that was passed in", func() {

			path := ""
			pathArgument := "./projects"

			expected := pathArgument

			actual, _ := checkRequiredFlags(pathArgument, path)
			g.Assert(actual).Equal(expected)
		})

		g.It("should set the pathDirectory to the path that was passed in", func() {

			path := "./projects"
			pathArgument := ""

			expected := path

			actual, _ := checkRequiredFlags(pathArgument, path)
			g.Assert(actual).Equal(expected)
		})

		g.It("should set the pathDirectory to the path that was passed in if it matches pathArgument", func() {

			path := "./projects"
			pathArgument := "./projects"

			expected := path

			actual, _ := checkRequiredFlags(pathArgument, path)
			g.Assert(actual).Equal(expected)
		})

		g.It("should throw an error if the path and pathArgument are passed in but they're different values", func() {

			path := "./projects"
			pathArgument := "./projects/oshalygin"

			expected := "Cannot set a path flag AND an argument"

			_, err := checkRequiredFlags(pathArgument, path)

			actual := err.Error()
			g.Assert(actual).Equal(expected)
		})

	})

}
