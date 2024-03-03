package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type model struct {
	projects []string
	cursor   int
	selected map[int]struct{}

	// where to store data
	data_dir string

	// timer
	timer_duration time.Duration
	timer          timer.Model
	last_update    time.Time
}

func main() {
	// set the timer duration
	timer_duration := 1 * time.Minute

	// set data directory
	data_dir := "./data/"

	// read projects
	f, err := os.Open("projects.txt")
	data, err := io.ReadAll(f)
	if err != nil {
		//log.Fatalln("unable to open projects.txt")
		log.Error("unable to open projects.txt")
	}

	// split the projects input into an individual string per project
	split_strings := strings.Split(string(data), "\n")
	projects := make([]string, 0)

	// normalize the inputs, trim the strings and drop empty lines
	for _, v := range split_strings {
		v = strings.TrimSpace(v)
		if v != "" {
			projects = append(projects, v)
		}
	}

	// initialize the models
	m := initModel(projects, timer_duration, data_dir)

	// start the timer
	m.timer.Start()

	// run the model
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initModel(projects []string, timer_duration time.Duration, data_dir string) model {
	return model{
		projects: projects,
		cursor:   0,
		selected: make(map[int]struct{}),

		data_dir: data_dir,

		timer_duration: timer_duration,
		timer:          timer.NewWithInterval(timer_duration, time.Second),
		last_update:    time.Now(),
	}
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "down", "j":
				m.cursor = (m.cursor + len(m.projects) + 1) % len(m.projects)
			case "up", "k":
				m.cursor = (m.cursor + len(m.projects) - 1) % len(m.projects)
			case "enter", " ":
				if _, ok := m.selected[m.cursor]; ok {
					m.selected = make(map[int]struct{})
				} else {
					m.selected = make(map[int]struct{})
					m.selected[m.cursor] = struct{}{}
				}
				err := write_event(m.projects, m.selected, m.data_dir)
				if err != nil {
					log.Error(err)
				}
			}
			// write update
		}
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case timer.TimeoutMsg:
		// reset timer
		m.timer.Timeout = m.timer_duration

		// if last update was longer than 5.5 minutes ago, reset selected and return
		if time.Since(m.last_update) > 5*time.Minute+30*time.Second {
			log.Error("time since last sync > 5.5 minutes, resetting timer")
			m.selected = make(map[int]struct{})
		}

		m.last_update = time.Now()

		// write update
		err := write_event(m.projects, m.selected, m.data_dir)
		if err != nil {
			log.Error(err)
		}

		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("Tani's Active Time Tracker: %s\n", m.timer.View())

	// Iterate over our choices
	for i, choice := range m.projects {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

type Data struct {
	Projects []string
	Selected string
	Time     time.Time
}

func write_event(projects []string, selectedMap map[int]struct{}, directory string) error {
	selected := ""
	for k := range selectedMap {
		selected = projects[k]
	}
	data := &Data{
		projects,
		selected,
		time.Now(),
	}
	file := fmt.Sprintf("test-%d.txt", time.Now().UnixMilli())
	b, _ := json.MarshalIndent(data, "", "  ")
	return os.WriteFile(directory+file, b, 0440)
}
