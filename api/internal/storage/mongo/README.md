## storage/mongo


NOTE:
Datatypes that are in the mongo directory should only be used by files in the mongo directory, to store data into the database. 

Mongo's `primitive.ObjectID` type should only be used for referencing itself in other documents. 


### Other:

consider using this to allow any Repo types to be returned by a given function 
```go
type Repos struct {
	BlogRepo  BlogRepo
	ListRepo  ListRepo
	ReactRepo ReactRepo
}

func GetAList(...interface{}) &Repos {
    repos := &Repos{
        ListRepo: &ListRepo{}
    }
    return repos 
}
```


