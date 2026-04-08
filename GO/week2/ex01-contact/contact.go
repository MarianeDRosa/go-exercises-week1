package main

import "fmt"

// Contact define o modelo de dados (como se fosse uma ficha)
type Contact struct {
	Name  string
	Email string
	Phone string
}

// CreateContact cria um novo contato
func CreateContact(name, email, phone string) Contact {
	return Contact{
		Name:  name,
		Email: email,
		Phone: phone,
	}
}

// UpdateEmail altera o email de um contato existente
// Usamos o * para o Go alterar o contato original, não uma cópia!
func (c *Contact) UpdateEmail(newEmail string) {
	c.Email = newEmail
}

// Display devolve o contato formatado como texto
func (c Contact) Display() string {
	return fmt.Sprintf("Nome: %s | Email: %s | Tel: %s", c.Name, c.Email, c.Phone)
}

func main() {
	// Exemplo de uso
	meuContato := CreateContact("Mariane", "mari@email.com", "51-99999-9999")
	fmt.Println("Antes:", meuContato.Display())

	meuContato.UpdateEmail("mari.nova@email.com")
	fmt.Println("Depois:", meuContato.Display())
}
