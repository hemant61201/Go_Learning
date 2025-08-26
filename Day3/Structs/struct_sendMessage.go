package main

type User1 struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User1 {
	if membershipType == "premium" {
		return User1{Name: name, Membership: Membership{Type: membershipType, MessageCharLimit: 1000}}
	} else {
		return User1{Name: name, Membership: Membership{Type: membershipType, MessageCharLimit: 100}}
	}

}

func (user User1) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= user.MessageCharLimit {
		return message, true
	} else {
		return "", false
	}
}
