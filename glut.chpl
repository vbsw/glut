/*
 * This is free and unencumbered software released into the public domain.
 * 
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 * 
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 * 
 * For more information, please refer to <http://unlicense.org>
 */


// display modes for InitDisplayMode
extern const GLUT_RGB: int(32);
extern const GLUT_RGBA: int(32);
extern const GLUT_INDEX: int(32);
extern const GLUT_SINGLE: int(32);
extern const GLUT_DOUBLE: int(32);
extern const GLUT_ACCUM: int(32);
extern const GLUT_ALPHA: int(32);
extern const GLUT_DEPTH: int(32);
extern const GLUT_STENCIL: int(32);
extern const GLUT_MULTISAMPLE: int(32);
extern const GLUT_STEREO: int(32);
extern const GLUT_LUMINANCE: int(32);

// mouse states
extern const GLUT_LEFT_BUTTON: int(32);
extern const GLUT_MIDDLE_BUTTON: int(32);
extern const GLUT_RIGHT_BUTTON: int(32);
extern const GLUT_DOWN: int(32);
extern const GLUT_UP: int(32);

// special key codes
extern const GLUT_KEY_F1: int(32);
extern const GLUT_KEY_F2: int(32);
extern const GLUT_KEY_F3: int(32);
extern const GLUT_KEY_F4: int(32);
extern const GLUT_KEY_F5: int(32);
extern const GLUT_KEY_F6: int(32);
extern const GLUT_KEY_F7: int(32);
extern const GLUT_KEY_F8: int(32);
extern const GLUT_KEY_F9: int(32);
extern const GLUT_KEY_F10: int(32);
extern const GLUT_KEY_F11: int(32);
extern const GLUT_KEY_F12: int(32);
extern const GLUT_KEY_LEFT: int(32);
extern const GLUT_KEY_UP: int(32);
extern const GLUT_KEY_RIGHT: int(32);
extern const GLUT_KEY_DOWN: int(32);
extern const GLUT_KEY_PAGE_UP: int(32);
extern const GLUT_KEY_PAGE_DOWN: int(32);
extern const GLUT_KEY_HOME: int(32);
extern const GLUT_KEY_END: int(32);
extern const GLUT_KEY_INSERT: int(32);

// entry/exit callback state
extern const GLUT_LEFT: int(32);
extern const GLUT_ENTERED: int(32);

// window and menu status
extern const GLUT_MENU_NOT_IN_USE: int(32);
extern const GLUT_MENU_IN_USE: int(32);
extern const GLUT_NOT_VISIBLE: int(32);
extern const GLUT_VISIBLE: int(32);
extern const GLUT_HIDDEN: int(32);
extern const GLUT_FULLY_RETAINED: int(32);
extern const GLUT_PARTIALLY_RETAINED: int(32);
extern const GLUT_FULLY_COVERED: int(32);

// RGB color component specification for GetColor
extern const GLUT_RED: int(32);
extern const GLUT_GREEN: int(32);
extern const GLUT_BLUE: int(32);

// the UseLayer parameters
extern const GLUT_NORMAL: int(32);
extern const GLUT_OVERLAY: int(32);

// fonts
param STROKE_ROMAN = 0;
param STROKE_MONO_ROMAN = 1;
param BITMAP_9_BY_15 = 2;
param BITMAP_8_BY_13 = 3;
param BITMAP_TIMES_ROMAN_10 = 4;
param BITMAP_TIMES_ROMAN_24 = 5;
param BITMAP_HELVETICA_10 = 6;
param BITMAP_HELVETICA_12 = 7;
param BITMAP_HELVETICA_18 = 8;

