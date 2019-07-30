package payment_repository

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/introphin/proto-payment-repository"
	"google.golang.org/grpc"
)

type MockCreateWalletPaymentRepository struct {
	*MockPaymentRepository
}

func (*MockCreateWalletPaymentRepository) CreateWallet(ctx context.Context, in *payment_repository.CreateWalletRequest, opts ...grpc.CallOption) (*payment_repository.Wallet, error) {
	return &payment_repository.Wallet{
		Id:              1,
		CreatedAt:       ptypes.TimestampNow(),
		UpdatedAt:       ptypes.TimestampNow(),
		UserId:          in.UserId,
		PaymentSystemId: in.PaymentSystemId,
		Name:            in.Name,
		Currency:        in.Currency,
		AutoPayment:     in.AutoPayment,
		AccountNumber:   in.AccountNumber,
		Checked:         in.Checked,
		EditedUserId:    in.EditedUserId,
		Type:            in.Type,
	}, nil
}
