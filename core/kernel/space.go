package kernel

// Spaces are containers of Gocs and Systems and can be treated like levels in
// a game context.
//
type Space interface {
	Init()
	DeInit()
	Update()
	AddGoc(goc *Goc)
}
