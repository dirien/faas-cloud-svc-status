package main

import (
	"fmt"
	"github.com/dirien/status-keeper/statuspage"
	"sort"
)

func main() {
	statusPages := statuspage.GetStatusPages()
	sort.Sort(statusPages)
	name := "STACKIT"
	var status *statuspage.StatusPage

	for _, t := range statusPages {
		if t.Name == name {
			status = &t
			break
		}
	}

	if status == nil {
		fmt.Println(fmt.Errorf("cannot get status: %s", name))
		return
	}

	fmt.Printf("Status %s\n", status.Name)
	fmt.Println("here")
	resp, err := statuspage.Download(status)
	fmt.Println("here2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)

}
