package main

import (
	"fmt"
	"github/Er-Sadiq/GoOverloader/services"
	"os"
)

func main() {
	var opt int

	fmt.Println("Less Talking More Doing..")
	fmt.Println()
	fmt.Println("Option 1: Launches a DoS attack.")
	fmt.Println("Option 2: Exits the program.")
	fmt.Println("Reminder: you can always exit by pressing ctrl+c")

	for {

		fmt.Print("Please select an option (1 or 2): ")
		fmt.Scan(&opt)

		switch opt {
		case 1:
			var threads, duration, port int
			var targetIp string

			fmt.Print("Enter The Number of Threads: ")
			fmt.Scan(&threads)

			fmt.Println("Enter Attack Duration in secs:")
			fmt.Scan(&duration)

			fmt.Println("Enter The Target IP or URL:")
			fmt.Scanln(&targetIp)

			fmt.Println("Enter A Valid Port Number:")
			fmt.Scan(&port)

			validatedTarged := services.GetIp(targetIp, port)
			fmt.Println("Target Validated & Locked:", validatedTarged)

			// lunching the attack
			services.Attack(validatedTarged, port, duration, threads)

			break

		case 2:

			fmt.Println("Thank you for using GoOverloader")
			fmt.Println("GoOverloader Signing Off")
			os.Exit(0)

		default:

			fmt.Println("Invalid option! Please choose either 1 or 2.")
		}
	}
}
