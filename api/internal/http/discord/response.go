package discord

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

// discord response
type response struct {
	// if request was posted
	DidPost bool `json:"did_post"`

	// response fields
	Date  *Field `json:"start_time,omitempty"`
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
	re := regexp.MustCompile(`DATE[:\s-/\n]+(.*)[\n+|TITLE]`)
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
	re = regexp.MustCompile(`TITLE[:\s-/\n]+(.*)[\n+|BODY]`)
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

	body, err := after(s, "BODY:")
	if err != nil || body == "" {
		resp.DidPost = false
		resp.Title = &Field{
			Err:   errors.New("parsing_error").Error(),
			Cause: "possible missing or misspelled BODY field",
		}
	} else {
		body = strings.TrimSpace(body)
		strings.TrimPrefix(body, ":")
		resp.Body = &Field{
			Match: &opt{s: body},
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

// ***Menu is the interface for a bot menu***
// type Menu interface {
// 	// Display will display any content that is fed into it
// 	// to the discord channel where the command originated.
// 	Display(interface{}) string

// 	// Command sends user commands for other methods to handle
// 	// before returning a respone back to the user.
// 	Command(string) interface{}

// 	// commands
// 	post()
// 	member()
// 	role()
// 	channel()
// }

// type menu struct {
// }

// func NewMenu() Menu {
// 	return &menu{}
// }

// func (m *menu) Command(string) interface{} {

// }

// func (m *menu) Display(interface{}) string {
// 	return fmt.Sprintf(
// 		`>>>
// 		Usage:
// 		!acm [COMMAND] [OPTIONS]

// 		Options:
// 		-h, --help

// 		Commands:
// 		post
// 			!acm create \t\t fast hand for creating a post. note: creates a post for the channel youre in.
// 			!acm post create \t\t Make a post with guidance.
// 				[OPTIONS]
// 				--date 03-17-2021:3:30 PM -0500 CDT
// 				--title Lorem Ipsum
// 				--body Lorem Ipsum dol sit a mir...
// 			!acm post inspect POST \t\t inspect a post
// 			!acm post ls \t\t show all posts
// 			!acm post link POST_ID \t\t post the url link of a post in current discord channel
// 		member
// 			!acm member inspect USERNAME|ID \t\t inspect a member or officer
// 			!acm member ls \t\t show all members
// 		role
// 			!acm role ls \t\t show a list of roles
// 			!acm role NAME|ID ls \t\t show a list of all members with a role
// 		channel
// 			!acm channel ls \t\t show a list of monitored channels.
// 				--[ocontrair] \t\t show a list of unmonitored channels.
// 			!acm channel add NAME|ID \t\t add a channel to the list of monitored.
// 			!acm channel NAME|ID view create --template "./web/pages/about/about.js" \t\t create a frontend view/page for the content of a channel. (too complex)
// 		`,
// 	)
// }

// // post command: if user is in the process of creating a post then store their changes to cache.

// func (m *menu) helpCreate() string {
// 	return fmt.Sprintf(
// 		`>>>
// 		Usage:
// 		!ACM create [OPTIONS]

// 		Options:
// 		-h, --help
// 		--date \t\t Start time (if an event).
// 		--title \t\t Title of the post.
// 		--body \t\t Body of the post.
// 		`,
// 	)
// }

// func (m *menu) helpEdit() {}

// func (m *menu) helpLs() {
// }

// func (m *menu) helpShow() {}

// func (m *menu) mk()   {}
// func (m *menu) edit() {}
// func (m *menu) ls() {
// 	// show a list of memebers or posts
// 	// returns:
// 	// ID 			   • Creator    • Channel • Date Created • Likes
// 	// 234597733823872 • QuantaCake • Events  • Apr 24, 2021 • 7
// }
// func (m *menu) cat() {}
