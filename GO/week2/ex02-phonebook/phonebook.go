package main

import (
	"bufio"
	"fmt"
	"os"
)

// Contact represents a phonebook entry
type Contact struct {
	Name  string
	Email string
}

// Global variable to store contacts
var Phonebook []Contact

// AddContact adds a new contact to the slice
func AddContact(c Contact) {
	Phonebook = append(Phonebook, c)
}

// SearchByName looks for a contact by name and returns a pointer or an error
func SearchByName(name string) (*Contact, error) {
	for _, c := range Phonebook {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("contact not found")
}

// DeleteContact removes a contact from the slice by name
func DeleteContact(name string) error {
	for i, c := range Phonebook {
		if c.Name == name {
			Phonebook = append(Phonebook[:i], Phonebook[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contact not found")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Phonebook Menu ---")
		fmt.Println("1. Add Contact")
		fmt.Println("2. Search Contact")
		fmt.Println("3. Delete Contact")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Enter name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Enter email: ")
			scanner.Scan()
			email := scanner.Text()

			AddContact(Contact{Name: name, Email: email})
			fmt.Println("Contact added successfully!")

		case "2":
			fmt.Print("Enter name to search: ")
			scanner.Scan()
			name := scanner.Text()

			contact, err := SearchByName(name)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Found: %s (%s)\n", contact.Name, contact.Email)
			}

		case "3":
			fmt.Print("Enter name to delete: ")
			scanner.Scan()
			name := scanner.Text()

			err := DeleteContact(name)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Contact deleted successfully!")
			}

		case "4":
			fmt.Println("Exiting... Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
