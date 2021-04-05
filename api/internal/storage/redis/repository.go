package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

// List - listing methods
type ListCache struct{}

func (r *ListCache) Get(channelID string) (*listing.Message, rest.Err) {
	result, err := GetClient().Get(ctx, channelID).Result()
	log.Println("GET cache result", result)
	if err != nil {
		return nil, rest.NewNotFoundError(fmt.Sprintf("error message not found with key: %v\n%vresult: ", err, result))
	}

	message := &listing.Message{}
	// unmarshal json into message object
	err = json.Unmarshal([]byte(result), message)
	if err != nil {
		return nil, rest.NewInternalServerError("error unmarshalling json", err)
	}

	log.Printf("Redis: %v : %v\n", channelID, *message)

	return message, nil
}

// Blog - blogging methods
type BlogCache struct{}

func (r *BlogCache) Set(message *listing.Message) rest.Err {

	json, err := json.Marshal(*message)
	if err != nil {
		return rest.NewInternalServerError("error unmarshalling json", err)
	}

	err = GetClient().Set(ctx, message.ChannelID, json, 10).Err()
	if err != nil {
		return rest.NewInternalServerError("error saving message in cache", err)
	}

	lc := &ListCache{}
	reply, restErr := lc.Get(message.ChannelID)
	if restErr != nil {
		log.Println("cannot find key: ChannelID")
		return restErr
	}
	log.Println("redis cache: ", reply)

	return nil
}

func (c *BlogCache) Delete(channelID string) rest.Err {
	err := GetClient().Del(ctx, channelID).Err()
	if err == redis.Nil {
		return rest.NewInternalServerError("error deleting key:value in cache", err)
	}

	return nil
}
