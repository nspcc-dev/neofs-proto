package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/pkg/accounting"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	v2accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

// Accounting contains methods related to balance querying.
type Accounting interface {
	// GetBalance returns balance of provided account.
	GetBalance(context.Context, *owner.ID, ...CallOption) (*accounting.Decimal, error)
}

func (c *clientImpl) GetBalance(ctx context.Context, owner *owner.ID, opts ...CallOption) (*accounting.Decimal, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2accounting.BalanceRequestBody)
	reqBody.SetOwnerID(owner.ToV2())

	req := new(v2accounting.BalanceRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.Balance(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, errors.Wrap(err, "transport error")
	}

	err = v2signature.VerifyServiceMessage(resp, callOptions.signOpts()...)
	if err != nil {
		return nil, errors.Wrap(err, "can't verify response message")
	}

	return accounting.NewDecimalFromV2(resp.GetBody().GetBalance()), nil
}
