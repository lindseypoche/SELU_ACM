package discord

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// snowflakeToUnix converts snowflake id to a unix
func snowflakeToUnix(snowflake string) int {
	v, _ := strconv.Atoi(snowflake)
	x := v>>22 + discordEpoch
	s := strconv.Itoa(x)
	v, _ = strconv.Atoi(s[:len(s)-3])
	return v
}

// HandleJson allows extraction and modification of the acm.json file
func HandleJson(field, key, value string) error {
	file := "acm.json"
	// Read json buffer from jsonFile
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	var conf config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return err
	}

	switch field {
	case "channels":
		switch key {
		case "acm_events", "acm_officers":
			conf.Channels[key] = value
			break
		default:
			return errors.New(fmt.Sprintf("unrecognized field %s", field))
		}
	case "roles":
		switch key {
		case "acm_admin", "acm_officer", "acm_member":
			conf.Roles[key] = value
			break
		default:
			return errors.New(fmt.Sprintf("unrecognized field %s", field))
		}
	}

	log.Println("\njson:\n", conf)

	b, err = json.Marshal(conf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	return err
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
