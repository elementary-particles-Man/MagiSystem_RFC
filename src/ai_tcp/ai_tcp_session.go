package ai_tcp

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"
)

// AITCPSession represents a mock AI-TCP session between two models.
type AITCPSession struct {
	ID        string
	ModelFrom string
	ModelTo   string
	StartTime time.Time
}

// InitAITCPSession initializes a new AITCPSession with the provided model names.
func InitAITCPSession(modelFrom, modelTo string) *AITCPSession {
	// Generate a 128-bit (16-byte) random session ID.
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		log.Fatalf("failed to generate session ID: %v", err)
	}

	return &AITCPSession{
		ID:        hex.EncodeToString(buf),
		ModelFrom: modelFrom,
		ModelTo:   modelTo,
		StartTime: time.Now(),
	}
}

// SimulateConnection prints a mock handshake between the two models.
func SimulateConnection(session *AITCPSession) {
	if session == nil {
		log.Println("no session provided")
		return
	}

	log.Printf("[AI-TCP] Session %s initialized at %s", session.ID, session.StartTime.Format(time.RFC3339))
	log.Printf("[AI-TCP] %s -> %s: initiating secure handshake", session.ModelFrom, session.ModelTo)
	log.Printf("[AI-TCP] %s -> %s: handshake acknowledged", session.ModelTo, session.ModelFrom)
	log.Printf("[AI-TCP] Session %s secure channel established", session.ID)
}
