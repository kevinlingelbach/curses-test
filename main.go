package main

import "github.com/kevinlingelbach/curses-test/curses"

func main() {
	curses.Initscr()

	curses.Move(20, 20)

	curses.StartColor()
	curses.InitPair(1, curses.ColorGreen, curses.ColorYellow)

	curses.Attron(curses.ColorPair(1))
	curses.Printw("Hello, world!")
	curses.Attroff(curses.ColorPair(1))

	curses.Refresh()
	curses.Getch()

	curses.Endwin()
}
