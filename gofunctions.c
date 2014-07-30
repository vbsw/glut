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
// goReshape
// goDisplay

#if defined(__APPLE__)
	#include <GLUT/glut.h>
#else
	#include <GL/glut.h>
#endif


// register callbacks

void register_reshape() {
	glutReshapeFunc(&goReshape);
}
void register_display() {
	glutDisplayFunc(&goDisplay);
}


// unregister callbacks

void unregister_reshape() {
	glutReshapeFunc(NULL);
}
void unregister_display() {
	glutDisplayFunc(NULL);
}


