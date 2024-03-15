package main

import "fmt"

type FileSystemNode interface {
	GetSize() int
}

type File struct {
	size int
}

func (f *File) GetSize() int {
	return f.size
}

type Directory struct {
	children []FileSystemNode
}

func (d *Directory) GetSize() int {
	size := 0
	for _, child := range d.children {
		size += child.GetSize()
	}
	return size
}

func (d *Directory) AddChild(child FileSystemNode) {
	d.children = append(d.children, child)
}

func main() {
	file1 := &File{size: 200}
	file2 := &File{size: 300}
	dir := Directory{}
	dir.AddChild(file1)
	dir.AddChild(file2)

	fmt.Println(dir.GetSize()) // Output: 500
}