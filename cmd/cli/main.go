package main

import (
	"context"
	"fmt"

	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/maximhq/bifrost/core"
	"github.com/maximhq/bifrost/core/schemas"
	"github.com/terriblethinking/cli/internal/app"
	"github.com/terriblethinking/core/account"
	"github.com/terriblethinking/core/config"
)

func main() {
	// var client *bifrost.Bifrost
	// var initErr error

	client, initErr := bifrost.Init(context.Background(), schemas.BifrostConfig{
		Account: &account.Account{},
	})

	for {
		if initErr != nil {

			if initErr.Error() == "base config does not exist" {
				fmt.Println("here")
				err := config.InitAll()

				if err != nil {
					panic(err)
				}
			} else {
				panic(initErr)
			}

		} else {

			break

		}

		client, initErr = bifrost.Init(context.Background(), schemas.BifrostConfig{
			Account: &account.Account{},
		})
	}

	defer client.Shutdown()

	p := tea.NewProgram(app.New(*client))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
