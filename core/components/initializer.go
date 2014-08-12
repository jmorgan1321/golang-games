package components

// Initalizer types can be constructed/destructed.
type Initializer interface {
	Construct() error
	Destruct() error
}
