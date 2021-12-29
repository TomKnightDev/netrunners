package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	infra "github.com/tomknightdev/netrunners/infrastructure"
)

var CurrentZone = infra.Zone{}

func Init() {
	CurrentZone = *infra.NewZone(0, "Zone0")

	for i := 0; i < 100; i++ {
		CurrentZone.Systems = append(CurrentZone.Systems, *infra.NewSystem(i, fmt.Sprintf("Sys%03d", i), *infra.NewZoneAddress(&CurrentZone)))
	}

	// CurrentZone.Systems = append(CurrentZone.Systems, *infra.NewSystem(0, "Sys0", *infra.NewZoneAddress(&CurrentZone)))
	// CurrentZone.Systems = append(CurrentZone.Systems, *infra.NewSystem(1, "Sys1", *infra.NewZoneAddress(&CurrentZone)))
	// CurrentZone.Systems = append(CurrentZone.Systems, *infra.NewSystem(2, "Sys2", *infra.NewZoneAddress(&CurrentZone)))
}

func main() {
	Init()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("user@machine: ")
		text, _ := reader.ReadString('\n')

		quit, vals := ProcessCommand(text)

		if quit {
			break
		}

		for _, sys := range vals {
			fmt.Println(sys)
		}
	}
}

func ProcessCommand(command string) (bool, []string) {
	input := strings.Split(command, " ")
	for i := range input {
		input[i] = strings.TrimSpace(input[i])
	}

	switch strings.ToLower(input[0]) {
	case Logout.String():
		return true, nil
	case ListSystems.String():
		return false, CurrentZone.GetSystemsInfo()
	case ListSystemHardware.String():
		if len(input) == 0 {
			return false, []string{"System name not provided"}
		}
		return false, []string{CurrentZone.GetSystemInfo(input[1])}
	}

	return false, nil
}

type Command int64

const (
	Logout Command = iota
	ListSystems
	ListSystemHardware
)

func (c Command) String() string {
	switch c {
	case Logout:
		return "logout"
	case ListSystems:
		return "ls"
	case ListSystemHardware:
		return "lshw"
	}
	return "unknown"
}
