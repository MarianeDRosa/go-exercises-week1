// @ts-nocheck
let phonebook = [];

function addContact(contact) {
  phonebook.push(contact);
}

function searchByName(name) {
  return phonebook.find(c => c.name === name) || null;
}

function deleteContact(name) {
  const initialLength = phonebook.length;
  phonebook = phonebook.filter(c => c.name !== name);
  return phonebook.length < initialLength;
}

// Para o teste conseguir limpar a agenda entre um teste e outro
function clearPhonebook() {
  phonebook = [];
}

module.exports = { addContact, searchByName, deleteContact, clearPhonebook };