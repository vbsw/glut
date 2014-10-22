
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#ifndef GOFUNCTIONS_H
#define GOFUNCTIONS_H


void register_display();
void register_overlayDisplay();
int create_menu();
int create_menuNullCallback();
void register_reshape();
void register_keyboard();
void register_mouse();
void register_motion();
void register_passiveMotion();
void register_visibility();
void register_entry();
void register_special();
void register_spaceballMotion();
void register_spaceballRotate();
void register_spaceballButton();
void register_buttonBox();
void register_dials();
void register_tabletMotion();
void register_tabletButton();
void register_menuStatus();
void register_menuState();
void register_idle();
void register_timer(unsigned int msecs, int timerId);

void unregister_display();
void unregister_overlayDisplay();
void unregister_reshape();
void unregister_keyboard();
void unregister_mouse();
void unregister_motion();
void unregister_passiveMotion();
void unregister_visibility();
void unregister_entry();
void unregister_special();
void unregister_spaceballMotion();
void unregister_spaceballRotate();
void unregister_spaceballButton();
void unregister_buttonBox();
void unregister_dials();
void unregister_tabletMotion();
void unregister_tabletButton();
void unregister_menuStatus();
void unregister_menuState();
void unregister_idle();

void *const_GLUT_STROKE_ROMAN();
void *const_GLUT_STROKE_MONO_ROMAN();
void *const_GLUT_BITMAP_9_BY_15();
void *const_GLUT_BITMAP_8_BY_13();
void *const_GLUT_BITMAP_TIMES_ROMAN_10();
void *const_GLUT_BITMAP_TIMES_ROMAN_24();
void *const_GLUT_BITMAP_HELVETICA_10();
void *const_GLUT_BITMAP_HELVETICA_12();
void *const_GLUT_BITMAP_HELVETICA_18();

#endif /* GOFUNCTIONS_H */






