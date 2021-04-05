package listing

// Pin stores data about pinned messages in a channel
type Pin struct {
	Message  *Message `json:"message" bson:"message"`
	PinnedAt int      `json:"pinned_at" bson:"pinned_at"`
}
