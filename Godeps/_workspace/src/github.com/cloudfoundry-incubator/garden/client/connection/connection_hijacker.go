package connection

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/garden/routes"
	"github.com/tedsuo/rata"
)

type hijackable struct {
	req               *rata.RequestGenerator
	noKeepaliveClient *http.Client
	dialer            func(string, string) (net.Conn, error)
}

func NewHijackStreamer(network, address string) HijackStreamer {
	dialer := func(string, string) (net.Conn, error) {
		return net.DialTimeout(network, address, time.Second)
	}

	return &hijackable{
		req:    rata.NewRequestGenerator("http://api", routes.Routes),
		dialer: dialer,
		noKeepaliveClient: &http.Client{
			Transport: &http.Transport{
				Dial:              dialer,
				DisableKeepAlives: true,
			},
		},
	}
}

func (h *hijackable) Hijack(handler string, body io.Reader, params rata.Params, query url.Values, contentType string) (net.Conn, *bufio.Reader, error) {
	request, err := h.req.CreateRequest(handler, params, body)
	if err != nil {
		return nil, nil, err
	}

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	if query != nil {
		request.URL.RawQuery = query.Encode()
	}

	conn, err := h.dialer("tcp", "api") // net/addr don't matter here
	if err != nil {
		return nil, nil, err
	}

	client := httputil.NewClientConn(conn, nil)

	httpResp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	if httpResp.StatusCode < 200 || httpResp.StatusCode > 299 {
		defer httpResp.Body.Close()

		errRespBytes, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, nil, fmt.Errorf("Backend error: Exit status: %d, error reading response body: %s", httpResp.StatusCode, err)
		}
		return nil, nil, fmt.Errorf("Backend error: Exit status: %d, message: %s", httpResp.StatusCode, errRespBytes)
	}

	hijackedConn, hijackedResponseReader := client.Hijack()

	return hijackedConn, hijackedResponseReader, nil
}

func (c *hijackable) Stream(handler string, body io.Reader, params rata.Params, query url.Values, contentType string) (io.ReadCloser, error) {
	request, err := c.req.CreateRequest(handler, params, body)
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	if query != nil {
		request.URL.RawQuery = query.Encode()
	}

	httpResp, err := c.noKeepaliveClient.Do(request)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode < 200 || httpResp.StatusCode > 299 {
		errResponse, err := ioutil.ReadAll(httpResp.Body)
		httpResp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("bad response: %s", httpResp.Status)
		}

		if httpResp.StatusCode == http.StatusServiceUnavailable {
			// The body has the actual error string formed at the server.
			return nil, garden.NewServiceUnavailableError(string(errResponse))
		}

		return nil, Error{httpResp.StatusCode, string(errResponse)}
	}

	return httpResp.Body, nil
}
