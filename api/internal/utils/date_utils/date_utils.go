package date_utils

import (
	"fmt"
	"time"
)

const (
	apiDateLayout          = "2006-01-02:02T15:04:05Z"
	apiDbLayout            = "2006-01-02 15:04:05"
	discordTimestampLayout = time.RFC3339
)

// GetNow returns a Time type of the current time in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowUnix gets current time from unix epoch
func GetNowUnix() int {
	return int(time.Now().UTC().Unix())
}

// GetNowString returns a string type of the current time in UTC
// with a layouf of "2006-01-02:02T15:04:05Z"
func GetNowString() string {
	// return GetNow().Format(apiDateLayout)
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat returns a string type of the current time in UTC
// with a layouf of "2006-01-02 15:04:05"
func GetNowDBFormat() string {
	// return GetNow().Format(apiDbLayout)
	return GetNow().Format(apiDbLayout)
}

// ToUnixTimestamp converts a format to a unix timestamp
func ToUnixTimestamp(timestamp string) string {
	t, err := time.Parse(discordTimestampLayout, timestamp)
	if err != nil {
		return fmt.Sprintf("%v", t.Unix())
	}
	return timestamp
}
