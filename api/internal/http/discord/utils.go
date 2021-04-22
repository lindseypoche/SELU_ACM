package discord

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/blogging"
)

func getAttachment(attachments []*discordgo.MessageAttachment) *blogging.MessageAttachment {
	if len(attachments) > 0 {
		return &blogging.MessageAttachment{
			ID:       attachments[0].ID,
			URL:      attachments[0].URL,
			Filename: attachments[0].Filename,
			Width:    attachments[0].Width,
			Height:   attachments[0].Height,
			Size:     attachments[0].Size,
		}
	}
	return nil
}

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
func after(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return string(value[adjustedPos])
}
