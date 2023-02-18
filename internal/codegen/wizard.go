package codegen

import (
	"fmt"
	"math/big"
	"net"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/go-bamboo/pkg/log"
)

type wizard struct {
	in   UserPrompter // Wrapper around stdin to allow reading user input
	lock sync.Mutex   // Lock to protect configs during concurrent service discovery
	uc   *biz
}

// read reads a single line from stdin, trimming if from spaces.
func (w *wizard) read(prompt string) string {
	text, err := w.in.PromptInput(prompt)
	if err != nil {
		log.Fatalf("Failed to read user input, %v", err)
	}
	return strings.TrimSpace(text)
}

// readString reads a single line from stdin, trimming if from spaces, enforcing
// non-emptyness.
func (w *wizard) readString(prompt string) string {
	for {
		fmt.Printf("> ")
		text, err := w.in.PromptInput(prompt)
		if err != nil {
			log.Fatalf("Failed to read user input, %v", err)
		}
		if text = strings.TrimSpace(text); text != "" {
			return text
		}
	}
}

// readDefaultString reads a single line from stdin, trimming if from spaces. If
// an empty line is entered, the default value is returned.
func (w *wizard) readDefaultString(prompt string, def string) string {
	text, err := w.in.PromptInput(prompt)
	if err != nil {
		log.Fatalf("Failed to read user input, %v", err)
	}
	if text = strings.TrimSpace(text); text != "" {
		return text
	}
	return def
}

// readDefaultYesNo reads a single line from stdin, trimming if from spaces and
// interpreting it as a 'yes' or a 'no'. If an empty line is entered, the default
// value is returned.
func (w *wizard) readDefaultYesNo(prompt string, def bool) bool {
	for {
		fmt.Printf("> ")
		text, err := w.in.PromptInput(prompt)
		if err != nil {
			log.Fatalf("Failed to read user input, %v", err)
		}
		if text = strings.ToLower(strings.TrimSpace(text)); text == "" {
			return def
		}
		if text == "y" || text == "yes" {
			return true
		}
		if text == "n" || text == "no" {
			return false
		}
		log.Errorf("Invalid input, expected 'y', 'yes', 'n', 'no' or empty")
	}
}

// readURL reads a single line from stdin, trimming if from spaces and trying to
// interpret it as a URL (http, https or file).
func (w *wizard) readURL(prompt string) *url.URL {
	for {
		fmt.Printf("> ")
		text, err := w.in.PromptInput(prompt)
		if err != nil {
			log.Fatalf("Failed to read user input, %v", err)
		}
		uri, err := url.Parse(strings.TrimSpace(text))
		if err != nil {
			log.Errorf("Invalid input, expected URL, err: %v", err)
			continue
		}
		return uri
	}
}

// readInt reads a single line from stdin, trimming if from spaces, enforcing it
// to parse into an integer.
func (w *wizard) readInt(prompt string) int {
	for {
		fmt.Printf("> ")
		text, err := w.in.PromptInput(prompt)
		if err != nil {
			log.Fatalf("Failed to read user input, %v", err)
		}
		if text = strings.TrimSpace(text); text == "" {
			continue
		}
		val, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			log.Errorf("Invalid input, expected integer, err: %v \n", err)
			continue
		}
		return val
	}
}

// readDefaultInt reads a single line from stdin, trimming if from spaces, enforcing
// it to parse into an integer. If an empty line is entered, the default value is
// returned.
func (w *wizard) readDefaultInt(prompt string, def int) int {
	for {
		fmt.Printf("> ")
		text, err := w.in.PromptInput(prompt)
		if err != nil {
			log.Fatalf("Failed to read user input, %v", err)
		}
		if text = strings.TrimSpace(text); text == "" {
			return def
		}
		val, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			log.Errorf("Invalid input, expected integer, err: %v\n", err)
			continue
		}
		return val
	}
}

// readDefaultBigInt reads a single line from stdin, trimming if from spaces,
// enforcing it to parse into a big integer. If an empty line is entered, the
// default value is returned.
func (w *wizard) readDefaultBigInt(prompt string, def *big.Int) *big.Int {
	for {
		// fmt.Printf("> ")
		text, err := w.in.PromptInput(prompt)
		if err != nil {
			log.Fatalf("Failed to read user input, %v", err)
		}
		if text = strings.TrimSpace(text); text == "" {
			return def
		}
		val, ok := new(big.Int).SetString(text, 0)
		if !ok {
			log.Errorf("Invalid input, expected big integer")
			continue
		}
		return val
	}
}

