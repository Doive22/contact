package main

import (
	"bufio"
	"fmt"
	"log"
	"net/mail"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2/app"
)

var myList []Contact

func main() {
	a := app.New()
	w := a.NewWindow("Contact")

	var err error
	myList, err = loadContacts("Contacts")
	if err != nil {
		log.Fatal(err)
	}

	g := newGUI()
	w.SetContent(g.makeUI())
	w.ShowAndRun()
}

type Contact struct {
	Name  string
	Email string
}

func loadContacts(contactDirPath string) ([]Contact, error) {
	entries, err := os.ReadDir(contactDirPath)
	if err != nil {
		log.Fatal(err)
	}

	var contacts []Contact
	for _, e := range entries {
		b, err := os.ReadFile(filepath.Join(contactDirPath, e.Name()))
		if err != nil {
			return contacts, err
		}

		ss := strings.Split(string(b), "\n")
		c := Contact{
			Name:  ss[0],
			Email: ss[1],
		}

		contacts = append(contacts, c)
	}

	return contacts, nil
}

func contactTerm() {
	fmt.Println("Hello Dadz")
	fmt.Println("Who do you wanna add?")

	input := bufio.NewReader(os.Stdin)
	var name string
	fmt.Print("Name: ")

	line, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(line)

	var email string
	fmt.Print("Email: ")
	line, err = input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	email = strings.TrimSpace(line)

	validEmail, err := mail.ParseAddress(email)
	if err != nil {
		log.Fatal(err)
	}

	contact := Contact{
		Name:  name,
		Email: validEmail.Address,
	}

	fmt.Println(contact)

	err = os.WriteFile(filepath.Join("Contacts", contact.Name), []byte(contact.Name+"\n"+contact.Email), 0o600)
	if err != nil {
		log.Fatal(err)
	}

}
