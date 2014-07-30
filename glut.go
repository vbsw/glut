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

const (
	CURSOR_RIGHT_ARROW = C.GLUT_CURSOR_RIGHT_ARROW
	CURSOR_LEFT_ARROW = C.GLUT_CURSOR_LEFT_ARROW
	CURSOR_INFO = C.GLUT_CURSOR_INFO
	CURSOR_DESTROY = C.GLUT_CURSOR_DESTROY
	CURSOR_HELP = C.GLUT_CURSOR_HELP
	CURSOR_CYCLE = C.GLUT_CURSOR_CYCLE
	CURSOR_SPRAY = C.GLUT_CURSOR_SPRAY
	CURSOR_WAIT = C.GLUT_CURSOR_WAIT
	CURSOR_TEXT = C.GLUT_CURSOR_TEXT
	CURSOR_CROSSHAIR = C.GLUT_CURSOR_CROSSHAIR
	CURSOR_UP_DOWN = C.GLUT_CURSOR_UP_DOWN
	CURSOR_LEFT_RIGHT = C.GLUT_CURSOR_LEFT_RIGHT
	CURSOR_TOP_SIDE = C.GLUT_CURSOR_TOP_SIDE
	CURSOR_BOTTOM_SIDE = C.GLUT_CURSOR_BOTTOM_SIDE
	CURSOR_LEFT_SIDE = C.GLUT_CURSOR_LEFT_SIDE
	CURSOR_RIGHT_SIDE = C.GLUT_CURSOR_RIGHT_SIDE
	CURSOR_TOP_LEFT_CORNER = C.GLUT_CURSOR_TOP_LEFT_CORNER
	CURSOR_TOP_RIGHT_CORNER = C.GLUT_CURSOR_TOP_RIGHT_CORNER
	CURSOR_BOTTOM_RIGHT_CORNER = C.GLUT_CURSOR_BOTTOM_RIGHT_CORNER
	CURSOR_BOTTOM_LEFT_CORNER = C.GLUT_CURSOR_BOTTOM_LEFT_CORNER
	CURSOR_FULL_CROSSHAIR = C.GLUT_CURSOR_FULL_CROSSHAIR
	CURSOR_NONE = C.GLUT_CURSOR_NONE
	CURSOR_INHERIT = C.GLUT_CURSOR_INHERIT
)

type tWindowCallbacks struct {
	display func()
	reshape func(width, height int)
	overlayDisplay func()
	keyboard func(key uint8, x, y int)
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
	keyboardUp func(key uint8, x, y int)
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

func CreateSubWindow(windowId, x, y, width, height int) (subWindowId int) {
	subWindowId = int(C.glutCreateSubWindow(C.int(windowId), C.int(x), C.int(y), C.int(width), C.int(height)))
	windowCallbacks[subWindowId] = new(tWindowCallbacks)
	return subWindowId
}

func SetWindow(windowId int) {
	C.glutSetWindow(C.int(windowId))
}

func GetWindow() (windowId int) {
	windowId = int(C.glutGetWindow())
	return windowId
}

func DestroyWindow(windowId int) {
	C.glutDestroyWindow(C.int(windowId))
	delete(windowCallbacks, windowId)
}

func PostRedisplay() {
	C.glutPostRedisplay()
}

func SwapBuffers() {
	C.glutSwapBuffers()
}

func PositionWindow(x, y int) {
	C.glutPositionWindow(C.int(x), C.int(y))
}

func ReshapeWindow(width, height int) {
	C.glutReshapeWindow(C.int(width), C.int(height))
}

func FullScreen() {
	C.glutFullScreen()
}

func PopWindow() {
	C.glutPopWindow()
}

func PushWindow() {
	C.glutPushWindow()
}

func ShowWindow() {
	C.glutShowWindow()
}

func HideWindow() {
	C.glutHideWindow()
}

func IconifyWindow() {
	C.glutIconifyWindow()
}

func SetWindowTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.glutSetWindowTitle(ctitle)
}

func SetIconTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.glutSetIconTitle(ctitle)
}

func SetCursor(cursor int) {
	C.glutSetCursor(C.int(cursor))
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

func KeyboardFunc(keyboard func(key uint8, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].keyboard = keyboard
	if keyboard != nil {
		C.register_keyboard()
	} else {
		C.unregister_keyboard()
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

//export goKeyboard
func goKeyboard(key C.uchar, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].keyboard(uint8(key), int(x), int(y))
}






