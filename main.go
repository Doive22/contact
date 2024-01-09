package main

import (
	"bufio"
	"fmt"
	"log"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
)

func main() {
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
	entries, err := os.ReadDir("Contacts")
	if err != nil {
		log.Fatal(err)
	}

	for i, e := range entries {
		b, err := os.ReadFile(filepath.Join("Contacts", e.Name()))
		if err != nil {
			log.Fatal(err)
		}

		ss := strings.Split(string(b), "\n")
		c := Contact{
			Name:  ss[0],
			Email: ss[1],
		}

		fmt.Println(i, c.Name, c.Email)
	}
}

type Contact struct {
	Name  string
	Email string
}
