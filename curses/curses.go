package curses

// #cgo LDFLAGS: -lncurses
// #include <curses.h>
// #include <stdlib.h>
//
// int cprintw(const char *fmt) {
// 		printw(fmt);
// }
import "C"
import "unsafe"

type Color C.short

// Color constants
const (
	ColorYellow  Color = C.COLOR_YELLOW
	ColorGreen         = C.COLOR_GREEN
	ColorCyan          = C.COLOR_CYAN
	ColorRed           = C.COLOR_RED
	ColorMagenta       = C.COLOR_MAGENTA
	ColorWhite         = C.COLOR_WHITE
	ColorBlack         = C.COLOR_BLACK
)

// InitPair calls C.init_pair()
func InitPair(index int16, f, b Color) {
	C.init_pair(C.short(index), C.short(f), C.short(b))
}

// Attron calls C.attron
func Attron(attrs int) {
	C.attron(C.int(attrs))
}

// Attroff calls C.attroff
func Attroff(attrs int) {
	C.attroff(C.int(attrs))
}

// ColorPair calls C.COLOR_PAIR
func ColorPair(i int) int {
	return int(C.COLOR_PAIR(C.int(i)))
}

// StartColor calls c.start_color()
func StartColor() {
	C.start_color()
}

// Initscr calls C.initscr()
func Initscr() {
	C.initscr()
}

// Refresh calls C.refresh()
func Refresh() {
	C.refresh()
}

// Printw calls C.printw()
func Printw(s string) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	C.cprintw(cStr)
}

// Getch calls C.getch()
func Getch() {
	C.getch()
}

// Endwin() calls C.endwin()
func Endwin() {
	C.endwin()
}

// Move() calls C.move()
func Move(x int, y int) {
	C.move(C.int(x), C.int(y))
}
