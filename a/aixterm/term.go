// Generated automatically.  DO NOT HAND-EDIT.

package aixterm

import "github.com/kckrinke/go-terminfo"

func init() {

	// IBM Aixterm Terminal Emulator
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "aixterm",
		Columns:      80,
		Lines:        25,
		Colors:       8,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[J",
		AttrOff:      "\x1b[0;10m\x1b(B",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Reverse:      "\x1b[7m",
		SetFg:        "\x1b[3%p1%dm",
		SetBg:        "\x1b[4%p1%dm",
		SetFgBg:      "\x1b[3%p1%d;4%p2%dm",
		ResetFgBg:    "\x1b[32m\x1b[40m",
		PadChar:      "\x00",
		AltChars:     "jjkkllmmnnqqttuuvvwwxx",
		EnterAcs:     "\x1b(0",
		ExitAcs:      "\x1b(B",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[139q",
		KeyDelete:    "\x1b[P",
		KeyBackspace: "\b",
		KeyHome:      "\x1b[H",
		KeyEnd:       "\x1b[146q",
		KeyPgUp:      "\x1b[150q",
		KeyPgDn:      "\x1b[154q",
		KeyF1:        "\x1b[001q",
		KeyF2:        "\x1b[002q",
		KeyF3:        "\x1b[003q",
		KeyF4:        "\x1b[004q",
		KeyF5:        "\x1b[005q",
		KeyF6:        "\x1b[006q",
		KeyF7:        "\x1b[007q",
		KeyF8:        "\x1b[008q",
		KeyF9:        "\x1b[009q",
		KeyF10:       "\x1b[010q",
		KeyF11:       "\x1b[011q",
		KeyF12:       "\x1b[012q",
		KeyF13:       "\x1b[013q",
		KeyF14:       "\x1b[014q",
		KeyF15:       "\x1b[015q",
		KeyF16:       "\x1b[016q",
		KeyF17:       "\x1b[017q",
		KeyF18:       "\x1b[018q",
		KeyF19:       "\x1b[019q",
		KeyF20:       "\x1b[020q",
		KeyF21:       "\x1b[021q",
		KeyF22:       "\x1b[022q",
		KeyF23:       "\x1b[023q",
		KeyF24:       "\x1b[024q",
		KeyF25:       "\x1b[025q",
		KeyF26:       "\x1b[026q",
		KeyF27:       "\x1b[027q",
		KeyF28:       "\x1b[028q",
		KeyF29:       "\x1b[029q",
		KeyF30:       "\x1b[030q",
		KeyF31:       "\x1b[031q",
		KeyF32:       "\x1b[032q",
		KeyF33:       "\x1b[033q",
		KeyF34:       "\x1b[034q",
		KeyF35:       "\x1b[035q",
		KeyF36:       "\x1b[036q",
		KeyClear:     "\x1b[144q",
		KeyBacktab:   "\x1b[Z",
	})
}
