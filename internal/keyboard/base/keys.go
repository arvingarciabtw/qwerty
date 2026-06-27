package base

import evdev "github.com/gvalkov/golang-evdev"

var (
	// ALPHANUMERIC KEYS
	KeyA = Key{Label: "A", Width: U1, Finger: Pinky, EvCode: evdev.KEY_A}
	KeyB = Key{Label: "B", Width: U1, Finger: Index, EvCode: evdev.KEY_B}
	KeyC = Key{Label: "C", Width: U1, Finger: Middle, EvCode: evdev.KEY_C}
	KeyD = Key{Label: "D", Width: U1, Finger: Middle, EvCode: evdev.KEY_D}
	KeyE = Key{Label: "E", Width: U1, Finger: Middle, EvCode: evdev.KEY_E}
	KeyF = Key{Label: "F", Width: U1, Finger: Index, EvCode: evdev.KEY_F}
	KeyG = Key{Label: "G", Width: U1, Finger: Index, EvCode: evdev.KEY_G}
	KeyH = Key{Label: "H", Width: U1, Finger: Index, EvCode: evdev.KEY_H}
	KeyI = Key{Label: "I", Width: U1, Finger: Middle, EvCode: evdev.KEY_I}
	KeyJ = Key{Label: "J", Width: U1, Finger: Index, EvCode: evdev.KEY_J}
	KeyK = Key{Label: "K", Width: U1, Finger: Middle, EvCode: evdev.KEY_K}
	KeyL = Key{Label: "L", Width: U1, Finger: Ring, EvCode: evdev.KEY_L}
	KeyM = Key{Label: "M", Width: U1, Finger: Index, EvCode: evdev.KEY_M}
	KeyN = Key{Label: "N", Width: U1, Finger: Index, EvCode: evdev.KEY_N}
	KeyO = Key{Label: "O", Width: U1, Finger: Ring, EvCode: evdev.KEY_O}
	KeyP = Key{Label: "P", Width: U1, Finger: Pinky, EvCode: evdev.KEY_P}
	KeyQ = Key{Label: "Q", Width: U1, Finger: Pinky, EvCode: evdev.KEY_Q}
	KeyR = Key{Label: "R", Width: U1, Finger: Index, EvCode: evdev.KEY_R}
	KeyS = Key{Label: "S", Width: U1, Finger: Ring, EvCode: evdev.KEY_S}
	KeyT = Key{Label: "T", Width: U1, Finger: Index, EvCode: evdev.KEY_T}
	KeyU = Key{Label: "U", Width: U1, Finger: Index, EvCode: evdev.KEY_U}
	KeyV = Key{Label: "V", Width: U1, Finger: Index, EvCode: evdev.KEY_V}
	KeyW = Key{Label: "W", Width: U1, Finger: Ring, EvCode: evdev.KEY_W}
	KeyX = Key{Label: "X", Width: U1, Finger: Ring, EvCode: evdev.KEY_X}
	KeyY = Key{Label: "Y", Width: U1, Finger: Index, EvCode: evdev.KEY_Y}
	KeyZ = Key{Label: "Z", Width: U1, Finger: Pinky, EvCode: evdev.KEY_Z}
	Key0 = Key{Label: "0", Width: U1, Finger: Pinky, EvCode: evdev.KEY_0}
	Key1 = Key{Label: "1", Width: U1, Finger: Pinky, EvCode: evdev.KEY_1}
	Key2 = Key{Label: "2", Width: U1, Finger: Ring, EvCode: evdev.KEY_2}
	Key3 = Key{Label: "3", Width: U1, Finger: Middle, EvCode: evdev.KEY_3}
	Key4 = Key{Label: "4", Width: U1, Finger: Index, EvCode: evdev.KEY_4}
	Key5 = Key{Label: "5", Width: U1, Finger: Index, EvCode: evdev.KEY_5}
	Key6 = Key{Label: "6", Width: U1, Finger: Index, EvCode: evdev.KEY_6}
	Key7 = Key{Label: "7", Width: U1, Finger: Index, EvCode: evdev.KEY_7}
	Key8 = Key{Label: "8", Width: U1, Finger: Middle, EvCode: evdev.KEY_8}
	Key9 = Key{Label: "9", Width: U1, Finger: Ring, EvCode: evdev.KEY_9}

	// CONTROL KEYS
	KeyEsc        = Key{Label: "Esc", Width: U1, Finger: Pinky, EvCode: evdev.KEY_ESC}
	KeyLeftShift  = Key{Label: "Shift", Width: U2p5, Finger: Pinky, EvCode: evdev.KEY_LEFTSHIFT}
	KeyRightShift = Key{Label: "Shift", Width: U3, Finger: Pinky, EvCode: evdev.KEY_RIGHTSHIFT}
	KeyLeftCtrl   = Key{Label: "Ctrl", Width: U2, Finger: Pinky, EvCode: evdev.KEY_LEFTCTRL}
	KeyRightCtrl  = Key{Label: "Ctrl", Width: U2, Finger: Pinky, EvCode: evdev.KEY_RIGHTCTRL}
	KeyLeftMeta   = Key{Label: "⌘", Width: U1, Finger: Ring, EvCode: evdev.KEY_LEFTMETA}
	KeyRightMeta  = Key{Label: "⌘", Width: U1, Finger: Ring, EvCode: evdev.KEY_RIGHTMETA}
	KeyLeftAlt    = Key{Label: "Alt", Width: U1p75, Finger: Thumb, EvCode: evdev.KEY_LEFTALT}
	KeyRightAlt   = Key{Label: "Alt", Width: U1p75, Finger: Thumb, EvCode: evdev.KEY_RIGHTALT}
	KeyCapsLock   = Key{Label: "Caps", Width: U2, Finger: Pinky, EvCode: evdev.KEY_CAPSLOCK}
	KeyNumLock    = Key{Label: "Nlk", Width: U1, Finger: Index, EvCode: evdev.KEY_NUMLOCK}
	KeyScrollLock = Key{Label: "⚲", Width: U1, Finger: Middle, EvCode: evdev.KEY_SCROLLLOCK}

	// FUNCTION KEYS
	KeyFn  = Key{Label: "Fn", Width: U1, Finger: Ring, EvCode: evdev.KEY_FN}
	KeyF1  = Key{Label: "F1", Width: U1, Finger: Pinky, EvCode: evdev.KEY_F1}
	KeyF2  = Key{Label: "F2", Width: U1, Finger: Ring, EvCode: evdev.KEY_F2}
	KeyF3  = Key{Label: "F3", Width: U1, Finger: Middle, EvCode: evdev.KEY_F3}
	KeyF4  = Key{Label: "F4", Width: U1, Finger: Index, EvCode: evdev.KEY_F4}
	KeyF5  = Key{Label: "F5", Width: U1, Finger: Index, EvCode: evdev.KEY_F5}
	KeyF6  = Key{Label: "F6", Width: U1, Finger: Index, EvCode: evdev.KEY_F6}
	KeyF7  = Key{Label: "F7", Width: U1, Finger: Index, EvCode: evdev.KEY_F7}
	KeyF8  = Key{Label: "F8", Width: U1, Finger: Middle, EvCode: evdev.KEY_F8}
	KeyF9  = Key{Label: "F9", Width: U1, Finger: Ring, EvCode: evdev.KEY_F9}
	KeyF10 = Key{Label: "F10", Width: U1, Finger: Pinky, EvCode: evdev.KEY_F10}
	KeyF11 = Key{Label: "F11", Width: U1, Finger: Pinky, EvCode: evdev.KEY_F11}
	KeyF12 = Key{Label: "F12", Width: U1, Finger: Pinky, EvCode: evdev.KEY_F12}

	// NAVIGATION AND EDITING KEYS
	KeyUp        = Key{Label: "↑", Width: U1, Finger: Middle, EvCode: evdev.KEY_UP}
	KeyDown      = Key{Label: "↓", Width: U1, Finger: Middle, EvCode: evdev.KEY_DOWN}
	KeyLeft      = Key{Label: "←", Width: U1, Finger: Index, EvCode: evdev.KEY_LEFT}
	KeyRight     = Key{Label: "→", Width: U1, Finger: Ring, EvCode: evdev.KEY_RIGHT}
	KeyPageUp    = Key{Label: "⇡", Width: U1, Finger: Ring, EvCode: evdev.KEY_PAGEUP}
	KeyPageDown  = Key{Label: "⇣", Width: U1, Finger: Ring, EvCode: evdev.KEY_PAGEDOWN}
	KeyHome      = Key{Label: "⌂", Width: U1, Finger: Middle, EvCode: evdev.KEY_HOME}
	KeyEnd       = Key{Label: "⌿", Width: U1, Finger: Middle, EvCode: evdev.KEY_END}
	KeyInsert    = Key{Label: "Ins", Width: U1, Finger: Index, EvCode: evdev.KEY_INSERT}
	KeyDelete    = Key{Label: "Del", Width: U1, Finger: Index, EvCode: evdev.KEY_DELETE}
	KeyBackspace = Key{Label: "<--", Width: U2p5, Finger: Pinky, EvCode: evdev.KEY_BACKSPACE}

	// SPECIAL AND PUNCTUATION KEYS
	KeyTab        = Key{Label: "Tab", Width: U1p75, Finger: Pinky, EvCode: evdev.KEY_TAB}
	KeyEnter      = Key{Label: "Enter↵", Width: U3, Finger: Pinky, EvCode: evdev.KEY_ENTER}
	KeySpace      = Key{Label: "Space", Width: U5p75, Finger: Thumb, EvCode: evdev.KEY_SPACE}
	KeyGrave      = Key{Label: "`", Width: U1, Finger: Pinky, EvCode: evdev.KEY_GRAVE}
	KeyMinus      = Key{Label: "-", Width: U1, Finger: Pinky, EvCode: evdev.KEY_MINUS}
	KeyEqual      = Key{Label: "=", Width: U1, Finger: Pinky, EvCode: evdev.KEY_EQUAL}
	KeyLeftBrace  = Key{Label: "[", Width: U1, Finger: Pinky, EvCode: evdev.KEY_LEFTBRACE}
	KeyRightBrace = Key{Label: "]", Width: U1, Finger: Pinky, EvCode: evdev.KEY_RIGHTBRACE}
	KeyBackslash  = Key{Label: "\\", Width: U1p75, Finger: Pinky, EvCode: evdev.KEY_BACKSLASH}
	KeySemicolon  = Key{Label: ";", Width: U1, Finger: Pinky, EvCode: evdev.KEY_SEMICOLON}
	KeyApostrophe = Key{Label: "'", Width: U1, Finger: Pinky, EvCode: evdev.KEY_APOSTROPHE}
	KeyComma      = Key{Label: ",", Width: U1, Finger: Middle, EvCode: evdev.KEY_COMMA}
	KeyDot        = Key{Label: ".", Width: U1, Finger: Ring, EvCode: evdev.KEY_DOT}
	KeySlash      = Key{Label: "/", Width: U1, Finger: Pinky, EvCode: evdev.KEY_SLASH}
	KeyPound      = Key{Label: "#", Width: U1, Finger: Pinky, EvCode: evdev.KEY_NUMERIC_POUND}
	KeyAcute      = Key{Label: "´", Width: U1, Finger: Pinky, EvCode: evdev.KEY_RESERVED}

	// NUMPAD KEYS
	KeyPad0        = Key{Label: "0", Width: U2p5, Finger: Thumb, EvCode: evdev.KEY_KP0}
	KeyPad1        = Key{Label: "1", Width: U1, Finger: Index, EvCode: evdev.KEY_KP1}
	KeyPad2        = Key{Label: "2", Width: U1, Finger: Middle, EvCode: evdev.KEY_KP2}
	KeyPad3        = Key{Label: "3", Width: U1, Finger: Ring, EvCode: evdev.KEY_KP3}
	KeyPad4        = Key{Label: "4", Width: U1, Finger: Index, EvCode: evdev.KEY_KP4}
	KeyPad5        = Key{Label: "5", Width: U1, Finger: Middle, EvCode: evdev.KEY_KP5}
	KeyPad6        = Key{Label: "6", Width: U1, Finger: Ring, EvCode: evdev.KEY_KP6}
	KeyPad7        = Key{Label: "7", Width: U1, Finger: Index, EvCode: evdev.KEY_KP7}
	KeyPad8        = Key{Label: "8", Width: U1, Finger: Middle, EvCode: evdev.KEY_KP8}
	KeyPad9        = Key{Label: "9", Width: U1, Finger: Ring, EvCode: evdev.KEY_KP9}
	KeyPadPlus     = Key{Label: "", Width: U1, Finger: Pinky, Gap: true, DivLabel: "+", EvCode: evdev.KEY_KPPLUS}
	KeyPadMinus    = Key{Label: "-", Width: U1, Finger: Pinky, EvCode: evdev.KEY_KPMINUS}
	KeyPadAsterisk = Key{Label: "*", Width: U1, Finger: Ring, EvCode: evdev.KEY_KPASTERISK}
	KeyPadSlash    = Key{Label: "/", Width: U1, Finger: Middle, EvCode: evdev.KEY_KPSLASH}
	KeyPadDot      = Key{Label: ".", Width: U1, Finger: Ring, EvCode: evdev.KEY_KPDOT}
	KeyPadEnter    = Key{Label: "", Width: U1, Finger: Pinky, Gap: true, DivLabel: "↵", EvCode: evdev.KEY_KPENTER}

	// MISC
	KeyBlank         = Key{Label: "", Width: U0p75, Finger: Any, EvCode: evdev.KEY_RESERVED}
	KeyPrintScreen   = Key{Label: "Prt", Width: U1, Finger: Index, EvCode: evdev.KEY_SYSRQ}
	KeyLightsToggle  = Key{Label: "☼", Width: U1, Finger: Ring, EvCode: evdev.KEY_LIGHTS_TOGGLE}
	KeyEnterISO      = Key{Label: "Ent", Width: U1p75, Finger: Pinky, EvCode: evdev.KEY_ENTER, Gap: true}
	KeyEnterISOBlank = Key{Label: "", Width: U1p5, Finger: KeyEnterISO.Finger, EvCode: KeyEnterISO.EvCode}
	KeyLeftShiftISO  = Key{Label: "Shft", Width: U1p5, Finger: KeyLeftShift.Finger, EvCode: KeyLeftShift.EvCode}
	KeyRightShiftISO = Key{Label: "Shft", Width: U2, Finger: KeyRightShift.Finger, EvCode: KeyRightShift.EvCode}
	KeyCedilla       = Key{Label: "Ç", Width: U1, Finger: Pinky, EvCode: evdev.KEY_RESERVED}
	KeyTilde         = Key{Label: "~", Width: U1, Finger: Pinky, EvCode: evdev.KEY_RESERVED}
	KeyColon         = Key{Label: ":", Width: U1, Finger: Pinky, EvCode: evdev.KEY_102ND}

	// JIS-SPECIFIC KEYS
	KeyYen      = Key{Label: "¥", Width: U1, Finger: Pinky, EvCode: evdev.KEY_YEN}
	KeyMuhenkan = Key{Label: "無", Width: U1, Finger: Thumb, EvCode: evdev.KEY_MUHENKAN}
	KeyHenkan   = Key{Label: "変", Width: U1, Finger: Thumb, EvCode: evdev.KEY_HENKAN}
	KeyKana     = Key{Label: "仮", Width: U1, Finger: Thumb, EvCode: evdev.KEY_KATAKANAHIRAGANA}

	// KS-SPECIFIC KEYS
	KeyWon     = Key{Label: "₩", Width: U1, Finger: Pinky, EvCode: evdev.KEY_RESERVED}
	KeyHanja   = Key{Label: "한자", Width: U1p5, Finger: Thumb, EvCode: evdev.KEY_HANJA}
	KeyHangeul = Key{Label: "한영", Width: U1p5, Finger: Thumb, EvCode: evdev.KEY_HANGEUL}
)
