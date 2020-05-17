package utils

import "time"

func StringToTimestampPostgres(arg string) (time.Time, error) {
	layout := time.RubyDate
	t, err := time.Parse(layout, arg)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
