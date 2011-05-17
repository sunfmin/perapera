package main
import(
  "fmt"
  "math"
  "github.com/hoisie/web.go"
  "github.com/hoisie/mustache.go"
  "github.com/Philio/GoMySQL"
)

type Video struct {
  Id int
  site string
  title string
  code string
}

type Subtitle struct {
  VideoId int
  StartTime string
  StartTimeInSeconds float64
  EndTime string
  LanguageCode string
  Content string
  Position int
  HtmlContent string
}

func (v *Video) YouKu() bool {
  return v.site == "YouKu"
}

func (v *Video) Picasa() bool {
  return v.site == "Picasa"
}

func (v *Video) GoogleDoc() bool {
  return v.site == "GoogleDoc"
}

func (v *Video) DirectLink() bool {
  return v.site == "DirectLink"
}

func (v *Video) Other() bool {
  return !v.YouKu() && !v.Picasa() && !v.GoogleDoc() && !v.DirectLink()
}

func (v *Video) SubtitleLocked() bool {
  return false
}

func (v *Subtitle) StartTimeInSecondsFixed() float64 {
  //
  return v.StartTimeInSeconds
}

func (v *Subtitle) ShortStartTime() string {

  seconds := v.StartTimeInSecondsFixed()

  sign := ""
  if seconds < 0 {
    sign = "-"
  }

  m := int(math.Fabs(seconds) / 60)
  s := int(math.Fmod(math.Fabs(seconds),  60))

  return sign + fmt.Sprintf("%02d:%02d", m, s)
}
func (v *Subtitle) SplitToWords() string {
  return v.Content
}

var db *mysql.Client


func (v *Video) Subtitles() []*Subtitle {
  var result []*Subtitle



  stmt, err := db.Prepare("SELECT video_id, start_time, start_time_in_seconds, end_time, language_code, content, html_content FROM subtitles WHERE video_id = ? ORDER BY position")

  if err != nil {
    panic(err)
  }


  err = stmt.BindParams(v.Id)

  if err != nil {
    panic(err)
  }


  err = stmt.Execute()

  if err != nil {
    panic(err)
  }

  var myrow *Subtitle

  for {
      myrow = new(Subtitle)
      // Bind result
      stmt.BindResult(&myrow.VideoId, &myrow.StartTime, &myrow.StartTimeInSeconds, &myrow.EndTime, &myrow.LanguageCode, &myrow.Content)

      eof, err := stmt.Fetch()

      if err != nil {
        panic(err)
      }

      if eof {
          break
      }

      result = append(result, myrow)
  }

/*  fmt.Println(result)*/

  return result
}


func videoShow(ctx *web.Context, videoId string) string {
  return mustache.RenderFile("show.html", &Video{1, "YouKu", "The Lord of War", "EQWEQWEQWEQWEQWE"})
}

func main() {
  db, _ = mysql.DialTCP("localhost", "root", "", "lingotv_dev")

  web.Get("/videos/(.*)", videoShow)
  web.Run("0.0.0.0:9999")
}
