package instagram

import (
   "encoding/base64"
   "encoding/binary"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
)

var logLevel format.LogLevel

func getID(shortcode string) (uint64, error) {
   for len(shortcode) <= 11 {
      shortcode = "A" + shortcode
   }
   buf, err := base64.URLEncoding.DecodeString(shortcode)
   if err != nil {
      return 0, err
   }
   return binary.BigEndian.Uint64(buf[1:]), nil
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

// This gets us image 1241 by 1241, but requires authentication.
func newMedia(id uint64) (*media, error) {
   buf := []byte("https://i.instagram.com/api/v1/media/")
   buf = strconv.AppendUint(buf, id, 10)
   buf = append(buf, "/info/"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {auth},
      "user-agent": {"Instagram 206.1.0.34.121 Android"},
   }
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   med := new(media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

type media struct {
   Items []item
}

type version struct {
   Width int
   Height int
   URL string
}

type item struct {
   Media_Type int
   Image_Versions2 struct {
      Candidates []version
   }
   Video_Versions []version
   Carousel_Media []struct {
      Media_Type int
      Image_Versions2 struct {
         Candidates []version
      }
      Video_Versions []version
   }
}
