package io

import tea "github.com/charmbracelet/bubbletea"

// BackMsg indicates that the currently-viewed experience is handing control back
type BackMsg int

var BackCmd = func() tea.Msg { return BackMsg(1) }
