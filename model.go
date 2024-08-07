package main

import (
  "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
  "time"
  "os"
  "golang.org/x/term"
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
  current_clock clock
  settings Settings
  start_time time.Time
  ticker *time.Ticker
  duration time.Duration
}
type tickMsg time.Time

func InitialModel() model {
  // Get Current time
  cur_time := time.Now()
  return model{
    current_mode: CM_CLOCK,
    current_clock: clock{
      hour: cur_time.Hour(),
      minute: cur_time.Minute(),
      seconds: cur_time.Second(),
    },
    settings: Settings{clock_size: [2]int{5,32}},
    start_time: time.Now(),
  }
}

func (m model) Init() tea.Cmd {
  //Create Ticker
  m.ticker = time.NewTicker(time.Second)
  return tickCmd(m.ticker)
}
func tickCmd(ticker *time.Ticker) tea.Cmd {
	return func() tea.Msg {
		return tickMsg(<-ticker.C)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tickMsg:
    m.duration = time.Since(m.start_time)
    return m, tickCmd(m.ticker)
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
  // time_now := time.Now()
  // m.currenct_clock.hour = time_now.Hour()
  // m.currenct_clock.minute = time_now.Minute()
  // m.currenct_clock.seconds = time_now.Second()

  // New Clock based on ticker
  duration := int(m.duration.Seconds())
  debug_logger.Println("Duration: ", duration)
  duration_clock := clock{hour: duration / 60 / 60, minute: duration / 60 % 60, seconds: duration % 60}
  debug_logger.Println("Duration Clock: ", duration_clock)
  // Render the clock.
  final_time_str := duration_clock.get_string()

  //Get Terminal Dimensions
  var width, height int
  if term.IsTerminal(int(os.Stdout.Fd())) {
    _w, _h, _e := term.GetSize(int(os.Stdout.Fd()))
    width = _w - 2
    height = _h - 2
    if _e != nil {
      panic(_e)
    }
  }

  centered := lipgloss.NewStyle().
    Width(width).
    Height(height).
    Align(lipgloss.Center, lipgloss.Center).
    Padding(1).
    Border(lipgloss.RoundedBorder()).
    Render(final_time_str)

  return centered
}
