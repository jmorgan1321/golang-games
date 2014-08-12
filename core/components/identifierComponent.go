package components

// Identifier types store basic (uniquely) identifying information about themselves.
type Identifier interface {
	Name() string
	SetName(string)
	// Type() string
}

// IdentifierComponent exists to allow a GOC to be identify itself.
type IdentifierComponent struct {
	name string
	ty   string
}

func (id *IdentifierComponent) SetName(s string) {
	id.name = s
}
func (id *IdentifierComponent) Name() string {
	return id.name
}
func (id *IdentifierComponent) Type() string {
	return id.ty
}

func (*IdentifierComponent) IsComponent() {}
