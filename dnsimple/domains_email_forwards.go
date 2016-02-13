package dnsimple

import (
	"fmt"
)

// EmailForwardResponse represents a response from an API method that returns an EmailForward struct.
type EmailForwardResponse struct {
	Response
	Data *EmailForward `json:"data"`
}

// EmailForwardsResponse represents a response from an API method that returns a collection of EmailForward struct.
type EmailForwardsResponse struct {
	Response
	Data []EmailForward `json:"data"`
}

// EmailForward represents an email forward in DNSimple.
type EmailForward struct {
	ID        int    `json:"id,omitempty"`
	DomainID  int    `json:"domain_id,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func emailForwardPath(accountID string, domain interface{}, forwardID int) string {
	path := fmt.Sprintf("%v/email_forwards", domainPath(accountID, domain))

	if forwardID != 0 {
		path += fmt.Sprintf("/%d", forwardID)
	}

	return path
}

// ListEmailForwards lists the email forwards for a domain.
//
// See https://developer.dnsimple.com/v2/domains/email-forwards/#list
func (s *DomainsService) ListEmailForwards(accountID string, domain interface{}) (*EmailForwardsResponse, error) {
	path := versioned(emailForwardPath(accountID, domain, 0))
	forwardsResponse := &EmailForwardsResponse{}

	resp, err := s.client.get(path, forwardsResponse)
	if err != nil {
		return nil, err
	}

	forwardsResponse.HttpResponse = resp
	return forwardsResponse, nil
}

// CreateEmailForward creates a new email forward.
//
// See https://developer.dnsimple.com/v2/domains/email-forwards/#create
func (s *DomainsService) CreateEmailForward(accountID string, domain interface{}, forwardAttributes EmailForward) (*EmailForwardResponse, error) {
	path := versioned(emailForwardPath(accountID, domain, 0))
	forwardResponse := &EmailForwardResponse{}

	resp, err := s.client.post(path, forwardAttributes, forwardResponse)
	if err != nil {
		return nil, err
	}

	forwardResponse.HttpResponse = resp
	return forwardResponse, nil
}

// GetEmailForward fetches an email forward.
//
// See https://developer.dnsimple.com/v2/domains/email-forwards/#get
func (s *DomainsService) GetEmailForward(accountID string, domain interface{}, forwardID int) (*EmailForwardResponse, error) {
	path := versioned(emailForwardPath(accountID, domain, forwardID))
	forwardResponse := &EmailForwardResponse{}

	resp, err := s.client.get(path, forwardResponse)
	if err != nil {
		return nil, err
	}

	forwardResponse.HttpResponse = resp
	return forwardResponse, nil
}

// DeleteEmailForward PERMANENTLY deletes an email forward from the domain.
//
// See https://developer.dnsimple.com/v2/domains/email-forwards/#delete
func (s *DomainsService) DeleteEmailForward(accountID string, domain interface{}, forwardID int) (*EmailForwardResponse, error) {
	path := versioned(emailForwardPath(accountID, domain, forwardID))
	forwardResponse := &EmailForwardResponse{}

	resp, err := s.client.delete(path, nil, nil)
	if err != nil {
		return nil, err
	}

	forwardResponse.HttpResponse = resp
	return forwardResponse, nil
}
