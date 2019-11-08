package onetimesecret

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// CreateSecretRequest is a structure that holds data that will be marshalled into https://onetimesecret.com query parameters
//
//  Query Params
//
//    secret: the secret value which is encrypted before being stored. There is a maximum length based on your plan that is enforced (1k-10k).
//    passphrase: a string that the recipient must know to view the secret. This value is also used to encrypt the secret and is bcrypted before being stored so we only have this value in transit.
//    ttl: the maximum amount of time, in seconds, that the secret should survive (i.e. time-to-live). Once this time expires, the secret will be deleted and not recoverable.
//    recipient: an email address. We will send a friendly email containing the secret link (NOT the secret itself).
type CreateSecretRequest struct {
	Secret     string
	Passphrase string
	TTL        int
	Recipient  []string
}

// Validate will verify that data in the parent data structure is present, and eventually, valid
//
// Variables:
//     None
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (C *CreateSecretRequest) Validate() error {
	if C.Secret == "" {
		return fmt.Errorf("secret can not be left blank")
	}

	return nil
}

// CreateSecretResponse is a structure that will hold data that is unmarshalled from a json response
//
//  Attributes
//
//    custid: this is you :]
//    metadata_key: the unique key for the metadata. DO NOT share this.
//    secret_key: the unique key for the secret you create. This is key that you can share.
//    ttl: The time-to-live (in seconds) that was specified (i.e. not the time remaining)
//    metadata_ttl: The remaining time (in seconds) that the metadata has left to live.
//    secret_ttl: The remaining time (in seconds) that the secret has left to live.
//    recipient: if a recipient was specified, this is an obfuscated version of the email address.
//    created: Time the secret was created in unix time (UTC)
//    updated: ditto, but the time it was last updated.
//    passphrase_required: If a passphrase was provided when the secret was created, this will be true. Otherwise false, obviously.
type CreateSecretResponse struct {
	CustID             string `json:"custid"`
	MetadataKey        string `json:"metadata_key"`
	SecretKey          string `json:"secret_key"`
	TTL                int `json:"ttl"`
	MetadataTTL        int `json:"metadata_ttl"`
	SecretTTL          int `json:"secret_ttl"`
	Recipient          []string `json:"recipient"`
	CreatedAt          int `json:"created"`
	UpdatedAt          int `json:"updated"`
	PassphraseRequired bool   `json:"passphrase_required"`
}

// Unmarshal will read a json formatted http response body and apply those fields to structure fields
//
// Variables:
//     httpResponseBody (io.ReadCloser): The interface returned from an http.Client.Do action
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (C *CreateSecretResponse) Unmarshal(httpResponseBody io.ReadCloser) error {
	b, err := ioutil.ReadAll(httpResponseBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, C)
}

// GenerateSecretRequest is a structure that holds data that will me marshalled into https://onetimesecret.com query parameters
//
//  Query Params
//
//    passphrase: a string that the recipient must know to view the secret. This value is also used to encrypt the secret and is bcrypted before being stored so we only have this value in transit.
//    ttl: the maximum amount of time, in seconds, that the secret should survive (i.e. time-to-live). Once this time expires, the secret will be deleted and not recoverable.
//    recipient: an email address. We will send a friendly email containing the secret link (NOT the secret itself).
type GenerateSecretRequest struct {
	Passphrase string
	TTL        int
	Recipient  []string
}

// Validate will verify that data in the parent data structure is present, and eventually, valid
//
// Variables:
//     None
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (G *GenerateSecretRequest) Validate() error {
	return nil
}

