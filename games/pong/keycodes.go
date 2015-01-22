package main

// import "fmt"

// type Keycode int

// const (
// 	Key_None Keycode = 0

// 	Key_Escape    = 0x01000000
// 	Key_Tab       = 0x01000001
// 	Key_Backtab   = 0x01000002
// 	Key_Backspace = 0x01000003
// 	Key_Return    = 0x01000004
// 	Key_Enter     = 0x01000005 //  Typically located on the keypad.
// 	Key_Delete    = 0x01000007
// 	Key_SysReq    = 0x0100000a

// 	Key_Left  = 0x01000012
// 	Key_Up    = 0x01000013
// 	Key_Right = 0x01000014
// 	Key_Down  = 0x01000015

// 	Key_Shift      = 0x01000020
// 	Key_Control    = 0x01000021 //  On Mac OS X, this corresponds to the Command keys.
// 	Key_Meta       = 0x01000022 //  On Mac OS X, this corresponds to the Control keys. On Windows keyboards, this key is mapped to the Windows key.
// 	Key_Alt        = 0x01000023
// 	Key_AltGr      = 0x01001103 //  On Windows, when the KeyDown event for this key is sent, the Ctrl+Alt modifiers are also set.
// 	Key_CapsLock   = 0x01000024
// 	Key_NumLock    = 0x01000025
// 	Key_ScrollLock = 0x01000026

// 	Key_F1  = 0x01000030
// 	Key_F2  = 0x01000031
// 	Key_F3  = 0x01000032
// 	Key_F4  = 0x01000033
// 	Key_F5  = 0x01000034
// 	Key_F6  = 0x01000035
// 	Key_F7  = 0x01000036
// 	Key_F8  = 0x01000037
// 	Key_F9  = 0x01000038
// 	Key_F10 = 0x01000039
// 	Key_F11 = 0x0100003a
// 	Key_F12 = 0x0100003b

// 	Key_Space = 0x20
// 	Key_Any   = Key_Space

// 	Key_0 = 0x30
// 	Key_1 = 0x31
// 	Key_2 = 0x32
// 	Key_3 = 0x33
// 	Key_4 = 0x34
// 	Key_5 = 0x35
// 	Key_6 = 0x36
// 	Key_7 = 0x37
// 	Key_8 = 0x38
// 	Key_9 = 0x39

// 	Key_A = 0x41
// 	Key_B = 0x42
// 	Key_C = 0x43
// 	Key_D = 0x44
// 	Key_E = 0x45
// 	Key_F = 0x46
// 	Key_G = 0x47
// 	Key_H = 0x48
// 	Key_I = 0x49
// 	Key_J = 0x4a
// 	Key_K = 0x4b
// 	Key_L = 0x4c
// 	Key_M = 0x4d
// 	Key_N = 0x4e
// 	Key_O = 0x4f
// 	Key_P = 0x50
// 	Key_Q = 0x51
// 	Key_R = 0x52
// 	Key_S = 0x53
// 	Key_T = 0x54
// 	Key_U = 0x55
// 	Key_V = 0x56
// 	Key_W = 0x57
// 	Key_X = 0x58
// 	Key_Y = 0x59
// 	Key_Z = 0x5a
// )

// func (k keycode) String() string {
// 	switch k {
// 	default:
// 		return fmt.Sprintf("keycode: %v", int(k))

// 	case Key_Escape:
// 		return "Key_Escape"
// 	case Key_Tab:
// 		return "Key_Tab"
// 	case Key_Backtab:
// 		return "Key_Backtab"
// 	case Key_Backspace:
// 		return "Key_Backspace"
// 	case Key_Return:
// 		return "Key_Return"
// 	case Key_Enter:
// 		return "Key_Enter"
// 	case Key_Delete:
// 		return "Key_Delete"
// 	case Key_SysReq:
// 		return "Key_SysReq"

// 	case Key_Left:
// 		return "Key_Left"
// 	case Key_Up:
// 		return "Key_Up"
// 	case Key_Right:
// 		return "Key_Right"
// 	case Key_Down:
// 		return "Key_Down"

// 	case Key_Shift:
// 		return "Key_Shift"
// 	case Key_Control:
// 		return "Key_Control"
// 	case Key_Meta:
// 		return "Key_Meta"
// 	case Key_Alt:
// 		return "Key_Alt"
// 	case Key_AltGr:
// 		return "Key_AltGr"
// 	case Key_CapsLock:
// 		return "Key_CapsLock"
// 	case Key_NumLock:
// 		return "Key_NumLock"
// 	case Key_ScrollLock:
// 		return "Key_ScrollLock"

// 	case Key_F1:
// 		return "Key_F1"
// 	case Key_F2:
// 		return "Key_F2"
// 	case Key_F3:
// 		return "Key_F3"
// 	case Key_F4:
// 		return "Key_F4"
// 	case Key_F5:
// 		return "Key_F5"
// 	case Key_F6:
// 		return "Key_F6"
// 	case Key_F7:
// 		return "Key_F7"
// 	case Key_F8:
// 		return "Key_F8"
// 	case Key_F9:
// 		return "Key_F9"
// 	case Key_F10:
// 		return "Key_F10"
// 	case Key_F11:
// 		return "Key_F11"
// 	case Key_F12:
// 		return "Key_F12"

// 	case Key_Space:
// 		return "Key_Space"

// 	case Key_0:
// 		return "Key_0"
// 	case Key_1:
// 		return "Key_1"
// 	case Key_2:
// 		return "Key_2"
// 	case Key_3:
// 		return "Key_3"
// 	case Key_4:
// 		return "Key_4"
// 	case Key_5:
// 		return "Key_5"
// 	case Key_6:
// 		return "Key_6"
// 	case Key_7:
// 		return "Key_7"
// 	case Key_8:
// 		return "Key_8"
// 	case Key_9:
// 		return "Key_9"

// 	case Key_A:
// 		return "Key_A"
// 	case Key_B:
// 		return "Key_B"
// 	case Key_C:
// 		return "Key_C"
// 	case Key_D:
// 		return "Key_D"
// 	case Key_E:
// 		return "Key_E"
// 	case Key_F:
// 		return "Key_F"
// 	case Key_G:
// 		return "Key_G"
// 	case Key_H:
// 		return "Key_H"
// 	case Key_I:
// 		return "Key_I"
// 	case Key_J:
// 		return "Key_J"
// 	case Key_K:
// 		return "Key_K"
// 	case Key_L:
// 		return "Key_L"
// 	case Key_M:
// 		return "Key_M"
// 	case Key_N:
// 		return "Key_N"
// 	case Key_O:
// 		return "Key_O"
// 	case Key_P:
// 		return "Key_P"
// 	case Key_Q:
// 		return "Key_Q"
// 	case Key_R:
// 		return "Key_R"
// 	case Key_S:
// 		return "Key_S"
// 	case Key_T:
// 		return "Key_T"
// 	case Key_U:
// 		return "Key_U"
// 	case Key_V:
// 		return "Key_V"
// 	case Key_W:
// 		return "Key_W"
// 	case Key_X:
// 		return "Key_X"
// 	case Key_Y:
// 		return "Key_Y"
// 	case Key_Z:
// 		return "Key_Z"
// 	}
// }
