package go_pgp_test

import (
	"fmt"
	"github.com/kbereza/go-pgp"
	"testing"
)

func TestSignature(t *testing.T) {
	fmt.Println("Signature test: START")
	entity, err := go_pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created private key entity.")

	signature, err := go_pgp.Sign(entity, []byte(TestMessage))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created signature of test message with private key entity.")

	publicKeyEntity, err := go_pgp.GetEntity([]byte(TestPublicKey), []byte{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created public key entity.")

	err = go_pgp.Verify(publicKeyEntity, []byte(TestMessage), signature)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Signature verified using public key entity.")
	fmt.Println("Signature test: END")
}
