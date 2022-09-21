package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

var A float64
var B float64
var C float64
var cubeWidth float64 = 20
var width, height float64 = 160, 44
var zBuffer [160 * 44]byte
var buffer [160 * 44]byte
var backgroundASCIICode byte = '.' //?
var distanceFromCam float64 = 100
var horizontalOffset float64
var K1 float64 = 40
var incrementSpeed float64 = 0.6
var x, y, z float64
var ooz float64
var xp, yp float64
var idx float64

func calculateX(i float64, j float64, k float64) float64 {
	return j*math.Sin(A)*math.Sin(B)*math.Cos(C) - k*math.Cos(A)*math.Sin(B)*math.Cos(C) +
		j*math.Cos(A)*math.Sin(C) + k*math.Sin(A)*math.Sin(C) + i*math.Cos(B)*math.Cos(C)
}

func calculateY(i float64, j float64, k float64) float64 {
	return j*math.Cos(A)*math.Cos(C) + k*math.Sin(A)*math.Cos(C) -
		j*math.Sin(A)*math.Sin(B)*math.Sin(C) + k*math.Cos(A)*math.Sin(B)*math.Sin(C) - i*math.Cos(B)*math.Sin(C)
}

func calculateZ(i float64, j float64, k float64) float64 {
	return k*math.Cos(A)*math.Cos(B) - j*math.Sin(A)*math.Cos(B) + i*math.Sin(B)
}

func calculateForSurface(cubeX float64, cubeY float64, cubeZ float64, ch int64) {
	x = calculateX(cubeX, cubeY, cubeZ)
	y = calculateY(cubeX, cubeY, cubeZ)
	z = calculateZ(cubeX, cubeY, cubeZ) + distanceFromCam

	ooz = 1 / z

	xp = width/2 + horizontalOffset + K1*ooz*x*2
	yp = height/2 + K1*ooz*y

	idx = xp + yp*width
	if idx >= 0 && idx < width*height {
		if byte(ooz) > zBuffer[int(idx)] {
			zBuffer[int(idx)] = byte(ooz)
			buffer[int(idx)] = byte(ch)
		}
	}
}

func MemSet(s unsafe.Pointer, c byte, n uintptr) {
	ptr := uintptr(s)
	var i uintptr
	for i = 0; i < n; i++ {
		pByte := (*byte)(unsafe.Pointer(ptr + i))
		*pByte = c
	}
}

func main() {
	fmt.Print("\033[H\033[2J")
	for true {
		MemSet(unsafe.Pointer(&buffer), backgroundASCIICode, unsafe.Sizeof(width*height))
		MemSet(unsafe.Pointer(&zBuffer), 0, unsafe.Sizeof(width*height*4))
		cubeWidth = 20
		horizontalOffset = -2 * cubeWidth
		// first cube
		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateForSurface(cubeX, cubeY, -cubeWidth, '@')
				calculateForSurface(cubeWidth, cubeY, cubeX, '$')
				calculateForSurface(-cubeWidth, cubeY, -cubeX, '~')
				calculateForSurface(-cubeX, cubeY, cubeWidth, '#')
				calculateForSurface(cubeX, -cubeWidth, -cubeY, ';')
				calculateForSurface(cubeX, cubeWidth, cubeY, '+')
			}
		}
		cubeWidth = 10
		horizontalOffset = 1 * cubeWidth
		// second cube
		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateForSurface(cubeX, cubeY, -cubeWidth, '@')
				calculateForSurface(cubeWidth, cubeY, cubeX, '$')
				calculateForSurface(-cubeWidth, cubeY, -cubeX, '~')
				calculateForSurface(-cubeX, cubeY, cubeWidth, '#')
				calculateForSurface(cubeX, -cubeWidth, -cubeY, ';')
				calculateForSurface(cubeX, cubeWidth, cubeY, '+')
			}
		}
		cubeWidth = 5
		horizontalOffset = 8 * cubeWidth
		// third cube
		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateForSurface(cubeX, cubeY, -cubeWidth, '@')
				calculateForSurface(cubeWidth, cubeY, cubeX, '$')
				calculateForSurface(-cubeWidth, cubeY, -cubeX, '~')
				calculateForSurface(-cubeX, cubeY, cubeWidth, '#')
				calculateForSurface(cubeX, -cubeWidth, -cubeY, ';')
				calculateForSurface(cubeX, cubeWidth, cubeY, '+')
			}
		}
		fmt.Print("\033[H\033[2J")
		for k := 0; k < int(width*height); k++ {
			//C.putchar(k % width ? buffer[k] : 10)

			if k%int(width) == 0 {
				fmt.Print(rune(buffer[k]))
			} else {
				fmt.Print(rune(10))
			}
		}

		A += 0.05
		B += 0.05
		C += 0.01
		//C.usleep(8000 * 2)
		time.Sleep(8000 * 2)
	}
}
