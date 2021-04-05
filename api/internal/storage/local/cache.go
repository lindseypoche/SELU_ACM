package local

import (
	"log"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
)

// Cache should store data that needs to be quickly accessed
// but cache is not a permanent means of storage.
// Data in local cache would likely also be stored in a
// db, but *could* be to store live data that doesn't
// get stored in a db.

var (
	pinCache = make(map[string]*listing.Message)
)

type BlogCache struct{}

func (c *BlogCache) SetMessage(message *listing.Message) rest.Err {
	pinCache[message.ChannelID] = message

	// check if message successfully stored
	if _, ok := pinCache[message.ChannelID]; !ok {
		return rest.NewNotFoundError("error value not found in cache")
	}
	// log.Println("GetMessage() result: ", result)
	log.Println("message successfully stored in cache: ", pinCache)
	return nil
}

func (c *BlogCache) EditMessage(message *listing.Message) rest.Err {
	// update the key with the new message
	pinCache[message.ChannelID] = message

	// check if message successfully stored
	if _, ok := pinCache[message.ChannelID]; !ok {
		return rest.NewNotFoundError("error message not found in cache")
	}
	log.Println("message successfully edited in cache: ", pinCache)
	return nil
}

func (c *BlogCache) DeleteMessage(channelID string) rest.Err {
	// check if item in cache
	if _, ok := pinCache[channelID]; !ok {
		log.Println("error message not found in cache")
		return rest.NewNotFoundError("error message not found in cache")
	}

	// delete item in pinCache where key = channel
	delete(pinCache, channelID)
	log.Println("message successfully deleted from cache: ", pinCache)

	return nil
}

// listing
type ListCache struct{}

// GetMessage searches cache from a channel id and returns a message if
// successful.
func (c *ListCache) GetMessage(channelID string) (*listing.Message, rest.Err) {

	result, ok := pinCache[channelID]
	log.Println("got pinned message")

	if !ok {
		log.Println("message not found in cache")
		return nil, rest.NewNotFoundError("error message not found in cache")
	}

	return result, nil
}

// GetMessages returns a list of all messages from cache
func (c *ListCache) GetMessages() (map[string]*listing.Message, rest.Err) {

	if len(pinCache) == 0 {
		return nil, rest.NewNotFoundError("pins not found in cache")
	}

	return pinCache, nil
}

type StatCache struct{}
