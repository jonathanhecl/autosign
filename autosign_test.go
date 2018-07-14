package autosign

import (
	"fmt"
	"testing"
	"time"
)

// Test Sign
func TestSign(t *testing.T) {

	startProcess := time.Now()

	hash := Init()

	elapsedProcess := time.Since(startProcess)
	fmt.Printf("Generated %s in %s\n", hash, elapsedProcess)

}
