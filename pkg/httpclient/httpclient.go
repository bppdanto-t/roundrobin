package httpclient

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func SendPostHTTPRequest(request string, address string) (*http.Response, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, address, strings.NewReader(request))
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}