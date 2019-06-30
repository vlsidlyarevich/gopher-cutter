package server

import (
	"github.com/gorilla/mux"
	"github.com/vlsidlyarevich/gopher-cutter/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Server struct {
	Db     *mongo.Database
	Router *mux.Router
	Dao    *internal.DAO
}

func NewServer(db *mongo.Database) (s *Server) {
	router := mux.NewRouter()
	server := Server{db, router, internal.NewLinkDAO(db)}
	server.routes()
	return &server
}

func (s *Server) CutURLEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	url := r.URL.Query().Get("url")
	if url == "" {
		log.Panicf("Cannot cut empty string")
	}
	var result *internal.Link
	result = s.Dao.FindByURL(url)
	if result == nil {
		var randId, e = internal.NewRandom().RInt(10)
		if e != nil {
			log.Panic(e)
		}

		shortUrl := internal.Encode(randId)

		result = s.Dao.Save(internal.Link{
			ShortURL: shortUrl,
			URL:      url,
		})
	}
	_, _ = w.Write([]byte(result.ShortURL))
}

func (s *Server) RedirectURLEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)
	short, _ := params["shortURL"]

	savedLink := s.Dao.FindByShortURL(short)
	if savedLink == nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Undefined url"))
		return
	}

	w.Header().Add("Location", savedLink.URL)
	w.WriteHeader(301)
}
