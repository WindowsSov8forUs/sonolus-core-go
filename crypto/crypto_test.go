package crypto_test

import (
	"testing"

	sonocrypto "github.com/WindowsSov8forUs/sonolus-core-go/crypto"
)

func TestHash(t *testing.T) {
	const want = "a9993e364706816aba3e25717850c26c9cd0d89d"

	if got := sonocrypto.Hash([]byte("abc")); got != want {
		t.Fatalf("Hash() = %q, want %q", got, want)
	}
}

func TestSignaturePublicKey(t *testing.T) {
	key, err := sonocrypto.SignaturePublicKey()
	if err != nil {
		t.Fatal(err)
	}

	if key.Curve == nil || key.X == nil || key.Y == nil {
		t.Fatalf("SignaturePublicKey() returned incomplete key: %#v", key)
	}
	if !key.Curve.IsOnCurve(key.X, key.Y) {
		t.Fatal("SignaturePublicKey() returned a point outside the curve")
	}
}
