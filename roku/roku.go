package roku

import (
   "encoding/json"
   "fmt"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
   "time"
)

var LogLevel format.LogLevel

type Content struct {
   Meta struct {
      MediaType string
   }
   Title string
   Series struct {
      Title string
   }
   SeasonNumber string
   EpisodeNumber string
   ReleaseDate string
   RunTimeSeconds int64
   ViewOptions []struct {
      License string
      Media struct {
         Videos []Video
      }
   }
}

func NewContent(id string) (*Content, error) {
   var addr url.URL
   addr.Scheme = "https"
   addr.Host = "content.sr.roku.com"
   addr.Path = "/content/v1/roku-trc/" + id
   addr.RawQuery = url.Values{
      "expand": {"series"},
      "include": {strings.Join([]string{
         "episodeNumber",
         "releaseDate",
         "runTimeSeconds",
         "seasonNumber",
         // this needs to be exactly as is, otherwise size blows up
         "series.seasons.episodes.viewOptions\u2008",
         "series.title",
         "title",
         "viewOptions",
      }, ",")},
   }.Encode()
   var buf strings.Builder
   buf.WriteString("https://therokuchannel.roku.com/api/v2/homescreen/content/")
   buf.WriteString(url.PathEscape(addr.String()))
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

func (c Content) Base() string {
   var buf strings.Builder
   buf.WriteString(c.Series.Title)
   buf.WriteByte('-')
   buf.WriteString(c.Title)
   buf.WriteByte('-')
   buf.WriteString(c.SeasonNumber)
   buf.WriteByte('-')
   buf.WriteString(c.EpisodeNumber)
   return buf.String()
}

func (c Content) Duration() time.Duration {
   return time.Duration(c.RunTimeSeconds) * time.Second
}

func (c Content) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "Type:", c.Meta.MediaType)
   fmt.Fprintln(f, "Title:", c.Title)
   if c.Meta.MediaType == "episode" {
      fmt.Fprintln(f, "Series:", c.Series.Title)
      fmt.Fprintln(f, "Season:", c.SeasonNumber)
      fmt.Fprintln(f, "Episode:", c.EpisodeNumber)
   }
   fmt.Fprintln(f, "Date:", c.ReleaseDate)
   fmt.Fprint(f, "Duration: ", c.Duration())
   if verb == 'a' {
      for _, opt := range c.ViewOptions {
         fmt.Fprint(f, "\nLicense: ", opt.License)
         for _, vid := range opt.Media.Videos {
            fmt.Fprint(f, "\nURL: ", vid.URL)
         }
      }
   }
}

func (c Content) Video() (*Video, error) {
   for _, opt := range c.ViewOptions {
      if opt.License == "Free" {
         for _, vid := range opt.Media.Videos {
            if vid.DrmAuthentication == nil {
               return &vid, nil
            }
         }
      }
   }
   return nil, errorString("drmAuthentication")
}

type Video struct {
   DrmAuthentication *struct{}
   URL string
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
