// @ts-nocheck
function createContact(name, email, phone) {
  return {
    name,
    email,
    phone,
    // Método para atualizar o email
    updateEmail(newEmail) {
      this.email = newEmail;
    },
    // Método para exibir formatado
    display() {
      return `Nome: ${this.name} | Email: ${this.email} | Tel: ${this.phone}`;
    }
  };
}

module.exports = { createContact };