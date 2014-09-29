package tok

type Token int

const (
	keywords_begin Token = iota + 1
	For
	If // OMIT
	keywords_end
	operators_begin
	Mul
	Sub
	operators_end
)

func (t Token) String() string {
	switch t {
	default: // OMIT
		return "unknown..." // OMIT
	case For:
		return "for"
	case If: // OMIT
		return "if" // OMIT
	case Mul:
		return "*"
	case Sub: // OMIT
		return "-" // OMIT
	}
}

func IsKeyword(t Token) bool { return t > keywords_begin && t < keywords_end }
