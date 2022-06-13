package main

type File struct {
}

func (f *File) Read(buf []byte) (n int, err error) {
	return 1, nil
}
func (f *File) Write(buf []byte) (n int, err error) {
	return 1, nil
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return 1, nil
}
func (f *File) Close() error {
	return nil
}
