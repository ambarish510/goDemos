package structDef1


type Person struct {
	Name  string
	Age   int
	Voter bool
}
// Any struct type that starts with a capital letter is exported and accessible from outside packages.
// Similarly, any struct field that starts with a capital letter is exported.

func New(name string) *Person {

	p := Person{Name: name}
	p.Age = 42
	return &p
}

// Methods can be defined for either pointer or value receiver types
// Go automatically handles conversion between values and pointers for method calls.
// You may want to use a pointer receiver type to avoid copying on method calls
// or to allow the method to mutate the receiving struct.
func (p *Person) IsMajor () bool {
	if p.Age > 18 {
		p.Voter = true
		return true
	}
	return false
}

func (p Person) CallByValue() {
	p.Name ="Name changed" //this will not affect the calling object
}
