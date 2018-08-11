package db

// Qoute type
type Qoute struct {
	Type   string
	Text   string
	Auther string
}

var myqoute = make([]Qoute, 0)

// Save new qoute
func Save(s, a string) error {
	myqoute = append(myqoute, Qoute{"text", s, a})
	return nil
}

//All func get all Qoutes
func All() ([]Qoute, error) {
	return myqoute, nil
}
