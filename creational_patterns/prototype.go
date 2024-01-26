package creational_patterns

import "fmt"

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

type folder struct {
	name string
	children []inode
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, c := range f.children {
		c.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	clonedFolder := &folder{}
	clonedFolder.name = f.name + "_clone"

	var tempChildren []inode

	for _, c := range f.children {
		clone := c.clone()
		tempChildren = append(tempChildren, clone)
	}

	clonedFolder.children = tempChildren

	return clonedFolder
}

func PrototypePattern() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := &folder{
		name: "folder1",
		children: []inode{file1},
	}

	folder2 := &folder{
		name: "folder2",
		children: []inode{folder1, file2, file3},
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}