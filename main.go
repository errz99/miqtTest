package main

import (
	"fmt"
	"os"

	"github.com/mappu/miqt/qt"
)

func main() {

	qt.NewQApplication(os.Args)

	btn := qt.NewQPushButton3("Hello world!")
	btn.SetFixedWidth(320)

	var counter int = 0

	btn.OnPressed(func() {
		counter++
		if counter > 1 {
			btn.SetText(fmt.Sprintf("Has pulsado el boton %d veces", counter))
		} else {
			btn.SetText(fmt.Sprintf("Has pulsado el boton 1 vez"))
		}
	})

	btn.Show()

	qt.QApplication_Exec()

	fmt.Println("OK!")
}
