package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/roku"
   "io"
   "net/http"
   "os"
)

func newMaster(id string, bandwidth int, info bool) error {
   content, err := roku.NewContent(id)
   if err != nil {
      return err
   }
   video, err := content.Video()
   if err != nil {
      return err
   }
   fmt.Println("GET", video.URL)
   res, err := http.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.NewScanner(res.Body).Master(res.Request.URL)
   if err != nil {
      return err
   }
   stream := master.Stream(bandwidth)
   if info {
      fmt.Println(content)
      for _, each := range master.Streams {
         if each.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      return download(stream, content.Base())
   }
   return nil
}

func download(stream *hls.Stream, base string) error {
   fmt.Println("GET", stream.URI)
   res, err := http.Get(stream.URI.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.NewScanner(res.Body).Segment(res.Request.URL)
   if err != nil {
      return err
   }
   file, err := os.Create(base + hls.TS)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.ProgressChunks(file, len(seg.Info))
   for _, info := range seg.Info {
      res, err := http.Get(info.URI.String())
      if err != nil {
         return err
      }
      pro.AddChunk(res.ContentLength)
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}