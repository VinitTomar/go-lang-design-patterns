package structural_patterns

import "fmt"

type component interface {
	search(string)
}

type folder struct {
	name string
	components []component
}

func (f *folder) search(keyword string) {
	fmt.Printf("Searching for keyword %v recursively in folder %v\n", keyword, f.name)

	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(composite component) {
	f.components = append(f.components, composite)
}

type file struct {
	name string
}

func (f file) search(keyword string) {
	fmt.Printf("Search for keyword %v in file %v\n", keyword, f.getName())
}

func (f file) getName() string {
	return f.name
}

func CompositePattern() {
	file1 := file{"file1"}
	file2 := file{"file2"}
	file3 := file{"file3"}

	folder1 := &folder{name: "folder1"}
	folder1.add(file1)

	folder2 := &folder{name: "folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("tulip")

}