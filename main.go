package main

import (
  "fmt"
  "github.com/charmbracelet/bubbletea"
)

func main() {
  p := tea.NewProgram(InitialModel(),tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error: %v\n", err)
  }
}
