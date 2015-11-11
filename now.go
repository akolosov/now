package now

import (
  "errors"
  "regexp"
  "time"
)

var FirstDayMonday = true

var TimeFormats = []string{"02.01.2006", "02.01.2006 15:04", "02.01.2006 15:04:05", "15:04 02.01.2006", "15:04:05 02.01.2006",
                           "2006-01-02", "2006-01-02 15:04", "2006-01-02 15:04:05", "15:04 2006-01-02", "15:04:05 2006-01-02",
                           "02/01/2006", "02/01/2006 15:04", "02/01/2006 15:04:05", "15:04 02/01/2006", "15:04:05 02/01/2006",
                           "15:04:05", "15:04"}

type Now struct {
  time.Time
}

func NewNow(t time.Time) *Now {
  return &Now{t}
}

func BeginningOfMinute() time.Time {
  return NewNow(time.Now()).BeginningOfMinute()
}

func BeginningOfHour() time.Time {
  return NewNow(time.Now()).BeginningOfHour()
}

func BeginningOfDay() time.Time {
  return NewNow(time.Now()).BeginningOfDay()
}

func BeginningOfWeek() time.Time {
  return NewNow(time.Now()).BeginningOfWeek()
}

func BeginningOfMonth() time.Time {
  return NewNow(time.Now()).BeginningOfMonth()
}

func BeginningOfYear() time.Time {
  return NewNow(time.Now()).BeginningOfYear()
}

func EndOfMinute() time.Time {
  return NewNow(time.Now()).EndOfMinute()
}

func EndOfHour() time.Time {
  return NewNow(time.Now()).EndOfHour()
}

func EndOfDay() time.Time {
  return NewNow(time.Now()).EndOfDay()
}

func EndOfWeek() time.Time {
  return NewNow(time.Now()).EndOfWeek()
}

func EndOfMonth() time.Time {
  return NewNow(time.Now()).EndOfMonth()
}

func EndOfYear() time.Time {
  return NewNow(time.Now()).EndOfYear()
}

func NextDay() time.Time {
  return NewNow(time.Now()).NextDay()
}

func PrevDay() time.Time {
  return NewNow(time.Now()).PrevDay()
}

func MonthLength() int {
  return NewNow(time.Now()).MonthLength()
}

func Parse(strs ...string) (time.Time, error) {
  return NewNow(time.Now()).Parse(strs...)
}

func MustParse(strs ...string) time.Time {
  return NewNow(time.Now()).MustParse(strs...)
}

func (now *Now) BeginningOfMinute() time.Time {
  return now.Truncate(time.Minute)
}

func (now *Now) BeginningOfHour() time.Time {
  return now.Truncate(time.Hour)
}

func (now *Now) BeginningOfDay() time.Time {
  d := time.Duration(-now.Hour()) * time.Hour
  return now.BeginningOfHour().Add(d)
}

func (now *Now) BeginningOfWeek() time.Time {
  t := now.BeginningOfDay()
  weekday := int(t.Weekday())
  if FirstDayMonday {
    if weekday == 0 {
      weekday = 7
    }
    weekday = weekday - 1
  }

  d := time.Duration(-weekday) * 24 * time.Hour
  return t.Add(d)
}

func (now *Now) BeginningOfMonth() time.Time {
  t := now.BeginningOfDay()
  d := time.Duration(-int(t.Day())+1) * 24 * time.Hour
  return t.Add(d)
}

func (now *Now) BeginningOfYear() time.Time {
  t := now.BeginningOfDay()
  d := time.Duration(-int(t.YearDay())+1) * 24 * time.Hour
  return t.Truncate(time.Hour).Add(d)
}

func (now *Now) EndOfMinute() time.Time {
  return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

func (now *Now) EndOfHour() time.Time {
  return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

func (now *Now) EndOfDay() time.Time {
  return now.BeginningOfDay().Add(24*time.Hour - time.Nanosecond)
}

func (now *Now) EndOfWeek() time.Time {
  return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (now *Now) EndOfMonth() time.Time {
  return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfYear() time.Time {
  return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (now *Now) NextDay() time.Time {
  return now.Add(24 * time.Hour)
}

func (now *Now) PrevDay() time.Time {
  return now.Add(-24 * time.Hour)
}

func (now *Now) MonthLength() int {
    return time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func parseWithFormat(str string) (t time.Time, err error) {
  for _, format := range TimeFormats {
    t, err = time.Parse(format, str)
    if err == nil {
      return
    }
  }
  err = errors.New("Can't parse string as time: " + str)
  return
}

func (now *Now) Parse(strs ...string) (t time.Time, err error) {
  var setCurrentTime bool
  parseTime := []int{}
  currentTime := []int{now.Second(), now.Minute(), now.Hour(), now.Day(), int(now.Month()), now.Year()}
  currentLocation := now.Location()

  for _, str := range strs {
    onlyTime := regexp.MustCompile(`^\s*\d+(:\d+)*\s*$`).MatchString(str) // match 15:04:05, 15

    t, err = parseWithFormat(str)
    location := t.Location()
    if location.String() == "UTC" {
      location = currentLocation
    }

    if err == nil {
      parseTime = []int{t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
      onlyTime = onlyTime && (parseTime[3] == 1) && (parseTime[4] == 1)

      for i, v := range parseTime {
        // Fill up missed information with current time
        if v == 0 {
          if setCurrentTime {
            parseTime[i] = currentTime[i]
          }
        } else {
          setCurrentTime = true
        }

        // Default day and month is 1, fill up it if missing it
        if (i == 3 || i == 4) && onlyTime {
          parseTime[i] = currentTime[i]
        }
      }
    }

    if len(parseTime) > 0 {
      t = time.Date(parseTime[5], time.Month(parseTime[4]), parseTime[3], parseTime[2], parseTime[1], parseTime[0], 0, location)
      currentTime = []int{t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
    }
  }
  return
}

func (now *Now) MustParse(strs ...string) (t time.Time) {
  t, _ = now.Parse(strs...)
  return t
}
