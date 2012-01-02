/*
Copyright 2011 Paul Ruane.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
    "errors"
    "fmt"
)

type FilesCommand struct{}

func (this FilesCommand) Name() string {
	return "files"
}

func (this FilesCommand) Summary() string {
	return "lists files with particular tags"
}

func (this FilesCommand) Help() string {
	return `  tmsu files TAG...
  tmsu files --all

Lists the files, if any, that have all of the TAGs specified.

  --all    show the complete set of tagged files`
}

func (this FilesCommand) Exec(args []string) error {
    argCount := len(args)

    if argCount == 1 && args[0] == "--all" {
        return this.listAllFiles()
    }

    return this.listFiles(args)
}

func (this FilesCommand) listAllFiles() error {
    db, error := OpenDatabase(databasePath())
    if error != nil { return error }
    defer db.Close()

    files, error := db.Files()
    if error != nil { return error }

    for _, file := range files {
        fmt.Println(file.Path())
    }

    return nil
}

func (this FilesCommand) listFiles(tagNames []string) error {
    if len(tagNames) == 0 { return errors.New("At least one tag must be specified.") }

    db, error := OpenDatabase(databasePath())
    if error != nil { return error }
    defer db.Close()

    files, error := db.FilesWithTags(tagNames)
    if error != nil { return error }

    for _, file := range files {
        fmt.Println(file.Path())
    }

    return nil
}
