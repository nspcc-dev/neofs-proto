package signature

import (
	"crypto/ecdsa"
	"io"

	crypto "github.com/nspcc-dev/neofs-crypto"
)

type DataSource interface {
	WriteSignedDataTo(w io.Writer) (int, error)
}

type DataWithSignature interface {
	DataSource
	GetSignatureWithKey() (key, sig []byte)
	SetSignatureWithKey(key, sig []byte)
}

type SignOption func(*Options)

type KeySignatureHandler func(key []byte, sig []byte)

type KeySignatureSource func() (key, sig []byte)

func DataSignature(key *ecdsa.PrivateKey, src DataSource, opts ...SignOption) ([]byte, error) {
	if key == nil {
		return nil, crypto.ErrEmptyPrivateKey
	}

	cfg := DefaultOptions()

	for i := range opts {
		opts[i](cfg)
	}

	return cfg.SignFunc(key, src)
}

func SignDataWithHandler(key *ecdsa.PrivateKey, src DataSource, handler KeySignatureHandler, opts ...SignOption) error {
	sig, err := DataSignature(key, src, opts...)
	if err != nil {
		return err
	}

	handler(crypto.MarshalPublicKey(&key.PublicKey), sig)

	return nil
}

func VerifyData(dataSrc DataSource, pub []byte, sig []byte, opts ...SignOption) error {
	cfg := DefaultOptions()

	for i := range opts {
		opts[i](cfg)
	}

	return cfg.VerifyFunc(
		cfg.UnmarshalPublic(pub),
		dataSrc,
		sig,
	)
}