// the Get parameters
extern const GLUT_WINDOW_X: int(32);
extern const GLUT_WINDOW_Y: int(32);
extern const GLUT_WINDOW_WIDTH: int(32);
extern const GLUT_WINDOW_HEIGHT: int(32);
extern const GLUT_WINDOW_BUFFER_SIZE: int(32);
extern const GLUT_WINDOW_STENCIL_SIZE: int(32);
extern const GLUT_WINDOW_DEPTH_SIZE: int(32);
extern const GLUT_WINDOW_RED_SIZE: int(32);
extern const GLUT_WINDOW_GREEN_SIZE: int(32);
extern const GLUT_WINDOW_BLUE_SIZE: int(32);
extern const GLUT_WINDOW_ALPHA_SIZE: int(32);
extern const GLUT_WINDOW_ACCUM_RED_SIZE: int(32);
extern const GLUT_WINDOW_ACCUM_GREEN_SIZE: int(32);
extern const GLUT_WINDOW_ACCUM_BLUE_SIZE: int(32);
extern const GLUT_WINDOW_ACCUM_ALPHA_SIZE: int(32);
extern const GLUT_WINDOW_DOUBLEBUFFER: int(32);
extern const GLUT_WINDOW_RGBA: int(32);
extern const GLUT_WINDOW_PARENT: int(32);
extern const GLUT_WINDOW_NUM_CHILDREN: int(32);
extern const GLUT_WINDOW_COLORMAP_SIZE: int(32);
extern const GLUT_WINDOW_NUM_SAMPLES: int(32);
extern const GLUT_WINDOW_STEREO: int(32);
extern const GLUT_WINDOW_CURSOR: int(32);
extern const GLUT_SCREEN_WIDTH: int(32);
extern const GLUT_SCREEN_HEIGHT: int(32);
extern const GLUT_SCREEN_WIDTH_MM: int(32);
extern const GLUT_SCREEN_HEIGHT_MM: int(32);
extern const GLUT_MENU_NUM_ITEMS: int(32);
extern const GLUT_DISPLAY_MODE_POSSIBLE: int(32);
extern const GLUT_INIT_WINDOW_X: int(32);
extern const GLUT_INIT_WINDOW_Y: int(32);
extern const GLUT_INIT_WINDOW_WIDTH: int(32);
extern const GLUT_INIT_WINDOW_HEIGHT: int(32);
extern const GLUT_INIT_DISPLAY_MODE: int(32);
extern const GLUT_ELAPSED_TIME: int(32);
// glut api version >= 4 or xlib implementation >= 13
extern const GLUT_WINDOW_FORMAT_ID: int(32);

// the DeviceGet parameters
extern const GLUT_HAS_KEYBOARD: int(32);
extern const GLUT_HAS_MOUSE: int(32);
extern const GLUT_HAS_SPACEBALL: int(32);
extern const GLUT_HAS_DIAL_AND_BUTTON_BOX: int(32);
extern const GLUT_HAS_TABLET: int(32);
extern const GLUT_NUM_MOUSE_BUTTONS: int(32);
extern const GLUT_NUM_SPACEBALL_BUTTONS: int(32);
extern const GLUT_NUM_BUTTON_BOX_BUTTONS: int(32);
extern const GLUT_NUM_DIALS: int(32);
extern const GLUT_NUM_TABLET_BUTTONS: int(32);

// glut api version >= 4 or xlib implementation >= 13

extern const GLUT_DEVICE_IGNORE_KEY_REPEAT: int(32);
extern const GLUT_DEVICE_KEY_REPEAT: int(32);
extern const GLUT_HAS_JOYSTICK: int(32);
extern const GLUT_OWNS_JOYSTICK: int(32);
extern const GLUT_JOYSTICK_BUTTONS: int(32);
extern const GLUT_JOYSTICK_AXES: int(32);
extern const GLUT_JOYSTICK_POLL_RATE: int(32);

// the LayerGet parameters
extern const GLUT_OVERLAY_POSSIBLE: int(32);
extern const GLUT_LAYER_IN_USE: int(32);
extern const GLUT_HAS_OVERLAY: int(32);
extern const GLUT_TRANSPARENT_INDEX: int(32);
extern const GLUT_NORMAL_DAMAGED: int(32);
extern const GLUT_OVERLAY_DAMAGED: int(32);


