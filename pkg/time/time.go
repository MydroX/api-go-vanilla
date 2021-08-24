package time

import "time"

// GetTimeNow returns the current time in time.Time value
func GetTimeNowString() string {
	loc, _ := time.LoadLocation("Europe/Paris")
	t := time.Now().In(loc)
	ts := t.Format("2006-01-02 15:04:05")
	return ts
}

func GetTimeNow() time.Time {
	loc, _ := time.LoadLocation("Europe/Paris")
	return time.Now().In(loc)
}

// MySQLTimeToTime converts a MySQL time string to a time.Time value
func MySQLTimeToTime(mysqlTime string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", mysqlTime)
	return t
}

func TimeToMySQLTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
