
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
// goOverlayDisplay
// goReshape
// goKeyboard
// goMouse
// goMotion
// goPassiveMotion
// goVisibility
// goEntry
// goTimer
// goIdle


#if defined(__APPLE__)
	#include <GLUT/glut.h>
#else
	#include <GL/glut.h>
#endif


// register callbacks

void register_display() {
	glutDisplayFunc(&goDisplay);
}
void register_overlayDisplay() {
	glutOverlayDisplayFunc(&goOverlayDisplay);
}
int create_menu() {
	return glutCreateMenu(&goMenu);
}
int create_menuNullCallback() {
	return glutCreateMenu(NULL);
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
void register_visibility() {
	glutVisibilityFunc(&goVisibility);
}
void register_entry() {
	glutEntryFunc(&goEntry);
}
void register_special() {
	glutSpecialFunc(&goSpecial);
}
void register_spaceballMotion() {
	glutSpaceballMotionFunc(&goSpaceballMotion);
}
void register_spaceballRotate() {
	glutSpaceballRotateFunc(&goSpaceballRotate);
}
void register_spaceballButton() {
	glutSpaceballButtonFunc(&goSpaceballButton);
}
void register_buttonBox() {
	glutButtonBoxFunc(&goButtonBox);
}
void register_dials() {
	glutDialsFunc(&goDials);
}
void register_tabletMotion() {
	glutTabletMotionFunc(&goTabletMotion);
}
void register_tabletButton() {
	glutTabletButtonFunc(&goTabletButton);
}
void register_menuStatus() {
	glutMenuStatusFunc(&goMenuStatus);
}
void register_menuState() {
	glutMenuStateFunc(&goMenuState);
}
void register_idle() {
	glutIdleFunc(&goIdle);
}
void register_timer(unsigned int msecs, int timerId) {
	glutTimerFunc(msecs, &goTimer, timerId);
}


// unregister callbacks

void unregister_display() {
	glutDisplayFunc(NULL);
}
void unregister_overlayDisplay() {
	glutOverlayDisplayFunc(NULL);
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
void unregister_visibility() {
	glutVisibilityFunc(NULL);
}
void unregister_entry() {
	glutEntryFunc(NULL);
}
void unregister_special() {
	glutSpecialFunc(NULL);
}
void unregister_spaceballMotion() {
	glutSpaceballMotionFunc(NULL);
}
void unregister_spaceballRotate() {
	glutSpaceballRotateFunc(NULL);
}
void unregister_spaceballButton() {
	glutSpaceballButtonFunc(NULL);
}
void unregister_buttonBox() {
	glutButtonBoxFunc(NULL);
}
void unregister_dials() {
	glutDialsFunc(NULL);
}
void unregister_tabletMotion() {
	glutTabletMotionFunc(NULL);
}
void unregister_tabletButton() {
	glutTabletButtonFunc(NULL);
}
void unregister_menuStatus() {
	glutMenuStatusFunc(NULL);
}
void unregister_menuState() {
	glutMenuStateFunc(NULL);
}
void unregister_idle() {
	glutIdleFunc(NULL);
}

// Fonts

void *const_GLUT_STROKE_ROMAN() {
	return GLUT_STROKE_ROMAN;
}
void *const_GLUT_STROKE_MONO_ROMAN() {
	return GLUT_STROKE_MONO_ROMAN;
}
void *const_GLUT_BITMAP_9_BY_15() {
	return GLUT_BITMAP_9_BY_15;
}
void *const_GLUT_BITMAP_8_BY_13() {
	return GLUT_BITMAP_8_BY_13;
}
void *const_GLUT_BITMAP_TIMES_ROMAN_10() {
	return GLUT_BITMAP_TIMES_ROMAN_10;
}
void *const_GLUT_BITMAP_TIMES_ROMAN_24() {
	return GLUT_BITMAP_TIMES_ROMAN_24;
}
void *const_GLUT_BITMAP_HELVETICA_10() {
	return GLUT_BITMAP_HELVETICA_10;
}
void *const_GLUT_BITMAP_HELVETICA_12() {
	return GLUT_BITMAP_HELVETICA_12;
}
void *const_GLUT_BITMAP_HELVETICA_18() {
	return GLUT_BITMAP_HELVETICA_18;
}

