# Tastatureingaben einlesen
Wer in Go Werte während der Laufzeit einlesen möchte, der kann das Paket für den buffered Input und Output verwenden. Dazu wird eine Instanz erstellt und anschließend die Scanner() Methode aufgerufen, Das Programm wartet, bis die Enter Taste gedrückt wurde. Die Eingabe kann dann mit der Methode Text() ausgelesen werden. Egal was der Nutzer eingibt, der Scanner liest ausschließlich Strings ein. Umwandlungen zu Zahlen müssen selbst abgewickelt werden.
```go
package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  Scanner := bufio.NewScanner(os.Stdin)
  scnner.Scan()
  input := scanner.Text()
  fmt.Println("Dein Input:",input)
}
```

Hier nochmal ein Beispiel zum einlesen des Alters:
```go
package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  Scanner := bufio.NewScanner(os.Stdin)
  fmt.Println()
  scnner.Scan("Wie alt bist du?")
  input, _ := strconv.ParseUint(scanner.Text(), 10, 8)
  fmt.Printf("Vor einem Jahr wärst du noch %2f Jahre alt gewesen", input - 1)
}
```
