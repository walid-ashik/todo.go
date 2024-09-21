package main

func main() {
  storage := NewStorage[Todos]("todos.json")
  todos := Todos{}
  storage.Load(&todos)

  cmdFlags := NewCmdFlgas()
  cmdFlags.Execute(&todos)

  storage.Save(todos)
}
