package main

import (
  "fmt"
  "github.com/charmbracelet/bubbletea"
  "os"
  "log"
)

var debug_logger *log.Logger

func main() {
  log_file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    log.Fatal(err)
  }
  debug_logger = log.New(log_file, "DEBUG ", log.Ldate | log.Ltime)
  p := tea.NewProgram(InitialModel(),tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error: %v\n", err)
  }
}
