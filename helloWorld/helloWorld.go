package main

import "github.com/banthar/Go-SDL/sdl"

func main() {
  var hello *sdl.Surface
  var screen *sdl.Surface

  sdl.Init(sdl.INIT_EVERYTHING)
  screen = sdl.SetVideoMode(640, 480, 32, sdl.SWSURFACE)

  hello = sdl.Load("hello.bmp")

  screen.Blit(nil, hello, nil)

  screen.Flip()

  sdl.Delay(2000)
}
