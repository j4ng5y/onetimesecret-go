package onetimesecret

import "net/http"

// Client is the main client for performing actions against the https://onetimesecret.com/ service
type Client struct {
	otsURL     string
	creds      *Credentials
	httpClient *http.Client
}

// Credentials are your https://onetimesecret.com user credentials to interact with the service API
type Credentials struct {
	// Username is your https://onetimesecret.com/ username.
	// See https://onetimesecret.com/signup to create a username.
	Username string

	// APIToken is your https://onetimesecret.com/ API token.
	// See https://onetimesecret.com/account#apikey-tab to view/generate one.
	APIToken string
}

// ClientOptions are provided to adjust and supplement the functionality of the Client
type ClientOptions struct {
	OneTimeSecretURL string
	Credentials      *Credentials
	HTTPClient       *http.Client
}

// New will generate a new Client with the default HTTP client
//
// Variables:
//     credentials (*Credentials): A pointer to a Credentials struct
//
// Returns:
//     (*Client): A pointer to a new instance of Client
func New(credentials *Credentials) *Client {
	var C Client
	C.otsURL = "https://onetimesecret.com"
	C.creds = credentials
	C.httpClient = http.DefaultClient
	return &C
}

// NewWithOptions will generate a new Client with a custom HTTP Client
//
// Variables:
//     opts (*ClientOptions): A pointer to a ClientOptions struct
//
// Returns:
//     (*Client): A pointer to a new instance of Client
func NewWithOptions(opts *ClientOptions) *Client {
	var C Client
	C.otsURL = opts.OneTimeSecretURL
	C.creds = opts.Credentials
	C.httpClient = opts.HTTPClient
	return &C
}
