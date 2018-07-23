package autosign

import (
	"fmt"
	"testing"
	"time"
)

// Test Sign
func TestSign(t *testing.T) {

	startProcess := time.Now()

	hash, created := Init()

	elapsedProcess := time.Since(startProcess)
	if created == true {
		fmt.Println("New file signed has been created!")
	}
	fmt.Printf("Generated %s in %s\n", hash, elapsedProcess)

}
