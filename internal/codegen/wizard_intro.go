package codegen

import (
	"context"
	"fmt"

	"github.com/go-bamboo/pkg/log"
)

// makeWizard creates and returns a new puppeth wizard.
func makeWizard() (*wizard, error) {
	d, err := New()
	if err != nil {
		return nil, err
	}
	uc := NewBiz(d)
	return &wizard{
		in: Stdin,
		uc: uc,
	}, nil
}

// run displays some useful infos to the user, starting on the journey of
// setting up a new or managing an existing Ethereum private network.
func (w *wizard) run(ctx context.Context) error {
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println("| Welcome to codegen                                        |")
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println()

	// Basics done, loop ad infinitum about what to do
	for {
		prompt := `What would you like to do? (default = exit)
	1. page
	2. menu
	3. exit
`
		fmt.Println("")
		fmt.Print(prompt)
		fmt.Println("")

		choice := w.read("choice:")
		switch {
		case choice == "1":
			if err := w.deployPage(ctx); err != nil {
				fmt.Printf("err : %v", err)
			}
		case choice == "2":
			if err := w.deployMenu(ctx); err != nil {
				fmt.Printf("err : %v", err)
			}
		case choice == "3":
			return nil
		default:
			log.Warn("That's not something I can do")
		}
	}
}
