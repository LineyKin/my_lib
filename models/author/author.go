package author

type Author struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FatherName string `json:"fatherName"`
	LastName   string `json:"lastName"`
}

func (a Author) Add() (int, error) {
	return 25, nil
}
