package main

import (
  "fmt"
  "github.com/charmbracelet/bubbletea"
  "github.com/ottersome/ttgo/internal/mgmt"
  // "os"
  // "log"
)


func main() {
  // TOREM: Whenever you dont feel like using the log_file
  // log_file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  // var debug_logger = log.New(log_file, "DEBUG ", log.Ldate | log.Ltime)
  // if err != nil {
  //   log.Fatal(err)
  // }

  p := tea.NewProgram(mgmt.InitialModel())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error: %v\n", err)
  }
}
