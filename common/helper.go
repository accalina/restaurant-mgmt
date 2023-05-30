package common

import "time"

func DateStringToDatetime(dateString string) (result time.Time, err error) {
	dateLayout := "2006-01-02"
	result, err = time.Parse(dateLayout, dateString)
	if err != nil {
		return
	}

	return
}
