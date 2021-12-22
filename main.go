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

	CurrentZone.Systems = append(CurrentZone.Systems, *infra.NewSystem(0, "Sys0", *infra.NewZoneAddress(&CurrentZone)))
	CurrentZone.Systems = append(CurrentZone.Systems, *infra.NewSystem(1, "Sys1", *infra.NewZoneAddress(&CurrentZone)))
	CurrentZone.Systems = append(CurrentZone.Systems, *infra.NewSystem(2, "Sys2", *infra.NewZoneAddress(&CurrentZone)))
}

func main() {
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

	switch strings.ToLower(strings.TrimSpace(input[0])) {
	case Logout.String():
		return true, nil
	case ListSystems.String():
		return false, CurrentZone.GetSystemsInfo()
	}

	return false, nil
}

type Command int64

const (
	Logout Command = iota
	ListSystems
)

func (c Command) String() string {
	switch c {
	case Logout:
		return "logout"
	case ListSystems:
		return "ls"
	}
	return "unknown"
}
