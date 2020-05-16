package utils

import "time"

func StringToTimestampPostgres(arg string) (time.Time, error) {
	tspFormat := "2006-01-02T15:04:05-0700"
	layout := time.RubyDate

	t, err := time.Parse(layout, arg)
	if err != nil {
		return time.Now(), err
	}
	result, err := time.Parse(tspFormat, t.String())
	if err != nil {
		return time.Now(), err
	}
	return result, nil
}
