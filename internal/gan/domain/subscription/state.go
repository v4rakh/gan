package subscription

type State string

const (
	Pending State = "pending"
	Active  State = "active"
)

func (e *State) Scan(value interface{}) error {
	*e = State(value.([]byte))
	return nil
}

func (e State) Value() string {
	return string(e)
}
