package main

import (
	"flag"
	"fmt"
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
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.List:
		todos.print()
	default:
		fmt.Println("Invalid commands")
	}
}
