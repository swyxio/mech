package main

import (
   "flag"
   "github.com/89z/mech/vimeo"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // f
   var height int
   flag.IntVar(&height, "f", 720, "target height")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      vimeo.LogLevel = 1
   }
   if address != "" {
      clip, err := vimeo.NewClip(address)
      if err != nil {
         panic(err)
      }
      if password != "" {
         err := doAuth(clip, password, height, info)
         if err != nil {
            panic(err)
         }
      } else {
         err := doAnon(clip, height, info)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
