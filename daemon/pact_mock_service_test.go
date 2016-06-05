package daemon

import "testing"

func TestPactMockService_NewService(t *testing.T) {
	s := &PactMockService{}
	port, svc := s.NewService([]string{"--foo bar"})

	if port <= 0 {
		t.Fatalf("Expected non-zero port but got: %d", port)
	}

	if svc == nil {
		t.Fatalf("Expected a non-nil object but got nil")
	}

	if s.Args[4] != "--foo bar" {
		t.Fatalf("Expected '--foo bar' argument to be passed")
	}
}
