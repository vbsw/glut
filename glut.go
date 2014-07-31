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


// special key codes
const (
	KEY_F1 = C.GLUT_KEY_F1
	KEY_F2 = C.GLUT_KEY_F2
	KEY_F3 = C.GLUT_KEY_F3
	KEY_F4 = C.GLUT_KEY_F4
	KEY_F5 = C.GLUT_KEY_F5
	KEY_F6 = C.GLUT_KEY_F6
	KEY_F7 = C.GLUT_KEY_F7
	KEY_F8 = C.GLUT_KEY_F8
	KEY_F9 = C.GLUT_KEY_F9
	KEY_F10 = C.GLUT_KEY_F10
	KEY_F11 = C.GLUT_KEY_F11
	KEY_F12 = C.GLUT_KEY_F12
	KEY_LEFT = C.GLUT_KEY_LEFT
	KEY_UP = C.GLUT_KEY_UP
	KEY_RIGHT = C.GLUT_KEY_RIGHT
	KEY_DOWN = C.GLUT_KEY_DOWN
	KEY_PAGE_UP = C.GLUT_KEY_PAGE_UP
	KEY_PAGE_DOWN = C.GLUT_KEY_PAGE_DOWN
	KEY_HOME = C.GLUT_KEY_HOME
	KEY_END = C.GLUT_KEY_END
	KEY_INSERT = C.GLUT_KEY_INSERT
)

// mouse states
const (
	LEFT_BUTTON = C.GLUT_LEFT_BUTTON
	MIDDLE_BUTTON = C.GLUT_MIDDLE_BUTTON
	RIGHT_BUTTON = C.GLUT_RIGHT_BUTTON
	DOWN = C.GLUT_DOWN
	UP = C.GLUT_UP
	// left as in leave
	LEFT = C.GLUT_LEFT
	ENTERED = C.GLUT_ENTERED
)

// display modes for InitDisplayMode
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

// the UseLayer parameters
const (
	NORMAL = C.GLUT_NORMAL
	OVERLAY = C.GLUT_OVERLAY
)

// the Get parameters
const (
	WINDOW_X = C.GLUT_WINDOW_X
	WINDOW_Y = C.GLUT_WINDOW_Y
	WINDOW_WIDTH = C.GLUT_WINDOW_WIDTH
	WINDOW_HEIGHT = C.GLUT_WINDOW_HEIGHT
	WINDOW_BUFFER_SIZE = C.GLUT_WINDOW_BUFFER_SIZE
	WINDOW_STENCIL_SIZE = C.GLUT_WINDOW_STENCIL_SIZE
	WINDOW_DEPTH_SIZE = C.GLUT_WINDOW_DEPTH_SIZE
	WINDOW_RED_SIZE = C.GLUT_WINDOW_RED_SIZE
	WINDOW_GREEN_SIZE = C.GLUT_WINDOW_GREEN_SIZE
	WINDOW_BLUE_SIZE = C.GLUT_WINDOW_BLUE_SIZE
	WINDOW_ALPHA_SIZE = C.GLUT_WINDOW_ALPHA_SIZE
	WINDOW_ACCUM_RED_SIZE = C.GLUT_WINDOW_ACCUM_RED_SIZE
	WINDOW_ACCUM_GREEN_SIZE = C.GLUT_WINDOW_ACCUM_GREEN_SIZE
	WINDOW_ACCUM_BLUE_SIZE = C.GLUT_WINDOW_ACCUM_BLUE_SIZE
	WINDOW_ACCUM_ALPHA_SIZE = C.GLUT_WINDOW_ACCUM_ALPHA_SIZE
	WINDOW_DOUBLEBUFFER = C.GLUT_WINDOW_DOUBLEBUFFER
	WINDOW_RGBA = C.GLUT_WINDOW_RGBA
	WINDOW_PARENT = C.GLUT_WINDOW_PARENT
	WINDOW_NUM_CHILDREN = C.GLUT_WINDOW_NUM_CHILDREN
	WINDOW_COLORMAP_SIZE = C.GLUT_WINDOW_COLORMAP_SIZE
	WINDOW_NUM_SAMPLES = C.GLUT_WINDOW_NUM_SAMPLES
	WINDOW_STEREO = C.GLUT_WINDOW_STEREO
	WINDOW_CURSOR = C.GLUT_WINDOW_CURSOR
	SCREEN_WIDTH = C.GLUT_SCREEN_WIDTH
	SCREEN_HEIGHT = C.GLUT_SCREEN_HEIGHT
	SCREEN_WIDTH_MM = C.GLUT_SCREEN_WIDTH_MM
	SCREEN_HEIGHT_MM = C.GLUT_SCREEN_HEIGHT_MM
	MENU_NUM_ITEMS = C.GLUT_MENU_NUM_ITEMS
	DISPLAY_MODE_POSSIBLE = C.GLUT_DISPLAY_MODE_POSSIBLE
	INIT_WINDOW_X = C.GLUT_INIT_WINDOW_X
	INIT_WINDOW_Y = C.GLUT_INIT_WINDOW_Y
	INIT_WINDOW_WIDTH = C.GLUT_INIT_WINDOW_WIDTH
	INIT_WINDOW_HEIGHT = C.GLUT_INIT_WINDOW_HEIGHT
	INIT_DISPLAY_MODE = C.GLUT_INIT_DISPLAY_MODE
	ELAPSED_TIME = C.GLUT_ELAPSED_TIME
	WINDOW_FORMAT_ID = C.GLUT_WINDOW_FORMAT_ID
)

