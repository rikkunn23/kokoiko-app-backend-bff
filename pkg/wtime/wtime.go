package wtime

import (
	"context"
	"fmt"
	"time"
)
type ctxKey string
const (
  ctxFreezeTimeKey ctxKey = "freeze_time_key"
  layout           string = "2006-01-02 15:04:05"
  formatMDWeek     string = "1/2(Mon)"
  formatHMM        string = "%d:%02d"
  // DateFormatYYYYMMDD ...
  DateFormatYYYYMMDD string = "2006-01-02"
)
// Now は context.Value から現在時刻を取得する
func Now(ctx context.Context) time.Time {
  return ctx.Value(ctxFreezeTimeKey).(time.Time)
}
// SetNow は context.Value に now で指定された任意の時刻を格納する
func SetNow(ctx context.Context, tmpNow time.Time) context.Context {
  now := time.Date(
    tmpNow.Year(),
    tmpNow.Month(),
    tmpNow.Day(),
    tmpNow.Hour(),
    tmpNow.Minute(),
    tmpNow.Second(),
    0,
    tmpNow.Location(),
  )
  return context.WithValue(ctx, ctxFreezeTimeKey, now)
}
// TimeToString ...
func TimeToString(t time.Time) string {
  str := t.Format(layout)
  return str
}
// TimeToStringf ...
func TimeToStringf(t time.Time, format string) string {
  str := t.Format(format)
  return str
}
// StringToTime ...
func StringToTime(str string) (time.Time, error) {
  return time.ParseInLocation(layout, str, time.Local)
}
// PtrStringToTime ...
func PtrStringToTime(str *string) (*time.Time, error) {
  if str == nil {
    return nil, nil
  }
  t, err := time.ParseInLocation(layout, *str, time.Local)
  return &t, err
}
// GetMidnightTime 渡された日付の0時0分0秒を返す
func GetMidnightTime(t time.Time) time.Time {
  return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
// GetOneWeekDate 渡された日付の前日曜日〜次土曜日を返す
// 日曜日の日付が渡された場合はその日付を起点に次の土曜日までの日付を返す
// 土曜日の日付が渡された場合はその日付を起点に前の日曜日から渡された土曜日までの日付を返す
func GetOneWeekDate(t time.Time) (time.Time, time.Time) {
  switch t.Weekday() {
  case time.Sunday:
    return t, t.AddDate(0, 0, 6)
  case time.Monday:
    return t.AddDate(0, 0, -1), t.AddDate(0, 0, 5)
  case time.Tuesday:
    return t.AddDate(0, 0, -2), t.AddDate(0, 0, 4)
  case time.Wednesday:
    return t.AddDate(0, 0, -3), t.AddDate(0, 0, 3)
  case time.Thursday:
    return t.AddDate(0, 0, -4), t.AddDate(0, 0, 2)
  case time.Friday:
    return t.AddDate(0, 0, -5), t.AddDate(0, 0, 1)
  case time.Saturday:
    return t.AddDate(0, 0, -6), t
  default:
    return t, t
  }
}
// GetTruncateSeconds 与えられた時刻の秒を切り捨て、"00" に設定します。
func GetTruncateSeconds(t time.Time) time.Time {
  return t.Truncate(time.Minute)
}
// GetTruncateSecondsPtr はポインタ型の時刻の秒を切り捨て、"00" に設定します。
func GetTruncateSecondsPtr(t *time.Time) *time.Time {
  if t == nil {
    return nil
  }
  truncated := GetTruncateSeconds(*t)
  return &truncated
}
// ConvertMDWeekFormat 渡された日時を"M/D(week)"形式に変換する
func ConvertMDWeekFormat(t time.Time) string {
  return t.Format(formatMDWeek)
}
// ConvertHMMFormat 渡された日時の時刻を"H:mm"形式に変換する
func ConvertHMMFormat(t time.Time) string {
  return fmt.Sprintf(formatHMM, t.Hour(), t.Minute())
}
