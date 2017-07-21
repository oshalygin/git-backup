package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test_DirectoryService(t *testing.T) {
	g := Goblin(t)
	g.Describe("Directory Service", func() {

		g.It("should return the directories in the current project directory", func() {

			path := "."

			directories := GetDirectoriesInPath(path)
			actual := len(directories) > 0

			g.Assert(actual).Equal(true)
		})

		g.It("should return 'vendor' as one of the directories in the current path", func() {

			path := "."
			expectedDirectory := "vendor"
			containsDirectory := false

			directories := GetDirectoriesInPath(path)
			for _, directory := range directories {
				if directory == expectedDirectory {
					containsDirectory = true
				}
			}

			actual := containsDirectory
			g.Assert(actual).Equal(true)
		})

		g.It("should return '.git' as one of the directories in the current path", func() {

			path := "."
			expectedDirectory := ".git"
			containsDirectory := false

			directories := GetDirectoriesInPath(path)
			for _, directory := range directories {
				if directory == expectedDirectory {
					containsDirectory = true
				}
			}

			actual := containsDirectory
			g.Assert(actual).Equal(true)
		})

	})
}
