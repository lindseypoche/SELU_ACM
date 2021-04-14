
## --- INTERNAL packages ---

### <verb>ing 
Pacakages that are named after verbs are part of the domain layer.
They contain the service and the models that it needs to respond to requests. 
These models are used in their own packages but also used in the storage package,
but only when the package needs to return a request to the client. 
The thing that is requested must be in it's `<verb>ing` package. 
Ie the request will contain the fields of the model which is stored in the `<verb>ing` package. 

### utils
General utilities that any package can use. 

### http 
rest
- http event handlers between the frontend and backend. 
discord
- webhook event handlers between discord and backend. 

## --- MongoDB QUERIES ---

This is a list of repository method queries that should be approved before deciding the database schema. 

#### ~~architecting repo~~ (postponed)
```go 
// used via discord
// store channels in the Channel document. 
// all other documents should reference a channel using `ChannelID: primitive.ObjectID 'json:"_id" bson:"_id"'`
type Repo struct {
    SaveChannel() rest.Err

    UpdateChannel() rest.Err

    DeleteChannel() rest.Err

    CreateRole() rest.Err

    DeleteRole() rest.Err

    UpdateRole() rest.Err
}
```

#### blogging repo
```go 
// used via discord
// store blogs/events/news in the Message document
// all other documents should reference a message using `MessageID: primitive.ObjectID 'json:"_id" bson:"_id"'`
type Repo struct {
    SaveMessage() rest.Err

    EditMessage() rest.Err

    DeleteMessage() rest.Err
}
```

#### ~~commenting repo~~ (postponed)
```go
// used via discord
// store comments with its correspoinding message in the Message document. append to array. 
// all other documents should reference a message using `MessageID: primitive.ObjectID 'json:"_id" bson:"_id"'`
type Repo struct {
    // a comment is added when a user explicitely replies to a message. 
    AddComment() rest.Err

    EditComment() rest.Err

    DeleteComment() rest.Err 

    AddReaction() rest.Err

    DeleteReaction() rest.Err
}
```

#### listing repo
```go
// used via frontend fetch calls 
// simply get calls 
type Repo struct {
    // get all message in chronological order. 
    // maybe set a limit though
    GetAllMessages() (*[]listing.Message, rest.Err)

    // get message by id from url parameter 
    GetMessageByID(string id) (*listing.Message, rest.Err)

    // gets all messages associated with a channel. set limit?
    GetMessagesByChannelID(string id) (*[]listing.Message, rest.Err)

    // get messages by username from url parameter
    GetMessagesByUsername(string username) (*listing.Message, rest.Err)

    // get from? mongo channel document? redis? local cache? 
    GetMessageByLatestPin(string channelID) (*[]listing.Message, rest.Err)

    // (postponed)
    ~~GetMessageByMostEmojis()~~ 

    // (postponed)
    ~~FindAllMessagesByTag(string tag)~~

    // (postponed)
    ~~GetAllArchives()~~ // for viewing archives. MessageID, Timestamp, Title, 

    // get all officers from users db. set limit?
    GetAllOfficers() (*[]listing.Member, rest.Err)

    // (postponed)
    ~~GetActiveOfficers()~~ (*[]listing.Message, rest.Err)
    
    // (postponed)
    ~~GetOfficersByYear()~~
    
    // (postponed)
    ~~GetOfficersByTitle()~~
}
```

#### ~~subscribing repo~~ (postponed)
```go
// used via discord 
// create a user mongo document 
type Repo struct {
    // add a user when they join the server
    // if a user joins that was once an officer, their officer status will remain inactive,
    // but they will not have any roles. 
    // when they are added an officer role, then simply set officer status to active. 
    AddUser() rest.Err
    
    // note, if user is an officer he/she will not be deleted from the user database. 
    // rather, status will be inactive or something
    DeleteUser() rest.Err 

    AddRoleToUser() rest.Err

    RemoveRoleFromUser() rest.Err
}
```

# -------

notice: find by mongo id should only be used between mongo Documents. 
- eg if you need to reference one document with another, then reference it with the mongo ID.