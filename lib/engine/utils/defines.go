package utils

type ReturnCode int

const (
	// ES_Success indicated that the running game successfully ran.
	ES_Success ReturnCode = iota

	// ES_Restart is used to indicate that the running game should be restarted.
	ES_Restart
)

func (r ReturnCode) String() string {
	switch r {
	default:
		return "Return Code: Unknown"
	case ES_Restart:
		return "Return Code: Restart"
	case ES_Success:
		return "Return Code: Success"
	}
}

const Epsilon = 0.025
