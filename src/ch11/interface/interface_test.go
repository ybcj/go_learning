package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgram struct {

}

func (g *GoProgram) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgram)
	t.Log(p.WriteHelloWorld())
}
