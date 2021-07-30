package main

type IOR interface {
	Reader() string
}

type IOW interface {
	Write(val string)
}

type File struct {
	Name string
	Size float64
	Val  string
}

func (t *File) Reader() string {
	return t.Val
}

func (t *File) Write(val string) {
	t.Val = val
}

func main() {
	file := File{"file", 12800, "this is file"}
	var r IOR
	r = &file
	println(&file)
	println(r.Reader())
	var w IOW
	w = r.(IOW)
	w.Write("this is new file")
	println(file.Reader())
}
