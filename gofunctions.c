//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#include <stdio.h>

// Go functions can not be passed to c directly.
// They can only be called from c.
// This code is an indirection to call Go callbacks.
// _cgo_export.h is generated automatically by cgo.
#include "_cgo_export.h"

// Exported functions from Go are:
// goDisplay
// goReshape
// goKeyboard
// goMouse
// goMotion
// goPassiveMotion


#if defined(__APPLE__)
	#include <GLUT/glut.h>
#else
	#include <GL/glut.h>
#endif


// register callbacks

void register_display() {
	glutDisplayFunc(&goDisplay);
}
void register_reshape() {
	glutReshapeFunc(&goReshape);
}
void register_keyboard() {
	glutKeyboardFunc(&goKeyboard);
}
void register_mouse() {
	glutMouseFunc(&goMouse);
}
void register_motion() {
	glutMotionFunc(&goMotion);
}
void register_passiveMotion() {
	glutPassiveMotionFunc(&goPassiveMotion);
}
void register_timer(unsigned int msecs, int timerId) {
	glutTimerFunc(msecs, &goTimer, timerId);
}
void register_idle() {
	glutIdleFunc(&goIdle);
}


// unregister callbacks

void unregister_display() {
	glutDisplayFunc(NULL);
}
void unregister_reshape() {
	glutReshapeFunc(NULL);
}
void unregister_keyboard() {
	glutKeyboardFunc(NULL);
}
void unregister_mouse() {
	glutMouseFunc(NULL);
}
void unregister_motion() {
	glutMotionFunc(NULL);
}
void unregister_passiveMotion() {
	glutPassiveMotionFunc(NULL);
}
void unregister_idle() {
	glutIdleFunc(NULL);
}


