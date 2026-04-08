package main

import "fmt"

// Reusamos a estrutura de Contact da tarefa anterior
type Contact struct {
	Name  string
	Email string
	Phone string
}

// O Phonebook é apenas uma lista (slice) de contatos
var Phonebook []Contact

// AddContact adiciona um contato à lista
func AddContact(c Contact) {
	Phonebook = append(Phonebook, c)
}

// SearchByName procura um contato pelo nome
func SearchByName(name string) *Contact {
	for i := range Phonebook {
		if Phonebook[i].Name == name {
			return &Phonebook[i]
		}
	}
	return nil
}

// DeleteContact remove um contato pelo nome
func DeleteContact(name string) bool {
	for i, c := range Phonebook {
		if c.Name == name {
			// Técnica de Go para remover item: junta tudo antes do 'i' com tudo depois do 'i'
			Phonebook = append(Phonebook[:i], Phonebook[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	c1 := Contact{Name: "Mariane", Email: "mari@email.com", Phone: "51-9999"}
	AddContact(c1)

	fmt.Println("Contatos na agenda:", len(Phonebook))

	found := SearchByName("Mariane")
	if found != nil {
		fmt.Println("Encontrado:", found.Email)
	}
}
