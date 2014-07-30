//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package glut

// #cgo LDFLAGS: -lGL -lGLU -lglut
// #include <stdlib.h>
// #include <GL/glut.h>
// #include "gofunctions.h"
import "C"
import (
	"unsafe"
	"os"
)


const (
	RGB = C.GLUT_RGB
	RGBA = C.GLUT_RGBA
	INDEX = C.GLUT_INDEX
	SINGLE = C.GLUT_SINGLE
	DOUBLE = C.GLUT_DOUBLE
	ACCUM = C.GLUT_ACCUM
	ALPHA = C.GLUT_ALPHA
	DEPTH = C.GLUT_DEPTH
	STENCIL = C.GLUT_STENCIL
	MULTISAMPLE = C.GLUT_MULTISAMPLE
	STEREO = C.GLUT_STEREO
	LUMINANCE = C.GLUT_LUMINANCE
)

type tWindowCallbacks struct {
	display func()
	reshape func(width, height int)
	overlayDisplay func()
	keyboard func(key byte, x, y int)
	mouse func(button, state, x, y int)
	motion func(x, y int)
	passiveMotion func(x, y int)
	visibility func(state int)
	entry func(state int)
	special func(key, x, y int)
	spaceballMotion func(x, y, z int)
	spaceballRotate func(x, y, z int)
	spaceballButton func(button, state int)
	buttonBox func(button, state int)
	dials func(dial, value int)
	tabletMotion func(x, y int)
	tabletButton func(button, state, x, y int)
	menuStatus func(status, x, y int)
	idle func()
	windowStatus func(state int)
	keyboardUp func(key byte, x, y int)
	specialUp func(key, x, y int)
	joystick func(buttonMask uint, x, y, z int)
}


var (
	windowCallbacks = make(map[int]*tWindowCallbacks)
)

func Init() {
	argc := C.int(len(os.Args))
	argv := make([]*C.char, argc)
	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
	}
	defer func() {
		for _, arg := range argv {
			C.free(unsafe.Pointer(arg))
		}
	}()

	C.glutInit(&argc, &argv[0])
}

func InitWindowSize(width, height int) {
	C.glutInitWindowSize(C.int(width), C.int(height))
}

func InitWindowPosition(width, height int) {
	C.glutInitWindowPosition(C.int(width), C.int(height))
}

func InitDisplayMode(mode uint) {
	C.glutInitDisplayMode(C.uint(mode))
}

func MainLoop() {
	C.glutMainLoop()
}

func CreateWindow(title string) (windowId int) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	windowId = int(C.glutCreateWindow(ctitle))
	windowCallbacks[windowId] = new(tWindowCallbacks)

	return windowId
}

func DestroyWindow(windowId int) {
	C.glutDestroyWindow(C.int(windowId))
	delete(windowCallbacks, windowId)
}

func DisplayFunc(display func()) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].display = display
	if display != nil {
		C.register_display()
	} else {
		C.unregister_display()
	}
}

func ReshapeFunc(reshape func(width, height int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].reshape = reshape
	if reshape != nil {
		C.register_reshape()
	} else {
		C.unregister_reshape()
	}
}


//export goReshape
func goReshape(width, height C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].reshape(int(width), int(height))
}

//export goDisplay
func goDisplay() {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].display()
}






