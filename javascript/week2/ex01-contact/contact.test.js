// @ts-nocheck
const { createContact } = require('./contact');

describe('Contact Object', () => {
  test('should create a contact correctly', () => {
    const contact = createContact('Mariane', 'mari@test.com', '123');
    expect(contact.name).toBe('Mariane');
    expect(contact.email).toBe('mari@test.com');
  });

  test('should update email using the method', () => {
    const contact = createContact('Mariane', 'mari@test.com', '123');
    contact.updateEmail('novo@test.com');
    expect(contact.email).toBe('novo@test.com');
  });
});