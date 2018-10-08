// periBlink.go
// Rich Robinson
// Sept 2018

package periBlink

import (
    "time"
    "periph.io/x/periph/host"
    "periph.io/x/periph/host/rpi"
)

const (
    numPx = 8
// used luminance instead of brightness
// note values 0 to 31 instead of float
    luminance = 3
    redI = 0
    greenI = 1
    blueI =2
    lumI = 3
)

type Pix struct {
    red ,green, blue, lum int
}

type Blinkt struct {
    pix [4] int
}

var (
    gpioSetUp bool = false
    clearOnExit bool = true
    pix []Pix
    blinkt [numPx]Blinkt
)

func Exit() {
    if clearOnExit {
        Clear()
        Show()
    }
}

func SetLuminance( lum int) {
    for i := range blinkt { 
        blinkt[i].pix[lumI] = lum
    }
}

func Clear() {
    for i := range blinkt { 
        blinkt[i].pix[redI] = 0
        blinkt[i].pix[greenI] = 0
        blinkt[i].pix[blueI] = 0
    }
}

func writeByte( val int) {
    for i := 0; i < 8; i++ {
        x := val&128
        if x == 0 {
            rpi.P1_16.Out(false)
        }else {
            rpi.P1_16.Out(true)
        }
        rpi.P1_18.Out(true)
        val = val << 1
        rpi.P1_18.Out(false)
    }
}

func eof() {
    rpi.P1_16.Out(false)
    for i := 0; i < 36; i++ {
        rpi.P1_18.Out(true)
        rpi.P1_18.Out(false)
    }
}

func sof() {
    rpi.P1_16.Out(false)
    for i := 0; i < 32; i++ {
        rpi.P1_18.Out(true)
        rpi.P1_18.Out(false)
    }
}

func Show() {
    if gpioSetUp == false { Setup() }
    sof()
    for i:= range blinkt {
        r := blinkt[i].pix[ redI]
        g := blinkt[i].pix[ greenI]
        b := blinkt[i].pix[ blueI]
        l := blinkt[i].pix[ lumI]
        bitwise := 224
        writeByte(bitwise | l)
        writeByte(b)
        writeByte(g)
        writeByte(r)
    }
    eof()
}

func SetAll(r int, g int, b int, l int) {
    for i := 0; i < numPx; i++ {
        SetPixel( i, r&255, g&255, b&255, l&31)
    }
}

func SetPixel(p int, r int, g int, b int, l int) {
    blinkt[p].pix[redI] = r & 255
    blinkt[p].pix[greenI] = g & 255
    blinkt[p].pix[blueI] = b & 255
    blinkt[p].pix[lumI] = l & 31
}

func GetPixel(p int) ( r int, g int, b int, l int ) {
    r = blinkt[p].pix[ redI]
    g = blinkt[p].pix[ greenI]
    b = blinkt[p].pix[ blueI]
    l = blinkt[p].pix[ lumI]
    return r, g, b, l
}

func SetclearOnExit(ce bool) {
    clearOnExit = ce
}

func delay(ms int) {
    time.Sleep(time.Duration(ms) * time.Millisecond)
}

func Setup() {
    host.Init()
    gpioSetUp = true
}