// glutVideoResizeGet parameters
// glut api version >= 4 or xlib implementation >= 9
extern const GLUT_VIDEO_RESIZE_POSSIBLE: int(32);
extern const GLUT_VIDEO_RESIZE_IN_USE: int(32);
extern const GLUT_VIDEO_RESIZE_X_DELTA: int(32);
extern const GLUT_VIDEO_RESIZE_Y_DELTA: int(32);
extern const GLUT_VIDEO_RESIZE_WIDTH_DELTA: int(32);
extern const GLUT_VIDEO_RESIZE_HEIGHT_DELTA: int(32);
extern const GLUT_VIDEO_RESIZE_X: int(32);
extern const GLUT_VIDEO_RESIZE_Y: int(32);
extern const GLUT_VIDEO_RESIZE_WIDTH: int(32);
extern const GLUT_VIDEO_RESIZE_HEIGHT: int(32);

// the GetModifiers parameters
extern const GLUT_ACTIVE_SHIFT: int(32);
extern const GLUT_ACTIVE_CTRL: int(32);
extern const GLUT_ACTIVE_ALT: int(32);

// the SetCursor parameters
extern const GLUT_CURSOR_RIGHT_ARROW: int(32);
extern const GLUT_CURSOR_LEFT_ARROW: int(32);
extern const GLUT_CURSOR_INFO: int(32);
extern const GLUT_CURSOR_DESTROY: int(32);
extern const GLUT_CURSOR_HELP: int(32);
extern const GLUT_CURSOR_CYCLE: int(32);
extern const GLUT_CURSOR_SPRAY: int(32);
extern const GLUT_CURSOR_WAIT: int(32);
extern const GLUT_CURSOR_TEXT: int(32);
extern const GLUT_CURSOR_CROSSHAIR: int(32);
extern const GLUT_CURSOR_UP_DOWN: int(32);
extern const GLUT_CURSOR_LEFT_RIGHT: int(32);
extern const GLUT_CURSOR_TOP_SIDE: int(32);
extern const GLUT_CURSOR_BOTTOM_SIDE: int(32);
extern const GLUT_CURSOR_LEFT_SIDE: int(32);
extern const GLUT_CURSOR_RIGHT_SIDE: int(32);
extern const GLUT_CURSOR_TOP_LEFT_CORNER: int(32);
extern const GLUT_CURSOR_TOP_RIGHT_CORNER: int(32);
extern const GLUT_CURSOR_BOTTOM_RIGHT_CORNER: int(32);
extern const GLUT_CURSOR_BOTTOM_LEFT_CORNER: int(32);
extern const GLUT_CURSOR_INHERIT: int(32);
extern const GLUT_CURSOR_NONE: int(32);
extern const GLUT_CURSOR_FULL_CROSSHAIR: int(32);

// glut api version >= 4 or xlib implementation >= 13
extern const GLUT_KEY_REPEAT_OFF: int(32);
extern const GLUT_KEY_REPEAT_ON: int(32);
extern const GLUT_KEY_REPEAT_DEFAULT: int(32);
extern const GLUT_JOYSTICK_BUTTON_A: int(32);
extern const GLUT_JOYSTICK_BUTTON_B: int(32);
extern const GLUT_JOYSTICK_BUTTON_C: int(32);
extern const GLUT_JOYSTICK_BUTTON_D: int(32);

// glut api version >= 4 or xlib implementation >= 13
extern const GLUT_GAME_MODE_ACTIVE: int(32);
extern const GLUT_GAME_MODE_POSSIBLE: int(32);
extern const GLUT_GAME_MODE_WIDTH: int(32);
extern const GLUT_GAME_MODE_HEIGHT: int(32);
extern const GLUT_GAME_MODE_PIXEL_DEPTH: int(32);
extern const GLUT_GAME_MODE_REFRESH_RATE: int(32);
extern const GLUT_GAME_MODE_DISPLAY_CHANGED: int(32);
