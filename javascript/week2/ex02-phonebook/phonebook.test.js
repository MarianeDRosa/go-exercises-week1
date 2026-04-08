// @ts-nocheck
const { addContact, searchByName, deleteContact, clearPhonebook } = require('./phonebook');

describe('Phonebook System', () => {
  beforeEach(() => {
    clearPhonebook();
  });

  test('should add and search a contact', () => {
    addContact({ name: 'Mariane', email: 'mari@test.com' });
    const found = searchByName('Mariane');
    expect(found.email).toBe('mari@test.com');
  });

  test('should delete a contact', () => {
    addContact({ name: 'Davi', email: 'davi@test.com' });
    const result = deleteContact('Davi');
    expect(result).toBe(true);
    expect(searchByName('Davi')).toBe(null);
  });
});