package inx

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/iotaledger/iota.go/v4/nodeclient"
)

const (
	APIRoundTripperBaseURL = "inx://"
)

type APIRoundTripper struct {
	client INXClient
}

func NewAPIRoundTripper(client INXClient) *APIRoundTripper {
	return &APIRoundTripper{
		client: client,
	}
}

func (r *APIRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	apiReq := &APIRequest{
		Method:  req.Method,
		Path:    req.URL.RequestURI(),
		Headers: HeadersFromHTTPHeader(req.Header),
	}

	if req.Body != nil {
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(req.Body); err != nil {
			_ = req.Body.Close()

			return nil, err
		}

		if err := req.Body.Close(); err != nil {
			return nil, err
		}

		apiReq.Body = buf.Bytes()
	}

	apiResp, err := r.client.PerformAPIRequest(req.Context(), apiReq)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode:    int(apiResp.GetCode()),
		ProtoMajor:    1,
		ProtoMinor:    0,
		Header:        apiResp.HTTPHeader(),
		Body:          io.NopCloser(bytes.NewBuffer(apiResp.GetBody())),
		ContentLength: int64(len(apiResp.GetBody())),
		Request:       req,
	}, nil
}

func NewHTTPClientOverINX(client INXClient) *http.Client {
	return &http.Client{
		Transport: NewAPIRoundTripper(client),
	}
}

func (x *APIRequest) HTTPHeader() http.Header {
	httpHeader := http.Header{}
	for k, v := range x.GetHeaders() {
		for _, i := range strings.Split(v, ", ") {
			httpHeader.Add(k, i)
		}
	}

	return httpHeader
}

func (x *APIResponse) HTTPHeader() http.Header {
	httpHeader := http.Header{}
	for k, v := range x.GetHeaders() {
		for _, i := range strings.Split(v, ", ") {
			httpHeader.Add(k, i)
		}
	}

	return httpHeader
}

func HeadersFromHTTPHeader(headers http.Header) map[string]string {
	h := map[string]string{}
	for k := range headers {
		h[http.CanonicalHeaderKey(k)] = strings.Join(headers.Values(k), ", ")
	}

	return h
}

func NewNodeclientOverINX(client INXClient) (*nodeclient.Client, error) {
	return nodeclient.New(APIRoundTripperBaseURL, nodeclient.WithHTTPClient(NewHTTPClientOverINX(client)))
}
