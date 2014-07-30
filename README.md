# glut

## Abstract
GLUT binding for the programming language Go. It works on linux.

## Installation

	go get github.com/vitalibaumtrok/glut
	go install github.com/vitalibaumtrok/glut

To update the package type

	go get -u github.com/vitalibaumtrok/glut

_Note: You need the programming language Go and git to run these commands._

## Development
Programming language Go is used.

## Coverage
* void glutInit(int *argcp, char **argv)
* void glutInitWindowSize(int width, int height)
* void glutInitWindowPosition(int x, int y)
* void glutInitDisplayMode(unsigned int mode)
* void glutMainLoop(void)
* int glutCreateWindow(char *name)
* void glutDestroyWindow(int win)
* void glutDisplayFunc(void (*func)(void))
* void glutReshapeFunc(void (*func)(int width, int height))

## Example
	package main

	import (
		"github.com/vitalibaumtrok/glut"
	)

	func main() {
		glut.Init()
		glut.InitDisplayMode(glut.SINGLE | glut.RGBA)
		glut.InitWindowSize(640, 480)
		glut.CreateWindow("Testing GLUT binding for Go");
		glut.ReshapeFunc(reshape)
		glut.DisplayFunc(display)
		glut.MainLoop()
	}

	func reshape(widht, height int) {
		println("reshape")
	}

	func display() {
		println("display")
	}

## Copyright
Copyright 2014 Vitali Baumtrok

This program is distributed under the terms of the Boost Software License,
Version 1.0, as published on <http://www.boost.org/LICENSE_1_0.txt>.
