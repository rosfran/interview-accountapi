package account

import (
	"net/url"
	"testing"

	"github.com/google/uuid"
	test "github.com/rosfran/interview-accountapi/test_assertions"
)

// Acceptance Tests
//
// The acceptance tests are used to test the actual behavior
// of the account client libraries connecting to the APIs
// provided by the docker-compose docker images.
func TestNewAccountRequest(t *testing.T) {
	b, _ := url.Parse("http://localhost:8080" + "/v1/organisation/accounts")
	client := NewAccountRequest(*b, "634e3a41-26b8-49f9-a23d-26fa92061f38")

	test.AssertNotNil(t, client)
}

var id string

// Create registers an existing bank account with Form3
// or create a new one
//
// Creating a bank account in the UK according to the
// following constraints:
// - United Kingdom	Country code: GB
// - Bank ID: required, 6 characters, UK sort code
// - BIC: required
// - Bank ID Code: required, has to be GBDSC
// - Account Number: optional, 8 characters, generated if not provided
// - IBAN: Generated if not provided
func TestCreateAccount(t *testing.T) {
	b, _ := url.Parse("http://localhost:8080" + "/v1/organisation/accounts")

	orgid := uuid.NewString()

	client := NewAccountRequest(*b, orgid)

	account, _ := client.Create("1stAccount", "GB", "GPB", "400302", "GBDSC", "NWBKGB42", "GB28NWBK40030212764204", orgid)

	id = account.Data.ID
	test.AssertNotNil(t, account.Data.Attributes.Country)
	test.AssertNotNil(t, account.Data.Attributes.AccountNumber)
	test.AssertNotNil(t, account.Data.Attributes.Iban)

	test.AssertEqual(t, "NWBKGB42", account.Data.Attributes.Bic)
}

func TestCreateAccountError(t *testing.T) {
	b, _ := url.Parse("http://localhost:8080" + "/v1/organisation/accounts")

	orgid := uuid.NewString()

	client := NewAccountRequest(*b, orgid)

	account, _ := client.Create("1stAccount", "GB", "GPB", "400302", "GBDSC", "", "GB28NWBK40030212764204", orgid)

	id = account.Data.ID
	test.AssertNotNil(t, account.Data.Attributes.Country)
	test.AssertNotNil(t, account.Data.Attributes.AccountNumber)
	test.AssertNotNil(t, account.Data.Attributes.Iban)

	test.AssertEqual(t, "NWBKGB42", account.Data.Attributes.Bic)
}

func TestCreateAccountWithError(t *testing.T) {
	b, _ := url.Parse("http://localhost:8080" + "/v1/organisation/accounts")
	client := NewAccountRequest(*b, "GB")

	_, err := client.Create("2ndAccount", "***", "GPB", "400302", "GBDSC", "NWBKGB42", "GB28NWBK40030212764204", uuid.NewString())

	test.AssertNotNil(t, err)

}

func TestFetchAccount(t *testing.T) {
	b, _ := url.Parse("http://localhost:8080" + "/v1/organisation/accounts")

	orgid := uuid.NewString()

	if id == "" || len(id) == 0 {
		id = "d38f9bbc-3180-4cae-9ed8-91dcc2c991ae"
	}
	client := NewAccountRequest(*b, orgid)

	accountL, _ := client.Fetch(id)

	account := accountL.Data[0]
	test.AssertNotNil(t, account)
	test.AssertNotNil(t, account)
	test.AssertNotNil(t, account.ID)
	test.AssertEqual(t, id, account.ID)
	test.AssertNotNil(t, account.Attributes.Country)
	test.AssertNotNil(t, account.Attributes.AccountNumber)
	test.AssertNotNil(t, account.Attributes.Iban)

	test.AssertEqual(t, "NWBKGB42", account.Attributes.Bic)
}

func TestDeleteAccount(t *testing.T) {
	b, _ := url.Parse("http://localhost:8080" + "/v1/organisation/accounts")

	orgid := uuid.NewString()

	if id == "" || len(id) == 0 {
		id = "d38f9bbc-3180-4cae-9ed8-91dcc2c991ae"
	}
	client := NewAccountRequest(*b, orgid)

	account, _ := client.Delete(id)

	test.AssertNotNil(t, account)
	test.AssertNotNil(t, account)
}

func TestDeleteAccountNotFound(t *testing.T) {
	b, _ := url.Parse("http://localhost:8080" + "/v1/organisation/accounts")

	orgid := uuid.NewString()

	if id == "" || len(id) == 0 {
		id = "oiof9bbc-3180-4cae-9ed8-91dcc2c991ae"
	}
	client := NewAccountRequest(*b, orgid)

	account, _ := client.Delete(id)

	test.AssertNil(t, account)

}
