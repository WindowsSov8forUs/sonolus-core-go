package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"errors"
	"math/big"

	"github.com/WindowsSov8forUs/sonolus-core-go/core/server"
)

func SignaturePublicKey() (*ecdsa.PublicKey, error) {
	jwk := server.SignaturePublicKey()
	if jwk.Kty != "EC" || jwk.Crv != "P-256" {
		return nil, errors.New("signature public key must be an EC P-256 key")
	}

	x, err := decodeBase64URLInteger(jwk.X)
	if err != nil {
		return nil, err
	}
	y, err := decodeBase64URLInteger(jwk.Y)
	if err != nil {
		return nil, err
	}

	curve := elliptic.P256()
	if !curve.IsOnCurve(x, y) {
		return nil, errors.New("signature public key point is not on P-256")
	}

	return &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}, nil
}

func decodeBase64URLInteger(value string) (*big.Int, error) {
	data, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}
