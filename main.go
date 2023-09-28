package main

import (
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fmt"
)

func main() {
  a := app.New()
  w := a.NewWindow("Hello")
  hello := widget.NewLabel("Hello Fyne!")
  w.SetContent(container.NewVBox(
    hello,
    widget.NewButton("Hi!", func() {
      hello.SetText("Welcome :)")
    }),
  ))
  w.ShowAndRun()
  tidyUp()

}


func tidyUp() {
  fmt.Println("Exited")
}
