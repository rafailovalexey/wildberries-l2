package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCommands(t *testing.T) {
	commands := GetMockCommands()

	arguments := map[string][]string{
		"cd":   {"cd", "."},
		"pwd":  {"pwd"},
		"echo": {"echo"},
		"exit": {"exit"},
	}

	for _, command := range commands {
		t.Run(command.name, func(t *testing.T) {
			output := CaptureOutput(func() {
				exit, err := command.handle(arguments[command.name])

				if err != nil {
					t.Errorf("unexpected error %v", err)
				}

				if exit && command.name != "exit" {
					t.Errorf("expected %s command to not exit", command.name)
				}
			})

			expected := GetExpectedOutput(command, t)

			if output != expected {
				t.Errorf("expected output to be %q got %q", expected, output)
			}
		})
	}
}

func GetExpectedOutput(command Command, t *testing.T) string {
	switch command.name {
	case "pwd":
		cwd, err := os.Getwd()

		if err != nil {
			t.Fatalf("error getting current working directory %v", err)
		}

		return fmt.Sprintf("%s\n", cwd)
	case "echo":
		return fmt.Sprintf("\n")
	case "ps":
		cmd := exec.Command("ps")

		output, err := cmd.Output()

		if err != nil {
			t.Fatalf("error running 'ps' command %v", err)
		}

		return fmt.Sprintf("%s", output)
	case "exit":
		return "выход из терминала\n"
	default:
		return ""
	}
}

func CaptureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	f()

	w.Close()

	out, _ := io.ReadAll(r)

	os.Stdout = old

	return string(out)
}

func GetMockCommands() map[string]Command {
	commands := &Commands{
		Command{
			name: "cd",
			handle: func(arguments []string) (bool, error) {
				if len(arguments) < 2 {
					return false, fmt.Errorf("usage: cd <directory>")
				}

				err := os.Chdir(arguments[1])

				if err != nil {
					return false, fmt.Errorf("error when changing directory %v", err)
				}

				return false, nil
			},
		},
		Command{
			name: "pwd",
			handle: func(arguments []string) (bool, error) {
				directory, err := os.Getwd()

				if err != nil {
					return false, fmt.Errorf("error when getting current directory %v", err)
				}

				fmt.Printf("%s\n", directory)

				return false, nil
			},
		},
		Command{
			name: "echo",
			handle: func(arguments []string) (bool, error) {
				fmt.Printf("%s\n", strings.Join(arguments[1:], " "))

				return false, nil
			},
		},
		Command{
			name: "exit",
			handle: func(arguments []string) (bool, error) {
				fmt.Printf("exit from the terminal\n")

				return true, nil
			},
		},
	}

	dictionary := make(map[string]Command, len(*commands))

	for _, command := range *commands {
		dictionary[command.name] = command
	}

	return dictionary
}
