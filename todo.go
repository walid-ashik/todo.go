package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
)

type Todo struct {
	Text        string
	Completed   bool
	CompletedAt *time.Time
	CreatedAt   time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Text:        title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(i int) error {
	if i < 0 || i >= len(*todos) {
		err := errors.New("Index not found")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	// delete by index, make 2 slice, merge them
	*todos = append(t[:index], t[:index+1]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Text = title
	return nil
}

func (todos *Todos) print() {
	t := table.New(os.Stdout)
	t.SetHeaders("index", "Do it", "Done")
	t.SetHeaderStyle(table.StyleBold)
	t.SetDividers(table.MarkdownDividers)

	t.SetBorderTop(false)
	t.SetBorderBottom(false)
	t.SetRowLines(false)

	for index, todo := range *todos {
		i := "<red>" + strconv.Itoa(index) + "</red>"
		text := "<red>" + todo.Text + "</red>"
		doneStatus := ""

		if todo.Completed {
			doneStatus = "x"
			i = "<green>" + strconv.Itoa(index) + "</green>"
			text = "<green>" + todo.Text + "</green>"
		}

		t.AddRow(tml.Sprintf(i), tml.Sprintf(text), doneStatus)
	}

	t.Render()
}
