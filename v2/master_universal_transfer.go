/*******************************************************************************
** @Author:					Alessandro Maioli <Nicemine>
** @Email:					alessandro.maioli@gmail.com
** @Date:					Wednesday 29 September 2021 - 18:51:45
** @Filename:				user_universal_transfer.go
**
** @Last modified by:
*******************************************************************************/

package binance

import (
	"context"
	"encoding/json"
)

// CreateMasterUniversalTransferService submits a transfer request.
//
// See https://binance-docs.github.io/apidocs/spot/en/#user-universal-transfer-user_data
type CreateMasterUniversalTransferService struct {
	c               *Client
	fromEmail       *string
	toEmail         *string
	fromAccountType string
	toAccountType   string
	clientTranId    *string
	asset           string
	amount          float64
}

// Asset sets the FromEmail parameter
func (s *CreateMasterUniversalTransferService) FromEmail(
	v string,
) *CreateMasterUniversalTransferService {
	s.fromEmail = &v
	return s
}

// Asset sets the ToEmail parameter
func (s *CreateMasterUniversalTransferService) ToEmail(
	v string,
) *CreateMasterUniversalTransferService {
	s.toEmail = &v
	return s
}

// Asset sets the FromAccountType parameter (MANDATORY).
func (s *CreateMasterUniversalTransferService) FromAccountType(
	v string,
) *CreateMasterUniversalTransferService {
	s.fromAccountType = v
	return s
}

// Asset sets the ToAccountType parameter (MANDATORY).
func (s *CreateMasterUniversalTransferService) ToAccountType(
	v string,
) *CreateMasterUniversalTransferService {
	s.toAccountType = v
	return s
}

// Asset sets the ClientTranId parameter
func (s *CreateMasterUniversalTransferService) ClientTranId(
	v string,
) *CreateMasterUniversalTransferService {
	s.clientTranId = &v
	return s
}

// Asset sets the Asset parameter (MANDATORY).
func (s *CreateMasterUniversalTransferService) Asset(
	v string,
) *CreateMasterUniversalTransferService {
	s.asset = v
	return s
}

// Amount sets the Amount parameter (MANDATORY).
func (s *CreateMasterUniversalTransferService) Amount(
	v float64,
) *CreateMasterUniversalTransferService {
	s.amount = v
	return s
}

// Do sends the request.
func (s *CreateMasterUniversalTransferService) Do(
	ctx context.Context,
) (*CreateMasterUniversalTransferResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}

	if v := s.fromEmail; v != nil {
		r.setParam("fromEmail", *v)
	}
	if v := s.toEmail; v != nil {
		r.setParam("toEmail", *v)
	}
	r.setParam("fromAccountType", s.fromAccountType)
	r.setParam("toAccountType", s.toAccountType)
	if v := s.clientTranId; v != nil {
		r.setParam("clientTranId", *v)
	}
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateMasterUniversalTransferResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateMasterUniversalTransferResponse represents a response from CreateMasterUniversalTransferResponse.
type CreateMasterUniversalTransferResponse struct {
	ID           int64 `json:"tranId"`
	ClientTranID int64 `json:"clientTranId"`
}
