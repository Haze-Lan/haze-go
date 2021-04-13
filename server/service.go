package server

func Run() {
	haze := NewServer()
	haze.handleSignals()
	if err := haze.Start(); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
