package wolf

import (
	"fmt"
	"flag"
	"io/ioutil"

	"github.com/JoshuaDoes/go-wolfram"
)

func main() {
	flag.Parse()

	Bid, err := ioutil.ReadFile("appid.txt")
	id := string(Bid)
	if err != nil { fmt.Errorf("%v\n", err) }
	fmt.Printf("Starting up with id [%s]\n", id)
	//Initialize a new client
	c := &wolfram.Client{AppID:id}

	fmt.Printf("Searching for '%s'\n", flag.Arg(0))
	res, err := c.GetQueryResult(flag.Arg(0), nil)

	if err != nil {
		fmt.Errorf("%v\n", err)
	}
	
	resObj := res.QueryResult

	//Iterate through the pods and subpods
	//and print out their title attributes
	for i := range resObj.Pods {
		fmt.Println(resObj.Pods[i].Title)

		for j := range resObj.Pods[i].SubPods {
			fmt.Println(resObj.Pods[i].SubPods[j].Title)
		}
	}
}