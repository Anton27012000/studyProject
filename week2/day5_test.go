package week2

import "testing"

func TestFindContactNormal(t *testing.T) {
	book := PhoneBook{
		Contacts: map[string]Contact{
			"Anton": {
				Name:  "Anton",
				Phone: "89857801695",
			},
		},
	}

	contact, ok := book.FindContact("Anton")

	if !ok {
		t.Error("expected contact to exist")
	}

	expected := Contact{Name: "Anton", Phone: "89857801695"}

	if contact != expected {
		t.Errorf("expected %v, got %v", expected, contact)
	}
}

func TestFindContactEmpty(t *testing.T) {

	book := PhoneBook{
		Contacts: map[string]Contact{},
	}

	_, ok := book.FindContact("Ivan")

	if ok {
		t.Error("expected contact to not exist")
	}
}

func TestFindContactEdgeCase(t *testing.T) {

	book := PhoneBook{
		Contacts: map[string]Contact{
			"": {
				Name:  "",
				Phone: "999999",
			},
		},
	}

	contact, ok := book.FindContact("")

	if !ok {
		t.Error("expected empty-name contact")
	}

	expected := "999999"

	if contact.Phone != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			contact.Phone,
		)
	}
}
