package main

import (
	"fmt"
	"log"
	"flag"
	"github.com/VojtechVitek/go-trello"
	"os/user"
	"github.com/BurntSushi/toml"
)


type tomlConfig struct {
	Title string
	Trello trelloConfig
}

// created from from config file
type trelloConfig struct {
	Appkey   string
	Secret  string
	User   string
}

func main() {
	newStoryFlag := flag.String("create", "", "title of new story")
	flag.Parse()
	newStory := *newStoryFlag
	if newStory != "" {
		fmt.Println(newStory)
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	var cfg tomlConfig
	if _, err := toml.DecodeFile(usr.HomeDir+"/.dvwrkflw.toml", &cfg); err != nil {
		fmt.Println(err)
		return
	}

	var trelloCfg = cfg.Trello
	fmt.Println(trelloCfg.Appkey)
	fmt.Println(trelloCfg.Appkey)
	fmt.Println(trelloCfg.Secret)
	fmt.Println(trelloCfg.User)
	fmt.Println(usr)
	fmt.Printf("Database: %s", cfg.Trello.User)
	trello, err := trello.NewAuthClient(cfg.Trello.Appkey, &cfg.Trello.Secret)
	if err != nil {
		log.Fatal(err)
	}

	// User @trello
	user, err := trello.Member(cfg.Trello.User)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.FullName)
	// New Trello Client
	fmt.Sprint("%s\n", cfg.Trello.User)

	// @trello Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	if len(boards) > 0 {
		board := boards[0]
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		// @trello Board Lists
		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fmt.Println("   - ", list.Name)

			// @trello Board List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
}
