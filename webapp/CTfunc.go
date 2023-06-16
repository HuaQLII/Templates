package webapp

import "time"

//custom template functions

func (app *WebApp) UnixToTime(timestamp int64) string {
	t := time.Duration(timestamp) * time.Second
	return t.String()
}
func (app *WebApp) UnixToDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
