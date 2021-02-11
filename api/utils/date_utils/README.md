# Dates 



the purpose of creating a constant for the date format is to be able to test that the format is correct.
``` Go
const (
	apiDateLayout = "2006-01-02:02T15:04:05Z"
)
```

the reason `GetNow()` function is in a function is because we will use it in the future. 
``` Go 
func GetNow() time.Time {
	return time.Now().UTC()
}
```


