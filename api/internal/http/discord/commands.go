package discord

// Commands

// var cmdErr = rest.NewBotCommandError("Command not recognized.")

// const (
// 	help = "!help"

// 	// !saveChannel --name events
// 	// !saveChannel --name officers
// 	saveChannel = "!saveChannel"
// )

// type commands struct{}

// // Do takes in a command and call
// func Do(command string) (string, rest.Err) {

// 	var msg string

// 	switch command {
// 	case help:
// 		// call help function
// 		msg = helpMenu()
// 	case saveChannel:
// 		// call saveChannel function
// 	default:
// 		// cannot read command
// 		return msg, cmdErr
// 	}

// 	return msg, nil
// }

// temporary help menu
func helpMenu() string {
	return `md
		>>> 
		**Bot Commands**: 
		!help
		!acm save channel events
		!acm save channel officers

		**Post**:
		_Post example_:
		DATE: 01/02/2006 3:04 PM -0500 CDT
		TITLE: ACM Game Night!
		BODY: The acm is having a game night next month ... 

		_Allowed_:
		- attachments (images)
		- roles (@acm, @officer, ...)

		**Comments**: 
		Replying to a post will display it has a comment on the web page.

		**Emojis**: 
		Adding an emoji to a post will display it on the web page. 
		`
}
