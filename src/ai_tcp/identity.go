package ai_tcp

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// AIIdentity represents the identity information for an LLM participant in an AI-TCP session.
type AIIdentity struct {
	UUID      string
	ModelName string
	PublicKey string
}

// GenerateAIIdentity creates a new AIIdentity for the provided model name.
// It generates a UUIDv4 and uses a fixed simulated public key.
func GenerateAIIdentity(modelName string) *AIIdentity {
	return &AIIdentity{
		UUID:      generateUUIDv4(),
		ModelName: modelName,
		PublicKey: "GPT_SIMULATED_PUBKEY",
	}
}

// SignPayload returns a simulated signature for the given payload using the identity's UUID.
// The signature is the SHA256 hash of the UUID concatenated with the payload.
func SignPayload(identity *AIIdentity, payload string) string {
	if identity == nil {
		return ""
	}

	h := sha256.New()
	h.Write([]byte(identity.UUID + payload))
	return hex.EncodeToString(h.Sum(nil))
}

// generateUUIDv4 generates a random UUIDv4 string in the form xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func generateUUIDv4() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		// In this simple implementation panic is acceptable because generation failure should never happen.
		panic(fmt.Errorf("failed to read random bytes: %w", err))
	}

	// Set version (4) and variant bits as per RFC 4122.
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4],
		b[4:6],
		b[6:8],
		b[8:10],
		b[10:16])
}