// the LayerGet parameters
const (
	OVERLAY_POSSIBLE = C.GLUT_OVERLAY_POSSIBLE
	LAYER_IN_USE = C.GLUT_LAYER_IN_USE
	HAS_OVERLAY = C.GLUT_HAS_OVERLAY
	TRANSPARENT_INDEX = C.GLUT_TRANSPARENT_INDEX
	NORMAL_DAMAGED = C.GLUT_NORMAL_DAMAGED
	OVERLAY_DAMAGED = C.GLUT_OVERLAY_DAMAGED
)

// the DeviceGet parameters
const (
	HAS_KEYBOARD = C.GLUT_HAS_KEYBOARD
	HAS_MOUSE = C.GLUT_HAS_MOUSE
	HAS_SPACEBALL = C.GLUT_HAS_SPACEBALL
	HAS_DIAL_AND_BUTTON_BOX = C.GLUT_HAS_DIAL_AND_BUTTON_BOX
	HAS_TABLET = C.GLUT_HAS_TABLET
	NUM_MOUSE_BUTTONS = C.GLUT_NUM_MOUSE_BUTTONS
	NUM_SPACEBALL_BUTTONS = C.GLUT_NUM_SPACEBALL_BUTTONS
	NUM_BUTTON_BOX_BUTTONS = C.GLUT_NUM_BUTTON_BOX_BUTTONS
	NUM_DIALS = C.GLUT_NUM_DIALS
	NUM_TABLET_BUTTONS = C.GLUT_NUM_TABLET_BUTTONS
	DEVICE_IGNORE_KEY_REPEAT = C.GLUT_DEVICE_IGNORE_KEY_REPEAT
	DEVICE_KEY_REPEAT = C.GLUT_DEVICE_KEY_REPEAT
	HAS_JOYSTICK = C.GLUT_HAS_JOYSTICK
	OWNS_JOYSTICK = C.GLUT_OWNS_JOYSTICK
	JOYSTICK_BUTTONS = C.GLUT_JOYSTICK_BUTTONS
	JOYSTICK_AXES = C.GLUT_JOYSTICK_AXES
	JOYSTICK_POLL_RATE = C.GLUT_JOYSTICK_POLL_RATE
)

// the GetModifiers parameters
const (
	ACTIVE_SHIFT = C.GLUT_ACTIVE_SHIFT
	ACTIVE_CTRL = C.GLUT_ACTIVE_CTRL
	ACTIVE_ALT = C.GLUT_ACTIVE_ALT
)

// the SetCursor parameters
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

// RGB color component specification for GetColor
const (
	RED = C.GLUT_RED
	GREEN = C.GLUT_GREEN
	BLUE = C.GLUT_BLUE
)

// additional keyboard and joystick constants
const (
	KEY_REPEAT_OFF = C.GLUT_KEY_REPEAT_OFF
	KEY_REPEAT_ON = C.GLUT_KEY_REPEAT_ON
	KEY_REPEAT_DEFAULT = C.GLUT_KEY_REPEAT_DEFAULT
	JOYSTICK_BUTTON_A = C.GLUT_JOYSTICK_BUTTON_A
	JOYSTICK_BUTTON_B = C.GLUT_JOYSTICK_BUTTON_B
	JOYSTICK_BUTTON_C = C.GLUT_JOYSTICK_BUTTON_C
	JOYSTICK_BUTTON_D = C.GLUT_JOYSTICK_BUTTON_D
)

