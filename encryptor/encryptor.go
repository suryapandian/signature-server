package encryptor

// An abstraction to interact with encryption methods like RSA, ed25519
type Encryptor interface {
	GetPublicKey() []byte
	Sign([]byte) []byte
}
