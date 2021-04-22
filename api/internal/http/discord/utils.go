package discord

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/blogging"
)

// Validate authors
func Validate(s *discordgo.Session, m *discordgo.Message) bool {

	hasAccess := false
	// validate user is not bot
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return hasAccess
	}

	// validate channels
	for _, channel := range Config.Channels {
		if channel == m.ChannelID {
			hasAccess = true
		}
	}

	// validate user roles
	for _, role := range m.Member.Roles {
		for _, accessRole := range Config.Roles {
			if role == accessRole && hasAccess == true {
				return hasAccess
			}
		}
	}
	return hasAccess
}

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

// discord response
type response struct {
	// if request was posted
	DidPost bool `json:"did_post"`

	// response fields
	Date  *Field `json:"date,omitempty"`
	Title *Field `json:"title,omitempty"`
	Body  *Field `json:"body,omitempty"`
}

type cause interface{}

// Field is a field in the response
type Field struct {
	// validated options
	Match *opt `json:"-"`

	// error
	Err string `json:"err,omitempty"`

	// cause of the error
	Cause cause `json:"cause,omitempty"`
}

// options for either string or integer
type opt struct {
	// for storing string types
	s string

	// for storing integer types
	i int
}

// parse the content from the message
func parseMessage(s string) (*response, error) {

	var resp response
	resp.DidPost = true // flag for should post

	// parse date
	re := regexp.MustCompile(`DATE[:\s-]+(.*)[\n+|TITLE]`)
	match := re.FindStringSubmatch(s)
	// if there is a match then parse it
	if len(match) > 1 {
		ts, err := parseTime(match[1])

		if err != nil {
			resp.DidPost = false
			resp.Date = &Field{
				Err:   errors.New("formatting_error").Error(),
				Cause: "possible date formatting error",
			}
		} else {

			resp.Date = &Field{
				Match: &opt{i: ts},
			}
		}
	} else {

		resp.DidPost = false
		resp.Date = &Field{
			Err:   errors.New("parsing_error").Error(),
			Cause: "possible misspelled DATE field",
		}
	}

	// parse title
	re = regexp.MustCompile(`TITLE[:\s-]+(.*)[\n+|BODY]`)
	match = re.FindStringSubmatch(s)
	if len(match) > 1 {

		resp.Title = &Field{
			Match: &opt{s: match[1]},
		}
	} else {

		resp.DidPost = false
		resp.Title = &Field{
			Err:   errors.New("parsing_error").Error(),
			Cause: "possible misspelled TITLE field",
		}
	}

	re = regexp.MustCompile(`BODY[:\s-]+(.*)`)
	match = re.FindStringSubmatch(s)
	if len(match) > 1 {

		resp.Body = &Field{
			Match: &opt{s: match[1]},
		}
	} else {

		resp.DidPost = false
		resp.Body = &Field{
			Err:   errors.New("parsing_error").Error(),
			Cause: "possible misspelled BODY field",
		}
	}

	if resp.DidPost == false {
		return &resp, errors.New("error from reading or maniuplating content")
	}

	return &resp, nil
}

func prettifyJSON(resp *response) string {

	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("error parsing data into json")
	}

	json := string(b)

	return fmt.Sprintf(">>> Error posting content\n %v", json)
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

func parseTime(entry string) (int, error) {

	layout := "01/02/2006 3:04 PM -0700 MST"
	dt, err := time.Parse(layout, strings.Trim(entry, " "))
	if err != nil {
		log.Println("Parsing error:", err)
		return 0, err
	}

	// return timestamp
	return int(dt.Unix()), nil
}