// GenerateSecretResponse is a structure that will hold data that is unmarshalled from a json response
//
//  Attributes
//
//    custid: this is you :]
//    metadata_key: the unique key for the metadata. DO NOT share this.
//    secret_key: the unique key for the secret you create. This is key that you can share.
//    ttl: The time-to-live (in seconds) that was specified (i.e. not the time remaining)
//    metadata_ttl: The remaining time (in seconds) that the metadata has left to live.
//    secret_ttl: The remaining time (in seconds) that the secret has left to live.
//    recipient: if a recipient was specified, this is an obfuscated version of the email address.
//    created: Time the secret was created in unix time (UTC)
//    updated: ditto, but the time it was last updated.
//    passphrase_required: If a passphrase was provided when the secret was created, this will be true. Otherwise false, obviously.
type GenerateSecretResponse struct {
	CustID             string `json:"custid"`
	Value              string `json:"value"`
	MetadataKey        string `json:"metadata_key"`
	SecretKey          string `json:"secret_key"`
	TTL                int `json:"ttl"`
	MetadataTTL        int `json:"metadata_ttl"`
	SecretTTL          int `json:"secret_ttl"`
	Recipient          []string `json:"recipient"`
	CreatedAt          int `json:"created"`
	UpdatedAt          int `json:"updated"`
	PassphraseRequired bool   `json:"passphrase_required"`
}

// Unmarshal will read a json formatted http response body and apply those fields to structure fields
//
// Variables:
//     httpResponseBody (io.ReadCloser): The interface returned from an http.Client.Do action
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (G *GenerateSecretResponse) Unmarshal(httpResponseBody io.ReadCloser) error {
	b, err := ioutil.ReadAll(httpResponseBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, G)
}

// RetrieveSecretRequest is a structure that holds data that will be marshalled into http://onetimesecret.com query parameters
//
//  Query Params
//
//    SECRET_KEY: the unique key for this secret.
//    passphrase (if required): the passphrase is required only if the secret was create with one.
type RetrieveSecretRequest struct {
	SecretKey  string
	Passphrase string
}

// Validate will verify that data in the parent data structure is present, and eventually, valid
//
// Variables:
//     None
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (R *RetrieveSecretRequest) Validate() error {
	if R.SecretKey == "" {
		return fmt.Errorf("SecretID must not be \"\"")
	}

	return nil
}

// RetrieveSecretResponse is a structure that will hold data that is unmarshalled from a json response
//
//   Attributes
//
//    secret_key: the unique key for the secret you create. This is key that you can share.
//    value: The actual secret. It should go without saying, but this will only be available one time.
type RetrieveSecretResponse struct {
	SecretKey   string `json:"secret_key"`
	SecretValue string `json:"value"`
}

// Unmarshal will read a json formatted http response body and apply those fields to structure fields
//
// Variables:
//     httpResponseBody (io.ReadCloser): The interface returned from an http.Client.Do action
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (R *RetrieveSecretResponse) Unmarshal(httpResponseBody io.ReadCloser) error {
	b, err := ioutil.ReadAll(httpResponseBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, R)
}

// RetrieveMetadataRequest is a structure that holds data required to perform an https://onetimesecret.com request
//
// Query Params
//
//    METADATA_KEY: the unique key for this metadata.
type RetrieveMetadataRequest struct {
	MetadataKey string
}

// Validate will verify that data in the parent data structure is present, and eventually, valid
//
// Variables:
//     None
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (R *RetrieveMetadataRequest) Validate() error {
	if R.MetadataKey == "" {
		return fmt.Errorf("MetadataKey must not be \"\"")
	}
	return nil
}

// RetrieveMetadataResponse is a structure that will hold data that is unmarshalled from a json response
//
//   Attributes
//
//    custid: this is you :]
//    metadata_key: the unique key for the metadata. DO NOT share this.
//    secret_key: the unique key for the secret you created. This is key that you can share.
//    ttl: The time-to-live that was specified (i.e. not the time remaining)
//    metadata_ttl: The remaining time (in seconds) that the metadata has left to live.
//    secret_ttl: The remaining time (in seconds) that the secret has left to live.
//    recipient: if a recipient was specified, this is an obfuscated version of the email address.
//    created: Time the metadata was created in unix time (UTC)
//    updated: ditto, but the time it was last updated.
//    received: Time the secret was received.
//    passphrase_required: If a passphrase was provided when the secret was created, this will be true. Otherwise false, obviously.
type RetrieveMetadataResponse struct {
	CustID             string `json:"custid"`
	MetadataKey        string `json:"metadata_key"`
	SecretKey          string `json:"secret_key"`
	TTL                int `json:"ttl"`
	MetadataTTL        int `json:"metadata_ttl"`
	SecretTTL          int `json:"secret_ttl"`
	Recipient          []string `json:"recipient"`
	CreatedAt          int `json:"created"`
	UpdatedAt          int `json:"updated"`
	Received           int `json:"received"`
	PassphraseRequired bool   `json:"passphrase_required"`
}

