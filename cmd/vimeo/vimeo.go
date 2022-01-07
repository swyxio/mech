package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/vimeo"
   "net/http"
   "os"
   "path"
   "strings"
)

func main() {
   var (
      info, verbose bool
      vFormat string
   )
   flag.StringVar(&vFormat, "f", "", "format")
   flag.BoolVar(&info, "i", false, "info only")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("vimeo [flags] [video ID]")
      flag.PrintDefaults()
      return
   }
   if verbose {
      format.Log.Level = 1
   }
   id := flag.Arg(0)
   videoID, err := vimeo.Parse(id)
   if err != nil {
      panic(err)
   }
   con, err := vimeo.NewConfig(videoID)
   if err != nil {
      panic(err)
   }
   vids, err := con.Videos()
   if err != nil {
      panic(err)
   }
   if info {
      fmt.Println("Owner:", con.Video.Owner.Name)
      fmt.Println("Title:", con.Video.Title)
      fmt.Println("Duration:", con.Video.Duration)
   }
   for _, vid := range vids {
      if info {
         fmt.Print("ID:", vid.ID)
         fmt.Print(" Width:", vid.Width)
         fmt.Print(" Height:", vid.Height)
         fmt.Println()
      } else if vid.ID == vFormat {
         err := download(con, vid.URL())
         if err != nil {
            panic(err)
         }
      }
   }
}

func download(con *vimeo.Config, addr string) error {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := con.Video.Owner.Name + "-" + con.Video.Title + path.Ext(addr)
   file, err := os.Create(strings.Map(format.Clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
