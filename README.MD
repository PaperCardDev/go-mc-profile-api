## 使用

`go get github.com/PaperCardDev/go-mc-profile-api`

## 代码示例：根据名字或ID获取（正版MC）

```go
package main

import (
	"fmt"
	"github.com/PaperCardDev/go-mc-profile-api"
)

func main() {
	p := go_mc_profile_api.McProfile{}

	err := p.FetchByName("jeb_")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Name: jeb_
	// Id: 853c80ef3c3749fdaa49938b674adae6
	fmt.Println("Name:", p.Name)
	fmt.Println("Id:", p.PlayerId)

	p2 := go_mc_profile_api.McProfile{}
	err = p2.FetchById("853c80ef3c3749fdaa49938b674adae6")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Name: jeb_
	// Id: 853c80ef3c3749fdaa49938b674adae6
	fmt.Println("Name:", p2.Name)
	fmt.Println("Id:", p2.PlayerId)
}

```