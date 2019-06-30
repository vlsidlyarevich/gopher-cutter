package server

func (s *internal.Server) routes() {
	s.Router.HandleFunc("/cut", s.CutURLEndpoint).Methods("GET")
	s.Router.HandleFunc("/{shortURL}", s.RedirectURLEndpoint).Methods("GET")
}
