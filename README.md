# GLUT

## Abstract
This is GLUT (version 3.7) binding for the programming language Chapel.

If you want the binding of freeglut visit <https://github.com/vitalibaumtrok/freeglut>.

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

## Development
Programming language Chapel is used.

There is a branch for the programming language Go. To check it out use

	$ git checkout master

The reference for the GLUT version 3 functions is taken from <https://www.opengl.org/resources/libraries/glut/glut-3.spec.pdf>. The reference for the GLUT version 3.7 functions is taken from source code available on <https://www.opengl.org/resources/libraries/glut/glut_downloads.php>.

## Installation

### Linux
[Download](http://chapel.cray.com/download.html) the programming language Chapel and install it as described here <http://chapel.cray.com/docs/1.12/usingchapel/README.html>. You also need Git and the GLUT development package. Install them for example with

        $ sudo apt-get install git freeglut3-dev

Clone this project to folder glut with

	$ git clone https://github.com/vitalibaumtrok/glut.git glut

or update your local copy of this branch with

	$ git pull origin chapel

Checkout the chapel branch with

	$ git checkout chapel

## Example

Will be added...
