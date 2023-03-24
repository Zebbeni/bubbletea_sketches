package component

type Model struct {
	isActive bool
}

func New() Model {
	return Model{isActive: false}
}

func (c Model) IsActive() bool {
	return c.isActive
}

func (c Model) Activate() {
	c.isActive = true
}

func (c Model) Deactivate() {
	c.isActive = false
}
