package chardet

var file0 = [...]byte{
	193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229,
	193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229,
	193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229,
	193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229, 193, 99, 197, 233, 164, 164, 164, 229,
	193, 99, 197, 233, 164, 164, 164, 229,
}

var file1 = [...]byte{
	188, 242, 204, 229, 214, 208, 206, 196, 188, 242, 204, 229, 214, 208, 206, 196, 188, 242, 204, 229, 214, 208, 206, 196, 188, 242, 204, 229, 214, 208, 206, 196,
	188, 242, 204, 229, 214, 208, 206, 196, 188, 242, 204, 229, 214, 208, 206, 196, 188, 242, 204, 229, 214, 208, 206, 196, 188, 242, 204, 229, 214, 208, 206, 196,
	188, 242, 204, 229, 214, 208, 206, 196, 188, 242, 204, 229, 214, 208, 206, 196, 188, 242, 204, 229, 214, 208, 206, 196,
}

var file2 = [...]byte{
	147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250,
	150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123,
	140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234,
	147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 147, 250, 150, 123, 140, 234, 10,
}

var file3 = [...]byte{
	230, 177, 137, 229, 173, 151, 230, 188, 162, 229, 173, 151, 231, 181, 177, 228, 184, 128, 231, 183, 168, 231, 162, 188, 232, 144, 172, 229, 156, 139, 231, 162,
	188, 10,
}

var embeddedfiles = map[string][]byte{
	"big5.txt":      file0[0:],
	"gb18030.txt":   file1[0:],
	"shift_jis.txt": file2[0:],
	"utf8.txt":      file3[0:],
}