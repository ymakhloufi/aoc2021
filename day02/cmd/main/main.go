package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ymakhloufi/aoc2021/day02/internal/pkg/vessels"
)

type Command struct {
	Keyword string
	Steps   int
}

func main() {
	commands, err := readLines("assets/input")
	if err != nil {
		fmt.Printf("error reading file: %s\n", err.Error())
		return
	}

	regularVessel := vessels.RegularVessel{}
	if err := calculatePosition(&regularVessel, commands); err != nil {
		fmt.Printf("failed to find position: %s\n", err.Error())
		return
	}
	fmt.Printf("The regular vessel reached pos %d x and %d y\n", regularVessel.Pos.X, regularVessel.Pos.Y)

	aimedVessel := vessels.AimedVessel{}
	if err := calculatePosition(&aimedVessel, commands); err != nil {
		fmt.Printf("failed to find position: %s\n", err.Error())
		return
	}
	fmt.Printf("The aimed vessel reached pos %d x and %d y\n", aimedVessel.Pos.X, aimedVessel.Pos.Y)
}

func calculatePosition(vessel vessels.Vessel, commands []Command) error {
	for _, command := range commands {
		switch command.Keyword {
		case "forward":
			if err := vessel.Forward(command.Steps); err != nil {
				return err
			}
		case "down":
			if err := vessel.Down(command.Steps); err != nil {
				return err
			}
		case "up":
			if err := vessel.Up(command.Steps); err != nil {
				return err
			}
		default:
			return fmt.Errorf("failed to parse command: %s", command.Keyword)
		}
	}
	return nil
}

func readLines(path string) ([]Command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var lines []Command
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		command := strings.Split(scanner.Text(), " ")
		if len(command) != 2 {
			return nil, errors.New(fmt.Sprintf(
				"The command #%d could not be parsed; it did not contain of two but %d parts.", i, len(command),
			))
		}
		steps, err := strconv.Atoi(command[1])
		if err != nil {
			return lines, err
		}
		lines = append(lines, Command{
			Keyword: command[0],
			Steps:   steps,
		})
	}
	return lines, scanner.Err()
}
