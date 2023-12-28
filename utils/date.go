package utils

import "time"

func FormartDateTime(dt time.Time) string {
	if dt.IsZero() {
		return ""
	}
	return dt.Format("2006-01-02 15:04:05")
}
