package discord

import (
	"errors"
	"strconv"
	"strings"
)

// func getAttachments(attachments []*discordgo.MessageAttachment) []*blogging.MessageAttachment {
// 	var ats []*blogging.MessageAttachment

// 	for _, attachment := range attachments {
// 		at := *&blogging.MessageAttachment{}
// 		ats = append(ats, attachment)
// 	}
// 	return &blogging.MessageAttachment{
// 		ID:       attachments[0].ID,
// 		URL:      attachments[0].URL,
// 		Filename: attachments[0].Filename,
// 		Width:    attachments[0].Width,
// 		Height:   attachments[0].Height,
// 		Size:     attachments[0].Size,
// 	}
// 	return nil
// }

// snowflakeToUnix converts snowflake id to a unix
func snowflakeToUnix(snowflake string) int {
	v, _ := strconv.Atoi(snowflake)
	x := v>>22 + discordEpoch
	s := strconv.Itoa(x)
	v, _ = strconv.Atoi(s[:len(s)-3])
	return v
}

// Get substring between two strings.
func between(value string, a string, b string) string {
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

// Get substring before a string.
func before(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

// Get substring after a string.
func after(value string, a string) (string, error) {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return "", errors.New("parsing error")
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return "", errors.New("parsing error")
	}
	return string(value[adjustedPos:len(value)]), nil
}