// Unmarshal will read a json formatted http response body and apply those fields to structure fields
//
// Variables:
//     httpResponseBody (io.ReadCloser): The interface returned from an http.Client.Do action
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (R *RetrieveMetadataResponse) Unmarshal(httpResponseBody io.ReadCloser) error {
	b, err := ioutil.ReadAll(httpResponseBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, R)
}

// BurnSecretRequest is a structure that holds data required to perform an https://onetimesecret.com request
//
//  Query Params
//
//    None
type BurnSecretRequest struct {
	MetadataKey string
}

// Validate will verify that data in the parent data structure is present, and eventually, valid
//
// Variables:
//     None
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (B *BurnSecretRequest) Validate() error {
	if B.MetadataKey == "" {
		return fmt.Errorf("MetadataKey must not be \"\"")
	}
	return nil
}

// BurnSecretResponse is a structure that will hold data that is unmarshalled from a json response
//
//   Attributes
//
//    custid: this is you :]
//    metadata_key: the unique key for the metadata. DO NOT share this.
//    secret_key: the unique key for the secret you created. This is key that you can share.
//    ttl: The time-to-live that was specified (i.e. not the time remaining)
//    metadata_ttl: The remaining time (in seconds) that the metadata has left to live.
//    secret_ttl: The remaining time (in seconds) that the secret has left to live.
//    recipient: if a recipient was specified, this is an obfuscated version of the email address.
//    created: Time the metadata was created in unix time (UTC)
//    updated: ditto, but the time it was last updated.
//    received: Time the secret was received.
//    passphrase_required: If a passphrase was provided when the secret was created, this will be true. Otherwise false, obviously.
type BurnSecretResponse struct {
	CustID             string `json:"custid"`
	MetadataKey        string `json:"metadata_key"`
	SecretKey          string `json:"secret_key"`
	TTL                int `json:"ttl"`
	MetadataTTL        int `json:"metadata_ttl"`
	SecretTTL          int `json:"secret_ttl"`
	Recipient          []string `json:"recipient"`
	CreatedAt          int `json:"created"`
	UpdatedAt          int `json:"updated"`
	Received           int `json:"received"`
	PassphraseRequired bool   `json:"passphrase_required"`
}

// Unmarshal will read a json formatted http response body and apply those fields to structure fields
//
// Variables:
//     httpResponseBody (io.ReadCloser): The interface returned from an http.Client.Do action
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (B *BurnSecretResponse) Unmarshal(httpResponseBody io.ReadCloser) error {
	b, err := ioutil.ReadAll(httpResponseBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, B)
}

// RetrieveRecentMetadataRequest is an empty struct for consistencies sake
type RetrieveRecentMetadataRequest struct{}

// RetrieveRecentMetadataResponse is a structure that will hold data that is unmarshalled from a json response
//
//   List of Attributes
//
//    custid: this is you :]
//    metadata_key: the unique key for the metadata. DO NOT share this.
//    secret_key: the unique key for the secret you created. This is key that you can share.
//    ttl: The time-to-live that was specified (i.e. not the time remaining)
//    metadata_ttl: The remaining time (in seconds) that the metadata has left to live.
//    secret_ttl: The remaining time (in seconds) that the secret has left to live.
//    recipient: if a recipient was specified, this is an obfuscated version of the email address.
//    created: Time the metadata was created in unix time (UTC)
//    updated: ditto, but the time it was last updated.
//    received: Time the secret was received.
//    passphrase_required: If a passphrase was provided when the secret was created, this will be true. Otherwise false, obviously.
type RetrieveRecentMetadataResponse []struct {
	CustID             string `json:"custid"`
	MetadataKey        string `json:"metadata_key"`
	SecretKey          string `json:"secret_key"`
	TTL                int `json:"ttl"`
	MetadataTTL        int `json:"metadata_ttl"`
	SecretTTL          int `json:"secret_ttl"`
	Recipient          []string `json:"recipient"`
	CreatedAt          int `json:"created"`
	UpdatedAt          int `json:"updated"`
	Received           int `json:"received"`
	PassphraseRequired bool   `json:"passphrase_required"`
}

