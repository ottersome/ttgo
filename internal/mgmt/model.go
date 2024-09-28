package mgmt

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

type SW_STATES int
const (
  SW_STOPPED = iota
  SW_RUNNING
  SW_PAUSED
)

type Settings struct{
  clock_size [2]int
  terminal_colors map[SW_STATES]string
}
type Stopwatch struct{
  last_tick_time time.Time
  ticker *time.Ticker
  duration time.Duration
  state SW_STATES
}

type model struct {
  current_mode MODES
  current_clock clock
  settings Settings
  stopwatch Stopwatch
}
type tickMsg time.Time

//TODO:
func getTerminalColors(){
  // Get Terminal Colors

}

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
    settings: Settings{
      clock_size: [2]int{5,32},
      terminal_colors:
        map[SW_STATES]string{
          SW_RUNNING: "#B7E8992",
          SW_PAUSED: "#ff6666",
          SW_STOPPED: "#ff6666",
      },
    },
    stopwatch: Stopwatch{
      last_tick_time: time.Now(),
      state: SW_RUNNING,
    },
  }
}

func (m model) Init() tea.Cmd {
  //Create Ticker
  m.stopwatch.ticker = time.NewTicker(time.Second)
  return tickCmd(m.stopwatch.ticker)
}

func tickCmd(ticker *time.Ticker) tea.Cmd {
	return func() tea.Msg {
		return tickMsg(<-ticker.C)
	}
}
func toggleStopwatchCmd(sw *Stopwatch) tea.Cmd {
  switch sw.state {
  case SW_RUNNING:
    sw.state = SW_PAUSED
    if sw.ticker != nil {
      sw.ticker.Stop()
    }
    return nil
  case SW_PAUSED:
    sw.state = SW_RUNNING
    //Continue the ticker
    sw.last_tick_time = time.Now()
    sw.ticker = time.NewTicker(time.Second)
    return tickCmd(sw.ticker)
  case SW_STOPPED:
    sw.state = SW_RUNNING
    sw.last_tick_time = time.Now()
    sw.ticker = time.NewTicker(time.Second)
    return tickCmd(sw.ticker)
  }
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tickMsg:
    m.stopwatch.duration += time.Since(m.stopwatch.last_tick_time)
    m.stopwatch.last_tick_time = time.Now()
    m.stopwatch.ticker = time.NewTicker(time.Second)
    return m, tickCmd(m.stopwatch.ticker)
  case tea.KeyMsg:
    switch msg.String() {
    case "q":
      return m, tea.Quit
    case " ":
      return m, toggleStopwatchCmd(&m.stopwatch)
    }
  }

  var cmds []tea.Cmd
  return m, tea.Batch(cmds...)
}


func (m model) View() string {
  // New Clock based on ticker
  duration := int(m.stopwatch.duration.Seconds())
  duration_clock := clock{hour: duration / 60 / 60, minute: duration / 60 % 60, seconds: duration % 60}
  // Render the clock.
  final_time_str := duration_clock.get_string(false)

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
  // width = 50
  // height = 7

  centered_watch := lipgloss.NewStyle().
    Width(width).
    Height(height).
    Align(lipgloss.Center, lipgloss.Center).
    Padding(1).
    Border(lipgloss.RoundedBorder()).
    Foreground(
      // Something like ternary depending on sw state
      lipgloss.Color(
        m.settings.terminal_colors[m.stopwatch.state],
      ),
    ).
    Render(final_time_str)

  return centered_watch
}
