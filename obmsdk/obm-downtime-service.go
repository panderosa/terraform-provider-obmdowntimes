package obmsdk

//petstore
import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-querystring/query"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

type Client struct {
	baseURL *url.URL
	headers http.Header
	http    *http.Client
	//limiter *rate.Limiter

	Downtimes Downtimes
}

func GetEnv(key string) *string {
	value, isDefined := os.LookupEnv(key)
	if isDefined {
		return &value
	} else {
		return nil
	}
}

// Instantiate a new OBM Downtime API REST Client
func NewClient(address *string, path *string, username *string, password *string) (*Client, error) {
	basicAuth := ""
	var baseURL *url.URL
	var headers = make(http.Header)
	if username != nil && password != nil {
		data := fmt.Sprintf("%v:%v", *username, *password)
		basicAuth = base64.StdEncoding.EncodeToString([]byte(data))
	}

	if (address == nil) || (path == nil) {
		return nil, fmt.Errorf("empty address and/or path")
	}

	plainText := fmt.Sprintf("%v%v", *address, *path)
	baseURL, err := url.Parse(plainText)
	if err != nil {
		return nil, fmt.Errorf("NewClient: failed to url-parse %v", plainText)
	}

	headers.Add("Content-Type", "application/xml")
	headers.Add("Accept", "application/xml")

	if basicAuth != "" {
		headers.Add("Authorization", fmt.Sprintf("Basic %v", basicAuth))
	}

	// Create the client
	client := &Client{
		baseURL: baseURL,
		headers: headers,
		http:    cleanhttp.DefaultPooledClient(),
	}

	client.Downtimes = &downtimes{client: client}

	return client, nil
}

func (c *Client) newRequest(method string, path string, v interface{}) (*http.Request, error) {
	u, err := url.Parse(c.baseURL.String() + path)
	if err != nil {
		return nil, err
	}

	var body io.Reader
	switch method {
	case "GET":
		if v != nil {
			q, err := query.Values(v)
			if err != nil {
				return nil, err
			}
			u.RawQuery = q.Encode()
		}
	case "POST", "PUT":
		if v != nil {
			dat, _ := xml.MarshalIndent(v, "", "  ")
			//log.Printf("[DEBUG] downtime body: " + string(dat))
			body = bytes.NewReader(dat)
		}
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	for k, v := range c.headers {
		req.Header[k] = v
	}

	return req, nil

}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)
	//log.Printf("[DEBUG] downtime requets: %v", req)

	// wake up the function ?
	tempReq, _ := c.newRequest("GET", "", nil)
	c.http.Do(tempReq)

	payload := ""
	if req.Body != nil {
		buf1 := new(bytes.Buffer)
		buf1.ReadFrom(req.Body)
		payload = buf1.String()
	}

	LogMe(req.Method, payload)

	resp, err := c.http.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}

	defer resp.Body.Close()
	//log.Printf("[DEBUG] downtime response: %v", resp)

	err = checkResponseCode(resp)
	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	// log body content
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	//log.Printf("[DEBUG] Body: %s", buf.String())
	LogMe("response", buf.String())
	return xml.Unmarshal(buf.Bytes(), v)
}

func checkResponseCode(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode <= 299 {
		return nil
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	return fmt.Errorf("HTTP Status Code: %v, Message: %v\n. %v", r.StatusCode, r.Status, buf.String())
}
