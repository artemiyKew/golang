package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct{}

type Aircraft struct{}

func drive(v Vehicle) {
	v.move()
}

func (c *Car) move() {
	fmt.Println("Автомобиль едет")
}

func (a *Aircraft) move() {
	fmt.Println("Самолет летит")
}

type Stream interface {
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}

func closeStream(stream Stream) {
	fmt.Println("Close stream")
	stream.close()
}

type File struct {
	text string
}

type Folder struct{}

func (f *File) read() string {
	return f.text
}

func (f *File) write(msg string) {
	f.text = msg
	fmt.Println("Запись в файл строки", msg)
}

func (f *File) close() {
	fmt.Println("File closed")
}

func (f *Folder) close() {
	fmt.Println("Folder closed")
}

func Test() {
	var tesla Vehicle = &Car{}
	var boing Vehicle = &Aircraft{}

	tesla.move()
	boing.move()
	drive(tesla)
	drive(boing)

	myFile := &File{}
	myFolder := Folder{}

	writeToStream(myFile, "Hello world")
	closeStream(myFile)

	// closeStream(&myFolder) Error тип *Folder не реализует интерфейс Stream
	myFolder.close()
}
