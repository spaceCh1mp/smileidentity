package smileidentity

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// Request model
type RequestBody struct {
	PartnerID      string `json:"partner_id"`
	Timestamp      int64  `json:"timestamp"`
	SecKey         string `json:"sec_key"`
	Country        string `json:"country"`
	IdType         string `json:"id_type"`
	IdNumber       string `json:"id_number"`
	*PartnerParams `json:"partner_params"`
}

type PartnerParams struct {
	JobId   string `json:"job_id"`
	UserId  string `json:"user_id"`
	JobType int    `json:"job_type"`
}

// Generic Response model
type Response struct {
	ResultCode    string `json:"ResultCode"`
	ResultText    string `json:"ResultText"`
	FullName      string `json:"FullName"`
	DOB           string `json:"DOB"` // format "2000-07-15"
	PhoneNumber   string `json:"PhoneNumber"`
	*BusinessData `json:"FullData"`
}

type BusinessData struct {
	TIN string `json:"TIN"`
	RC  string `json:"CAC_Reg_No"`
}

func (sc *Client) GetBVNDetails(bvn string) {
	sc.GetDetailsWithID(bvn, NIGERIA, BVN)
}

func (sc *Client) GetTINDetails(tin string) {
	sc.GetDetailsWithID(tin, NIGERIA, TIN)
}

func (sc *Client) GetCACDetails(cac string) {
	sc.GetDetailsWithID(cac, NIGERIA, CAC)
}

func (sc *Client) GetDetailsWithID(idNumber string, country countryCode, idType smileIDType) *Response {

	timestamp := time.Now().UTC().UnixNano()
	secretKey, err := calculateSecretKey(sc.config.smileAPIKey, sc.config.smilePartnerID, timestamp)
	if err != nil {
		return nil
	}

	reqBytes, err := translateParametersToRequestBody(sc.config.smilePartnerID, secretKey, idNumber, country, idType, timestamp)
	if err != nil {
		return nil
	}

	req, err := http.NewRequest("POST", getVerificationURL(sc.config.prod), bytes.NewReader(reqBytes))
	if err != nil {
		return nil
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	return &Response{}
}

func getVerificationURL(prod bool) string {
	if prod {
		return string(SMILEAPIURLIDVERIFICATIONPROD)
	}
	return string(SMILEAPIURLIDVERIFICATIONTEST)
}

func translateParametersToRequestBody(partnerID, secretKey, idNumber string, country countryCode, idType smileIDType, timestamp int64) ([]byte, error) {
	return json.Marshal(&RequestBody{
		PartnerID: partnerID,
		Timestamp: timestamp,
		SecKey:    secretKey,
		Country:   string(country),
		IdType:    string(idType),
		IdNumber:  idNumber,
		PartnerParams: &PartnerParams{
			JobId:   "",
			UserId:  "",
			JobType: 5,
		},
	})
}