// Unmarshal will read a json formatted http response body and apply those fields to structure fields
//
// Variables:
//     httpResponseBody (io.ReadCloser): The interface returned from an http.Client.Do action
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (R *RetrieveRecentMetadataResponse) Unmarshal(httpResponseBody io.ReadCloser) error {
	b, err := ioutil.ReadAll(httpResponseBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, R)
}

// CreateSecret will create a secret using the https://onetimesecret.com service
//
// Variables:
//     request (*CreateSecretRequest): A pointer to a CreateSecretRequest struct
//
// Returns:
//     (*CreateSecretResponse): A pointer to the response struct that is generated, nil if an error occurred
//     (error):                 An error if one exists, nil otherwise
func (C *Client) CreateSecret(request *CreateSecretRequest) (*CreateSecretResponse, error) {
	var (
		u        string
		err      error
		resp     = new(CreateSecretResponse)
		httpReq  *http.Request
		httpResp *http.Response
	)

	if err := request.Validate(); err != nil {
		return nil, err
	}

	u = fmt.Sprintf("https://onetimesecret.com/api/v1/share?secret=%s", request.Secret)

	if request.Passphrase != "" {
		u = fmt.Sprintf("%s,passphrase=%s", request.Passphrase)
	}
	if request.TTL != 0 {
		u = fmt.Sprintf("%s,ttl=%d", request.TTL)
	}
	if request.Recipient != nil {
		u = fmt.Sprintf("%s,recipient=%s", strings.Join(request.Recipient, ","))
	}

	httpReq, err = http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(C.creds.Username, C.creds.APIToken)

	httpResp, err = C.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service returned a non-200 status code: %d", httpResp.StatusCode)
	}

	if err := resp.Unmarshal(httpResp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

// GenerateSecret will generate a secret using the https://onetimesecret.com service
//
// Variables:
//     request (*GenerateSecretRequest): A pointer to a GenerateSecretRequest struct
//
// Returns:
//     (*GenerateSecretResponse): A pointer to the response struct that is generated, nil if an error occurred
//     (error):                   An error if one exists, nil otherwise
func (C *Client) GenerateSecret(request *GenerateSecretRequest) (*GenerateSecretResponse, error) {
	var (
		params []string
		u      string
		err      error
		resp     = new(GenerateSecretResponse)
		httpReq  *http.Request
		httpResp *http.Response
	)

	if err := request.Validate(); err != nil {
		return nil, err
	}

	u = fmt.Sprintf("https://onetimesecret.com/api/v1/generate")

	if request.Passphrase != "" || request.TTL != 0 || request.Recipient != nil {
		if request.Passphrase != "" {
			params = append(params, fmt.Sprintf("passphrase=%s", request.Passphrase))
		}
		if request.TTL != 0 {
			params = append(params, fmt.Sprintf("ttl=%d", request.TTL))
		}
		if request.Recipient != nil {
			params = append(params, fmt.Sprintf("recipient=%s", strings.Join(request.Recipient, ",")))
		}
		u = fmt.Sprintf("%s?%s", u, strings.Join(params, ","))
	}


	httpReq, err = http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(C.creds.Username, C.creds.APIToken)

	httpResp, err = C.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service returned a non-200 status code: %d", httpResp.StatusCode)
	}

	if err := resp.Unmarshal(httpResp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

// RetrieveSecret will retrieve a secret using the https://onetimesecret.com service
//
// Variables:
//     request (*RetrieveSecretRequest): A pointer to a RetrieveSecretRequest struct
//
// Returns:
//     (*RetrieveSecretResponse): A pointer to the response struct that is generated, nil if an error occurred
//     (error):                   An error if one exists, nil otherwise
func (C *Client) RetrieveSecret(request *RetrieveSecretRequest) (*RetrieveSecretResponse, error) {
	var (
		u      string
		err      error
		resp     = new(RetrieveSecretResponse)
		httpReq  *http.Request
		httpResp *http.Response
	)

	if err := request.Validate(); err != nil {
		return nil, err
	}

	u = fmt.Sprintf("https://onetimesecret.com/api/v1/secret/%s", request.SecretKey)

	if request.Passphrase != "" {
		u = fmt.Sprintf("%s?passphrase=%s", u, request.Passphrase)
	}

	httpReq, err = http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(C.creds.Username, C.creds.APIToken)

	httpResp, err = C.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service returned a non-200 status code: %d", httpResp.StatusCode)
	}

	if err := resp.Unmarshal(httpResp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

// RetrieveMetadata will retrieve metadata for a secret using the https://onetimesecret.com service
//
// Variables:
//     request (*RetrieveMetadataRequest): A pointer to a RetrieveMetadataRequest struct
//
// Returns:
//     (*RetrieveMetadataResponse): A pointer to the response struct that is generated, nil if an error occurred
//     (error):                     An error if one exists, nil otherwise
func (C *Client) RetrieveMetadata(request *RetrieveMetadataRequest) (*RetrieveMetadataResponse, error) {
	var (
		u      string
		err      error
		resp     = new(RetrieveMetadataResponse)
		httpReq  *http.Request
		httpResp *http.Response
	)

	if err := request.Validate(); err != nil {
		return nil, err
	}

	u = fmt.Sprintf("https://onetimesecret.com/api/v1/private/%s", request.MetadataKey)

	httpReq, err = http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(C.creds.Username, C.creds.APIToken)

	httpResp, err = C.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service returned a non-200 status code: %d", httpResp.StatusCode)
	}

	if err := resp.Unmarshal(httpResp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

// BurnSecret will destroy a secret using the https://onetimesecret.com service
//
// Variables:
//     request (*BurnSecretRequest): A pointer to a BurnSecretRequest struct
//
// Returns:
//     (*BurnSecretResponse): A pointer to the response struct that is generated, nil if an error occurred
//     (error):               An error if one exists, nil otherwise
func (C *Client) BurnSecret(request *BurnSecretRequest) (*BurnSecretResponse, error) {
	var (
		u      string
		err      error
		resp     = new(BurnSecretResponse)
		httpReq  *http.Request
		httpResp *http.Response
	)

	if err := request.Validate(); err != nil {
		return nil, err
	}

	u = fmt.Sprintf("https://onetimesecret.com/api/v1/private/%s/burn", request.MetadataKey)

	httpReq, err = http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(C.creds.Username, C.creds.APIToken)

	httpResp, err = C.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service returned a non-200 status code: %d", httpResp.StatusCode)
	}

	if err := resp.Unmarshal(httpResp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

// RetrieveRecentMetadata will retrieve all recent metadata using the https://onetimesecret.com service
//
// Variables:
//     request (*RetrieveRecentMetadataRequest): A pointer to a RetrieveRecentMetadataRequest struct
//
// Returns:
//     (*RetrieveRecentMetadataResponse): A pointer to the response struct that is generated, nil if an error occurred
//     (error):                           An error if one exists, nil otherwise
func (C *Client) RetrieveRecentMetadata(request *RetrieveRecentMetadataRequest) (*RetrieveRecentMetadataResponse, error) {
	var (
		url      string
		err      error
		resp     = new(RetrieveRecentMetadataResponse)
		httpReq  *http.Request
		httpResp *http.Response
	)

	url = fmt.Sprintf("https://onetimesecret.com/api/v1/private/recent")

	httpReq, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(C.creds.Username, C.creds.APIToken)

	httpResp, err = C.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service returned a non-200 status code: %d", httpResp.StatusCode)
	}

	if err := resp.Unmarshal(httpResp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}
