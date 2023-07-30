# ghostDB
a minimal and easy to use database written in go
# Installation
`go get www.github.com/sammo12/ghostDB`
# Getting Started
here is a golang example
```
package main
import (
  "github.com/sammo12/ghostdb/ghostdb-go"
  "fmt"
)
func main()  {

ghost.Init()
\\reqired to initialize the database

ghost.Set("collection-1","examplekey","examplevalue")
\\ creating data in in collection 'collection-1'

ghost.Get("collection-1","examplekey")
\\ reading the data 
  
ghost.Del("collection-1","examplekey")
\\ delete the data
  
ghost.Delcol("collection-1")
\\ delete the whole data collections

ghost.Update("collection-1","examplekey","an exaple value")
\\ updating data

}

```
