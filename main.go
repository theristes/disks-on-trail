package main

import (
	"fmt"
	"strconv"
	"time"
)

type Disk struct {
	Name  string
	Color string
	Size  int
}
type Trail struct {
	Name  string
	Disks []Disk
}

type TrailManager struct {
	Trails []Trail
}

func hexTo256Color(hexColor string) int {
	// Convert hex color to RGB
	r, _ := strconv.ParseInt(hexColor[1:3], 16, 32)
	g, _ := strconv.ParseInt(hexColor[3:5], 16, 32)
	b, _ := strconv.ParseInt(hexColor[5:7], 16, 32)

	// Convert RGB to 256 color
	return int(16 + (r/51)*36 + (g/51)*6 + (b / 51))
}

func (tm *TrailManager) Print() {

	fmt.Print("\033[H\033[2J") // clear the screen

	// print fancy and colorful trails
	for _, trail := range tm.Trails {
		// print trail name in bold and underlined
		fmt.Print("\033[1;4m")
		fmt.Printf("#%s  T R A I L :         %v\n", trail.Name, len(trail.Disks))
		fmt.Print("\033[0m") // reset formatting

		// print disks
		for _, disk := range trail.Disks {
			// change the color of cli text based on disk color
			colorCode := hexTo256Color(disk.Color)
			fmt.Print("\033[38;5;" + strconv.Itoa(colorCode) + "m")
			// print disk name
			fmt.Print(disk.Name)

			// print new line
			fmt.Println()
			fmt.Print("\033[0m") // reset formatting
		}
		fmt.Println() // add an extra line between trails
	}
}

func (tm *TrailManager) MoveDisk(from int, to int) {
	// move a last disk from a trail to another only if the disk is smaller than the last disk of the destination trail
	fromTrail := &tm.Trails[from-1]
	toTrail := &tm.Trails[to-1]

	if len(fromTrail.Disks) == 0 {
		fmt.Println("No disk to move")
		return
	}
	if len(toTrail.Disks) == 0 || fromTrail.Disks[len(fromTrail.Disks)-1].Size < toTrail.Disks[len(toTrail.Disks)-1].Size {
		toTrail.Disks = append(toTrail.Disks, fromTrail.Disks[len(fromTrail.Disks)-1])
		fromTrail.Disks = fromTrail.Disks[:len(fromTrail.Disks)-1]
	} else {
		fmt.Println("Cannot move disk")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// create a trail manager with 3 trails and a first trail with 4 disks

	trailManager := &TrailManager{
		Trails: []Trail{
			{
				Name: "1",
				Disks: []Disk{
					{
						Name:  "A  B I G G E R    D I S K",
						Color: "#FFFF00",
						Size:  4,
					},
					{
						Name:  "T H I S   B I G   D I S K",
						Color: "#FFA500",
						Size:  3,
					},
					{
						Name:  "A   M E D I U M   D I S K",
						Color: "#FF0000",
						Size:  2,
					},
					{
						Name:  "A    S M A L L    D I S K",
						Color: "#FFFFFF",
						Size:  1,
					},
				},
			},
			{
				Name:  "2",
				Disks: []Disk{},
			},
			{
				Name:  "3",
				Disks: []Disk{},
			},
		},
	}
	trailManager.Print()

	for {
		var command string
		fmt.Scan(&command)

		if command == "P" || command == "p" {
			trailManager.Print()
		} else if command == "M" || command == "m" {
			// show the instructions
			fmt.Println("Enter the index of the trail you want to move from and press enter")
			var fromTrailIndex int
			var toTrailIndex int
			fmt.Scan(&fromTrailIndex)
			fmt.Println("Enter the index of the trail you want to move to and press enter")
			fmt.Scan(&toTrailIndex)

			if fromTrailIndex < 1 || fromTrailIndex > 3 || toTrailIndex < 1 || toTrailIndex > 3 {
				fmt.Println("Invalid trail index, please enter a number between 1 and 3")
				continue
			}

			trailManager.MoveDisk(fromTrailIndex, toTrailIndex)

			if len(trailManager.Trails[1].Disks) == 4 || len(trailManager.Trails[2].Disks) == 4 {
				fmt.Println("Congratulations! You have won the game!")
				break
			}
			trailManager.Print()
		} else if command == "H" || command == "h" {
			fmt.Println("P: Print the trails")
			fmt.Println("M: Move a disk")
			fmt.Println("H: Print this help")
			fmt.Println("Q: Quit")
		} else if command == "Q" || command == "q" {
			break
		} else {
			fmt.Println("Invalid command")
			fmt.Println("Enter H for help")
		}
	}
}
