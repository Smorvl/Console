package main

import (
	"bufio"
	"fmt"
	"os"
)

type Note struct {
	name    string
	surname string
	note    string
}

func main() {
	var sl []Note

	f := input()
	sl = append(sl, f)

	fl := true
	for fl == true {
		fl = choose(&sl)
	}
}

func input() Note {
	CallClear()

	Scan := bufio.NewScanner(os.Stdin)

	fmt.Println("Name")
	Scan.Scan()
	name := Scan.Text()

	fmt.Println("Surname")
	Scan.Scan()
	surname := Scan.Text()

	fmt.Println("Info")
	Scan.Scan()
	note := Scan.Text()

	f := Note{name, surname, note}

	fmt.Printf("\nEntered data:")
	fmt.Printf("\nName - %s; Surname - %s; Info - %s\n", f.name, f.surname, f.note)

	return f
}

func choose(sl *[]Note) bool {

	fmt.Printf("\nWhat's next?\n")
	fmt.Printf("\ny - proceed, n - finish, p - show all\n")

	var choice string
	_, err := fmt.Scanf("%s\n", &choice)
	if err != nil {
		return false
	}

	if choice == "n" {
		return false
	}

	if choice == "y" {
		v := input()
		*sl = append(*sl, v)
		return true
	}

	if choice == "p" {
		CallClear()
		for i := 0; i < len(*sl); i++ {
			fmt.Printf("Note № %d\n", i+1)
			fmt.Printf("Name - %s\nSurname - %s\nInfo - %s\n\n", (*sl)[i].name, (*sl)[i].surname, (*sl)[i].note)
		}
		return true
	}

	CallClear()
	fmt.Printf("Unidentified command, please try again\n")
	return true
}
func CallClear() {
	fmt.Print("\033[H\033[2J")
}
