package main

import (
  "github.com/charmbracelet/bubbletea"
  "time"
)

type MODES int
const (
  CM_CLOCK = iota
  CM_MENU
)

type Settings struct{
  clock_size [2]int
}

type model struct {
  current_mode MODES
  currenct_clock clock
  settings Settings
}

func InitialModel() model {
  // Get Current time
  cur_time := time.Now()
  return model{
    CM_CLOCK,
    clock{hour: cur_time.Hour(), minute: cur_time.Minute(), seconds: cur_time.Second()},
    Settings{clock_size: [2]int{5,32}},
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
  // Use Clock to get current time
  // Update the clock
  time_now := time.Now()
  m.currenct_clock.hour = time_now.Hour()
  m.currenct_clock.minute = time_now.Minute()
  m.currenct_clock.seconds = time_now.Second()


  // Render the clock.
  final_time_str := m.currenct_clock.get_string()
  return final_time_str
}
