package app

func (m Model) leftPanelWidth() int {
	return 25
}

func (m Model) rightPanelWidth() int {
	return m.w - m.leftPanelWidth()
}
