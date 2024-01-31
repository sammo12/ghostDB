# ghostDB
a minimal and easy to use database written in go for better CRUD operations
# Installation
`go get www.github.com/sammo12/ghostDB`
# Getting Started
here is a golang example
```go
package main
import (
db "github.com/sammo12/ghostdb/ghostdb-go"
  "fmt"
)
func main()  {

db.Init()
\\reqired to initialize the database

db.Set("collection-1","examplekey","examplevalue")
\\ creating data in in collection 'collection-1'

db.Get("collection-1","examplekey")
\\ reading the data 
  
db.Del("collection-1","examplekey")
\\ delete the data
  
db.Delcol("collection-1")
\\ delete the whole data collections

db.Update("collection-1","examplekey","an exaple value")
\\ updating data

db.Getall("collection-1")
\\get all data from entire collection
}

```
