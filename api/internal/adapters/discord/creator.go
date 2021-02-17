package discord

// Creator is the creator of a message in a channel
type Creator struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Email         string `json:"email"`
}

// Discord User
// {
// 	"id": "80351110224678912",
// 	"username": "Nelly",
// 	"discriminator": "1337",
// 	"avatar": "8342729096ea3675442027381ff50dfe",
// 	"verified": true,
// 	"email": "nelly@discord.com",
// 	"flags": 64,
// 	"premium_type": 1,
// 	"public_flags": 64
// }
