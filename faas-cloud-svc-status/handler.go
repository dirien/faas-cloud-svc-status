package function

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dirien/status-keeper/statuspage"
	handler "github.com/openfaas/templates-sdk/go-http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	statusPages := statuspage.GetStatusPages()

	var status *statuspage.StatusPage

	for _, t := range statusPages {
		if t.Name == req.QueryString {
			status = &t
			break
		}
	}

	if status == nil {
		fmt.Println(fmt.Errorf("cannot get status: %s", req.QueryString))
		return handler.Response{}, nil
	}

	fmt.Printf("Status %s\n", status.Name)
	statusResponse, err := statuspage.Download(status)
	if err != nil {
		fmt.Println(err)
	}
	e, err := json.Marshal(statusResponse)
	if err != nil {
		fmt.Println(err)
	}
	return handler.Response{
		Body: e,
		Header: map[string][]string{
			"Content-Type": []string{"application/json"}},
		StatusCode: http.StatusOK,
	}, err
}
