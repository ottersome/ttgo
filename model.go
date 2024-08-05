package main

import (
  "fmt"
  "github.com/charmbracelet/bubbletea"
)

type model struct {
  choices []string
  cursor int
  selected map[int]struct{}
}

func InitialModel() model {
  return model{
    choices: []string{"a", "b", "c"},
    cursor: 0,
    selected: make(map[int]struct{}),
  }
}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "q":
      return m, tea.Quit
    case "down":
      m.cursor++
      if m.cursor >= len(m.choices) {
        m.cursor = 0
      }
    case "up":
      m.cursor--
      if m.cursor < 0 {
        m.cursor = len(m.choices) - 1
      }
    case "enter":
      m.selected[m.cursor] = struct{}{}
    }
  }

  var cmds []tea.Cmd
  // for _, choice := range m.choices {
  //   if _, ok := m.selected[m.cursor]; ok {
  //     // cmds = append(initialModel)
  //   } else {
  //     // cmds = append(cmds, tea.Print(choice + " (unselected)"))
  //   }
  // }
  // cmds = append(cmds, tea.Print(""))
  return m, tea.Batch(cmds...)
}

func (m model) View() string {
  return fmt.Sprintf("Choices: %v\nCursor: %v\nSelected: %v", m.choices, m.cursor, m.selected)
}
