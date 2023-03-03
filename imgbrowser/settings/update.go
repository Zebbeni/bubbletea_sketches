package settings

func (m Model) handleEsc() Model {
	m.ShouldClose = true
	return m
}
