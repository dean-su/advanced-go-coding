package main

import (
	"github.com/ChimeraCoder/anaconda"
	"fmt"
)

func main() {
	anaconda.SetConsumerKey("1oDpGoBRbfF9ZFeCznh68NT3W")
	anaconda.SetConsumerSecret("Xo4aVlLaseUNbhXECJNvF5tsdSpOAhhFSxCRos0ZBakeQYgvoH")
	api := anaconda.NewTwitterApi("4885835964-Wx022eBN6rfl5HtLvDroaIRDGPbfqfgM5M2eXz6", "kA3WOeffK68pu7XlMJ957Dfh28tTn76kpTLRxwj2UzH4w")

	searchResult, _ := api.GetSearch("golang", nil)
	for _ , tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}

	t, err := api.PostTweet("Hello Twitter API test:https://www.youtube.com/watch?v=YBo0RWyZJV4", nil)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(t)
}

