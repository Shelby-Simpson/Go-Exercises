package Entity

type Entity struct {
	ID         int
	Address    string
	EntityType string
}

func (e *Entity) ChangeAddress(newAddr string) {
	e.Address = newAddr
}