/*
// readFloat reads a single line from stdin, trimming if from spaces, enforcing it
// to parse into a float.
func (w *wizard) readFloat() float64 {
	for {
		fmt.Printf("> ")
		text, err := w.in.ReadString('\n')
		if err != nil {
			log.Crit("Failed to read user input", "err", err)
		}
		if text = strings.TrimSpace(text); text == "" {
			continue
		}
		val, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
		if err != nil {
			log.Error("Invalid input, expected float", "err", err)
			continue
		}
		return val
	}
}
*/

// readDefaultFloat reads a single line from stdin, trimming if from spaces, enforcing
// it to parse into a float. If an empty line is entered, the default value is returned.
func (w *wizard) readDefaultFloat(prompt string, def float64) float64 {
	for {
		// fmt.Printf("> ")
		text, err := w.in.PromptInput(prompt)
		if err != nil {
			log.Fatalf("Failed to read user input, %v", err)
		}
		if text = strings.TrimSpace(text); text == "" {
			return def
		}
		val, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
		if err != nil {
			log.Errorf("Invalid input, expected float", "err", err)
			continue
		}
		return val
	}
}

// readPassword reads a single line from stdin, trimming it from the trailing new
// line and returns it. The input will not be echoed.
// func (w *wizard) readPassword() string {
// 	fmt.Printf("> ")
// 	text, err := terminal.ReadPassword(int(os.Stdin.Fd()))
// 	if err != nil {
// 		log.Crit("Failed to read password", "err", err)
// 	}
// 	fmt.Println()
// 	return string(text)
// }

// readAddress reads a single line from stdin, trimming if from spaces and converts
// it to an Ethereum address.
// func (w *wizard) readAddress() *common.Address {
// 	for {
// 		// Read the address from the user
// 		fmt.Printf("> 0x")
// 		text, err := w.in.ReadString('\n')
// 		if err != nil {
// 			log.Crit("Failed to read user input", "err", err)
// 		}
// 		if text = strings.TrimSpace(text); text == "" {
// 			return nil
// 		}
// 		// Make sure it looks ok and return it if so
// 		if len(text) != 40 {
// 			log.Error("Invalid address length, please retry")
// 			continue
// 		}
// 		bigaddr, _ := new(big.Int).SetString(text, 16)
// 		address := common.BigToAddress(bigaddr)
// 		return &address
// 	}
// }

// readDefaultAddress reads a single line from stdin, trimming if from spaces and
// converts it to an Ethereum address. If an empty line is entered, the default
// value is returned.
// func (w *wizard) readDefaultAddress(def common.Address) common.Address {
// 	for {
// 		// Read the address from the user
// 		fmt.Printf("> 0x")
// 		text, err := w.in.ReadString('\n')
// 		if err != nil {
// 			log.Crit("Failed to read user input", "err", err)
// 		}
// 		if text = strings.TrimSpace(text); text == "" {
// 			return def
// 		}
// 		// Make sure it looks ok and return it if so
// 		if len(text) != 40 {
// 			log.Error("Invalid address length, please retry")
// 			continue
// 		}
// 		bigaddr, _ := new(big.Int).SetString(text, 16)
// 		return common.BigToAddress(bigaddr)
// 	}
// }

// readJSON reads a raw JSON message and returns it.
// func (w *wizard) readJSON(prompt string) string {
// 	var blob json.RawMessage

// 	for {
// 		// fmt.Printf("> ")
// 		text, err := w.in.PromptInput(prompt)
// 		if err != nil {
// 			log.Crit("Failed to read user input", "err", err)
// 		}
// 		if err := json.(&blob); err != nil {
// 			log.Error("Invalid JSON, please try again", "err", err)
// 			continue
// 		}
// 		return string(blob)
// 	}
// }

// readIPAddress reads a single line from stdin, trimming if from spaces and
// returning it if it's convertible to an IP address. The reason for keeping
// the user input format instead of returning a Go net.IP is to match with
// weird formats used by ethstats, which compares IPs textually, not by value.
func (w *wizard) readIPAddress(prompt string) string {
	for {
		// Read the IP address from the user
		// fmt.Printf("> ")
		text, err := w.in.PromptInput(prompt)
		if err != nil {
			log.Fatalf("Failed to read user input, %v", err)
		}
		if text = strings.TrimSpace(text); text == "" {
			return ""
		}
		// Make sure it looks ok and return it if so
		if ip := net.ParseIP(text); ip == nil {
			log.Infof("Invalid IP address, please retry")
			continue
		}
		return text
	}
}
