package privacy

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

// KeyPair contains hex-encoded public and private keys
type KeyPair struct {
	PrivateKey string
	PublicKey  string
}

// GenerateECDSAKeys generates a new ECDSA P256 key pair
func GenerateECDSAKeys() (*KeyPair, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	// Concatenate X and Y of public key
	pub := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)

	return &KeyPair{
		PrivateKey: hex.EncodeToString(priv.D.Bytes()),
		PublicKey:  hex.EncodeToString(pub),
	}, nil
}