// game modes
const (
	GAME_MODE_ACTIVE = C.GLUT_GAME_MODE_ACTIVE
	GAME_MODE_POSSIBLE = C.GLUT_GAME_MODE_POSSIBLE
	GAME_MODE_WIDTH = C.GLUT_GAME_MODE_WIDTH
	GAME_MODE_HEIGHT = C.GLUT_GAME_MODE_HEIGHT
	GAME_MODE_PIXEL_DEPTH = C.GLUT_GAME_MODE_PIXEL_DEPTH
	GAME_MODE_REFRESH_RATE = C.GLUT_GAME_MODE_REFRESH_RATE
	GAME_MODE_DISPLAY_CHANGED = C.GLUT_GAME_MODE_DISPLAY_CHANGED
)

type tWindowCallbacks struct {
	display func()
	overlayDisplay func()
	reshape func(width, height int)
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
	windowStatus func(state int)
	keyboardUp func(key uint8, x, y int)
	specialUp func(key, x, y int)
	joystick func(buttonMask uint, x, y, z int)
	idle func()
}


var (
	windowCallbacks = make(map[int]*tWindowCallbacks)
	timers = make(map[int]func(timerId int))
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

func EstablishOverlay() {
	C.glutEstablishOverlay()
}

// Use glut.NORMAL or glut.OVERLAY as layer.
func UseLayer(layer int) {
	C.glutUseLayer(C.GLenum(layer))
}

func RemoveOverlay() {
	C.glutRemoveOverlay()
}

func PostOverlayRedisplay() {
	C.glutPostOverlayRedisplay()
}

func PostWindowOverlayRedisplay(windowId int) {
	C.glutPostWindowOverlayRedisplay(C.int(windowId))
}

func ShowOverlay() {
	C.glutShowOverlay()
}

func HideOverlay() {
	C.glutHideOverlay()
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

func OverlayDisplayFunc(overlayDisplay func()) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].overlayDisplay = overlayDisplay
	if overlayDisplay != nil {
		C.register_overlayDisplay()
	} else {
		C.unregister_overlayDisplay()
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

func MouseFunc(mouse func(button, state, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].mouse = mouse
	if mouse != nil {
		C.register_mouse()
	} else {
		C.unregister_mouse()
	}
}

func MotionFunc(motion func(x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].motion = motion
	if motion != nil {
		C.register_motion()
	} else {
		C.unregister_motion()
	}
}

func PassiveMotionFunc(passiveMotion func(x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].passiveMotion = passiveMotion
	if passiveMotion != nil {
		C.register_passiveMotion()
	} else {
		C.unregister_passiveMotion()
	}
}

// Do not register a second timer with the same id, before the first run out.
// timer should not be nil.
func TimerFunc(msecs int, timer func(timerId int), timerId int) {
	timers[timerId] = timer
	C.register_timer(C.uint(msecs), C.int(timerId))
}

func IdleFunc(idle func()) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].idle = idle
	if idle != nil {
		C.register_idle()
	} else {
		C.unregister_idle()
	}
}

func SetColor(cell int, red, green, blue float32) {
	C.glutSetColor(C.int(cell), C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

func GetColor(cell, component int) (float32) {
	return float32(C.glutGetColor(ccell, component))
}

func CopyColormap(windowId int) {
	C.glutCopyColormap(C.int(windowId))
}

func Get(state int) int {
	return int(C.glutGet(C.GLenum(state)))
}

func LayerGet(info int) int {
	return int(C.glutLayerGet(C.GLenum(info)))
}

func DeviceGet(info int) int {
	return int(C.glutDeviceGet(C.GLenum(info)))
}

func GetModifiers() int {
	return int(C.glutGetModifiers())
}

func ExtensionSupported(extension string) (supported bool) {
	cextension := C.CString(extension)
	defer C.free(unsafe.Pointer(cextension))
	supported = C.glutExtensionSupported(cextension) != 0
	return supported
}


//export goDisplay
func goDisplay() {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].display()
}

//export goOverlayDisplay
func goOverlayDisplay() {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].overlayDisplay()
}

//export goReshape
func goReshape(width, height C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].reshape(int(width), int(height))
}

//export goKeyboard
func goKeyboard(key C.uchar, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].keyboard(uint8(key), int(x), int(y))
}

//export goMouse
func goMouse(button, state, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].mouse(int(button), int(state), int(x), int(y))
}

//export goMotion
func goMotion(x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].motion(int(x), int(y))
}

//export goPassiveMotion
func goPassiveMotion(x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].passiveMotion(int(x), int(y))
}

//export goTimer
func goTimer(timerId C.int) {
	goTimerId := int(timerId)
	timer := timers[goTimerId]
	delete(timers, goTimerId)
	timer(goTimerId)
}

//export goIdle
func goIdle() {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].idle()
}






