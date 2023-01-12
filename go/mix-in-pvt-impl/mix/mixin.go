package mix

type IDer interface {
	InitID(id int)
	GetID() int
}

type IDerMixin struct {
	id int
}

// enforce static constraint:  MixImpl implements Mix
var _ IDer = new(IDerMixin)

func (m *IDerMixin) InitID(id int) {
	m.id = id
}

func (m *IDerMixin) GetID() int {
	return m.id
}
