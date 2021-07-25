package main

import (
   "github.com/myusuf3/imghash"
   "image/jpeg"
   "os"
)

func main() {
   fa, err := os.Open("mb.jpg")
   if err != nil {
      panic(err)
   }
   defer fa.Close()
   ia, err := jpeg.Decode(fa)
   if err != nil {
      panic(err)
   }
   fb, err := os.Open("crop.jpg")
   if err != nil {
      panic(err)
   }
   defer fb.Close()
   ib, err := jpeg.Decode(fb)
   if err != nil {
      panic(err)
   }
   d := imghash.Distance(imghash.Average(ia), imghash.Average(ib))
   println(d)
}