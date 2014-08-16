package kernel

type State int

const (
	Running State = iota + 1
	Stopped
	Rebooting
	Terminated
)

func (s State) String() string {
	switch s {
	default:
		return "Kernel State: Unknown"
	case Running:
		return "Kernel State: Running"
	case Stopped:
		return "Kernel State: Stopped"
	case Rebooting:
		return "Kernel State: Rebooting"
	case Terminated:
		return "Kernel State: Terminated"
	}
}
