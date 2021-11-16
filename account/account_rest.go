package account

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/google/uuid"
	h "github.com/rosfran/interview-accountapi/http"
)

const AccountURI = "/v1/organisation/accounts"

const VERSION = 0

type AccountRequest struct {
	ID             string
	OrganisationID string
	RESTClient     h.RESTInterface
}

func NewAccountRequest(baseURL url.URL, organisationID string) *AccountRequest {
	return &AccountRequest{
		ID:             uuid.NewString(),
		OrganisationID: organisationID,
		RESTClient:     &h.RESTClient{BaseURL: &baseURL},
	}
}

func (c *AccountRequest) Create(name, orgID string, attributes map[string]string) (a *Account, err error) {

	var names []string = []string{name}
	attr := &AccountAttributesField{
		Name: names,
	}

	var country, foundCountry = attributes["Country"]
	if foundCountry {
		attr.Country = &country
	}

	var baseCurrency, foundBaseCurrency = attributes["BaseCurrency"]
	if foundBaseCurrency {
		attr.BaseCurrency = baseCurrency

	}

	var bankID, foundBankID = attributes["BankID"]
	if foundBankID {
		attr.BankID = bankID

	}

	var bankIDCode, foundBankIDCode = attributes["BankIDCode"]
	if foundBankIDCode {
		attr.BankIDCode = bankIDCode

	}

	var bic, foundBic = attributes["Bic"]
	if foundBic {
		attr.Bic = bic

	}

	var iban, foundIban = attributes["Iban"]
	if foundIban {
		attr.Iban = iban

	}

	account := &AccountDataField{
		Attributes:     attr,
		OrganisationID: orgID,
		ID:             orgID,
	}

	jsonReq, err := c.marshalAccount(&Account{
		Data: *account,
	})

	if err != nil {
		return nil, err
	}

	jsonRes, err := c.RESTClient.Post(jsonReq)

	if err != nil {
		if accountErr, jsonErr := c.unmarshalError(err.Error()); jsonErr != nil {
			return nil, jsonErr
		} else {
			return nil, errors.New(accountErr)
		}
	}

	accountDataResponse, err := c.unmarshalAccount(jsonRes)

	if err != nil {
		return nil, err
	}

	return &accountDataResponse, nil

}

func (c *AccountRequest) Fetch(ID string) (a *AccountArray, err error) {

	jsonRes, err := c.RESTClient.Get(ID)

	if err != nil {
		if accountErr, jsonErr := c.unmarshalError(err.Error()); jsonErr != nil {
			return nil, jsonErr
		} else {
			return nil, errors.New(accountErr)
		}
	}

	accountDataResponse, err := c.unmarshalAccountArray(jsonRes)

	if err != nil {
		return nil, err
	}

	return &accountDataResponse, nil

}

func (c *AccountRequest) Delete(ID string) (a *Account, err error) {

	jsonRes, err := c.RESTClient.Delete(ID, "0")

	if err != nil {
		if accountErr, jsonErr := c.unmarshalError(err.Error()); jsonErr != nil {
			return nil, jsonErr
		} else {
			return nil, errors.New(accountErr)
		}
	}

	accountDataResponse, err := c.unmarshalAccount(jsonRes)

	if err != nil {
		return nil, err
	}

	return &accountDataResponse, nil

}

func (c *AccountRequest) marshalAccount(account *Account) (string, error) {

	request := &Account{
		AccountDataField{

			Type:           "accounts",
			ID:             c.ID,
			OrganisationID: c.OrganisationID,
			Attributes:     account.Data.Attributes,
			Version:        func() *int64 { i := int64(VERSION); return &i }(),
		},
	}

	requestString, err := json.Marshal(request)

	return string(requestString), err
}

func (c *AccountRequest) unmarshalAccountArray(jsonData string) (accountData AccountArray, err error) {

	err = json.Unmarshal([]byte(jsonData), &accountData)
	return accountData, err
}

func (c *AccountRequest) unmarshalAccount(jsonData string) (accountData Account, err error) {
	err = json.Unmarshal([]byte(jsonData), &accountData)
	return accountData, err
}

func (c *AccountRequest) unmarshalError(jsonData string) (accountErr string, err error) {
	err = json.Unmarshal([]byte(jsonData), &accountErr)
	return accountErr, err
}
