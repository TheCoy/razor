package app

import "fmt"

//BaseApp defines how app works
type BaseApp interface {
	run() error
	initApp() error
}

//CLI show the entrance for main program
func CLI(args []string, app BaseApp) int {
	if err := app.initApp(); err != nil {
		fmt.Println(err)
		return 2
	}

	if err := app.run(); err != nil {
		fmt.Println(err)
		return 3
	}

	return 0
}
