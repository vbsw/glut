# GLUT

[![GoDoc](https://godoc.org/github.com/vbsw/glut?status.svg)](https://godoc.org/github.com/vbsw/glut)

## Abstract
This is GLUT (version 3.7) binding (published at <https://github.com/vbsw/glut>)
for the programming language Go.

If you want the binding of freeglut visit <https://github.com/vbsw/freeglut>.

## Copying
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org>

## Example

```golang
package main

import (
	"github.com/vbsw/glut"
)

func main() {
	glut.Init()
	glut.InitDisplayMode(glut.SINGLE | glut.RGBA)
	glut.InitWindowSize(640, 480)
	glut.CreateWindow("Testing GLUT binding for Go");
	glut.ReshapeFunc(reshape)
	glut.DisplayFunc(display)
	glut.KeyboardFunc(keyboard)
	glut.MainLoop()
}

func reshape(width, height int) {
	println("reshape")
}

func display() {
	println("display")
}

func keyboard(key uint8, x, y int) {
	if key==27 { // escape
		glut.DestroyWindow(glut.GetWindow())
	} else {
		if (glut.GetModifiers() & glut.ACTIVE_CTRL) > 0 {
			println("key pressed: ctrl +", key)
		} else {
			println("key pressed:", key)
		}
	}
}
```

## Installation

### Linux
Install the programming language [Go](https://golang.org/doc/install) and
the source code management system [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).
You also need GCC, the C standard library development package and the GLUT development package. Install them for example with

        $ sudo apt-get install git gcc libc6-dev freeglut3-dev

Get this project with

	$ go get github.com/vbsw/glut

or update your local copy with

	$ go get -u github.com/vbsw/glut

or compile it with

	$ go install github.com/vbsw/glut

### Windows
Install the programming language [Go](https://golang.org/doc/install).
Install [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

Get this project with (run this preferably from the Git Bash)

	$ go get github.com/vbsw/glut

or update your local copy with

	$ go get -u github.com/vbsw/glut

Cgo needs another compiler to compile c files. On a 64 bit system a 64 bit compiler is needed. I tried it with the gcc. MinGW provides only 32 bit binaries. So go to <http://tdm-gcc.tdragon.net> to download 64 bit gcc binaries. Install it.

Then download freeglut from <http://www.transmissionzero.co.uk/software/freeglut-devel/> and unpack it for example to

	C:/Users/Alice/Downloads/freeglut

Then open the file github.com/vbsw/glut/glut.go and change the first line from

	// #cgo LDFLAGS: -lGL -lGLU -lglut
	// #include <stdlib.h>
	// #include <GL/glut.h>
	// #include "gofunctions.h"

to

	// #cgo CFLAGS: -IC:/Users/Alice/Downloads/freeglut/include
	// #cgo LDFLAGS: -LC:/Users/Alice/Downloads/freeglut/bin/x64 -l:freeglut.dll
	// #include <stdlib.h>
	// #include <GL/glut.h>
	// #include "gofunctions.h"

To compile run

	$ go install github.com/vbsw/glut

## References

- <https://www.opengl.org/resources/libraries/glut/glut-3.spec.pdf>
- <https://www.opengl.org/resources/libraries/glut/glut_downloads.php>

