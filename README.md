# Errors
Error as json []byte
Every error has a code (something like http codes) and string id.

```go
err := errors.Access("Wrong creds")
err.Ref("L")
fmt.Println(string(err))

// {"code":401,"id":"access","time":1598828671,"comment":"Wrong creds","ref":"L1598828671589406"}

```
## Create
```go
err := errors.New(503, "access", "Can't connect to database")
```
