package main
import "fmt"

// main function
func main() {
	const statement = "Five Finger Death Punch rocks!"
	fmt.Println(statement)

	// ###########################################################################

	var helloWorld string = "Hello World"
  var eineNull int = 0
	fmt.Println(helloWorld)
	fmt.Println(eineNull)

	// ###########################################################################

	statement2 := "Five Finger Death Punch rocks!"
	fmt.Println(statement2)

	// ###########################################################################

	var dreiLeckereEissorten = []string{"Schoko", "Keks", "Vanille"}
  fmt.Println("Array:", dreiLeckereEissorten)

  // oder

  var dreiLeckereEissorten [3]string
  dreiLeckereEissorten[0] = "Schoko"
  dreiLeckereEissorten[1] = "Keks"
  dreiLeckereEissorten[2] = "Vanille"

  // Ausgabe
  fmt.Println(dreiLeckereEissorten) 		// [Schoko Keks Vanille]
  fmt.Println(dreiLeckereEissorten[1:]) // [Keks Vanille]

	// ###########################################################################

	// allgemeine Formatierung
  var str string = "Tutorial"
  fmt.Printf("Variable ist vom Type \"%T\" und hat den Wert \"%v\"\n", str, str)

  // boolean
  var testBoolean bool = false
  fmt.Printf("Hausaufgaben gemacht? -> %t\n", testBoolean)

  // binary
  var testInteger int8 = 12
  fmt.Printf("12 ist binär: %b (12 + 8)\n", testInteger)

  // floats
  var testFloats float32 = 30.5655
  fmt.Printf("e -> %e\n", testFloats)
  fmt.Printf("f -> %f\n", testFloats)
  fmt.Printf("g -> %g\n", testFloats)

	// ###########################################################################

	var testFloats float32 = 30.5655
  var formatString string = fmt.Sprintf("e -> %e\n", testFloats)
  fmt.Println(formatString)
}
