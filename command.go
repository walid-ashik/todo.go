package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Toggle int
	Edit   string
	List   bool
}

func NewCmdFlgas() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "add new todo")
	flag.StringVar(&cf.Edit, "edit", "", "edit existing todo")
	flag.IntVar(&cf.Del, "del", -1, "delete todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "toggle todo")
	flag.BoolVar(&cf.List, "list", false, "show todos")

	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {

	case cf.Add != "":
		todos.add(cf.Add)
		todos.print()

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use index:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid commad")
			os.Exit(1)
		}

		todos.edit(index, parts[1])
		todos.print()

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
		todos.print()

	case cf.Del != -1:
		todos.delete(cf.Del)
		todos.print()

	case cf.List:
		todos.print()

	default:
		fmt.Println("Invalid commands")
	}
}
