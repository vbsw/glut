//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)

#ifndef GOFUNCTIONS_H
#define GOFUNCTIONS_H

void register_display();
void register_overlayDisplay();
void register_reshape();
void register_keyboard();
void register_mouse();
void register_motion();
void register_passiveMotion();
void register_timer(unsigned int msecs, int timerId);
void register_idle();

void unregister_display();
void unregister_overlayDisplay();
void unregister_reshape();
void unregister_keyboard();
void unregister_mouse();
void unregister_motion();
void unregister_passiveMotion();
void unregister_idle();

#endif /* GOFUNCTIONS_H */


