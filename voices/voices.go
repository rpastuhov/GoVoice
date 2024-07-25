package voices

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

func ChooseVoice() string {
	reader := bufio.NewReader(os.Stdin)

	clearConsole()

	for {
		languages := getKeys(PiperVoices)
		selectedLanguage, back := selectOption("language", languages, reader)
		if back {
			return ""
		}

		clearConsole()

		for {
			voice := getKeys(PiperVoices[selectedLanguage])
			selectedVoice, back := selectOption("voice", voice, reader)
			if back {
				clearConsole()
				break
			}

			clearConsole()

			for {
				qualities := getKeys(PiperVoices[selectedLanguage][selectedVoice])
				selectedQuality, back := selectOption("quality", qualities, reader)
				if back {
					clearConsole()
					break
				}

				clearConsole()

				selectedURL := PiperVoices[selectedLanguage][selectedVoice][selectedQuality]

				return fmt.Sprintf("https://huggingface.co/rhasspy/piper-voices/resolve/v1.0.0/%s", selectedURL)
			}
		}
	}
}

func selectOption(prompt string, options []string, reader *bufio.Reader) (string, bool) {
	// if len(options) == 1 {
	// 	return options[0], false
	// }
	slices.Sort(options)
	for {
		fmt.Printf("Choose a %s:\n", prompt)
		for i, option := range options {
			fmt.Printf("%d: %s\n", i+1, option)
		}
		fmt.Print("Enter number (press 0 to go back): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "0" {
			return "", true
		}
		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(options) {
			clearConsole()
			fmt.Println("Invalid selection. Please try again.")
			continue
		}
		return options[index-1], false
	}
}

var clear = make(map[string]func())

func init() {
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearConsole() {
	if value, ok := clear[runtime.GOOS]; ok {
		value()
	}
}

func getKeys(m interface{}) []string {
	var keys []string
	switch v := m.(type) {
	case map[string]map[string]map[string]string:
		for k := range v {
			keys = append(keys, k)
		}
	case map[string]map[string]string:
		for k := range v {
			keys = append(keys, k)
		}
	case map[string]string:
		for k := range v {
			keys = append(keys, k)
		}
	}
	return keys
}
