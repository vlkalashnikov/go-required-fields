# go-required-fields

```
type User struct {
	Login       string  `json:"login" required:"true"`      // required
	FirstName   string  `json:"firstName" required:"false"` // not required
	LastName    string  `json:"lastName" required:"-"`      // not required
	SurName     string  `json:"surName"`					// not required
}
```

Example of usage:
```go
package main

import (
	"fmt"
	"github.com/vlkalashnikov/go-required-fields"
)

func main() {
	user := User{
		Login: "login",
		FirstName: "firstName",
		LastName: "lastName",
		SurName: "surName",
	}

	err := required.Validate(&user)
	if err != nil {
		fmt.Println(err)
	}
}