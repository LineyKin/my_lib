package author

type Author struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FatherName string `json:"fatherName,omitempty"`
	LastName   string `json:"lastName"`
}
