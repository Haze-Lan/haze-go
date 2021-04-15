package server

func Run() {
	haze := NewServer()
	if err := haze.Start(); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
