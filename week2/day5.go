package week2

type Contact struct {
	Name  string
	Phone string
}

type PhoneBook struct {
	Contacts map[string]Contact
}

func (pb *PhoneBook) FindContact(name string) (Contact, bool) {

	contact, ok := pb.Contacts[name]

	return contact, ok
}
