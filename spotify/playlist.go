package main

import (
   "fmt"
   "net/http"
)

func main() {
   req, err := http.NewRequest(
      "HEAD",
      "https://api.spotify.com/v1/playlists/6rZ28nCpmG5Wo1Ik64EoDm",
      nil,
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("Authorization", "Bearer BQC1D86LnOO67vOcMVg0nBzMo3Sr5J6yS7SqSm610E-uLPonxBZ2Ava5QuZ2siQpB7iSnlx25Y0kgVHVPAE")
   q := req.URL.Query()
   q.Set("additional_types", "track,episode")
   q.Set("fields", "collaborative,description,followers(total),images,name,owner(display_name,id,images,uri),public,tracks(items(track.type,track.duration_ms),total),uri")
   q.Set("market", "US")
   req.URL.RawQuery = q.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", res)
}
