package server

import (
	"github.com/gorilla/mux"
	"github.com/vlsidlyarevich/gopher-cutter/internal/app/gopher-cutter/link"
	"github.com/vlsidlyarevich/gopher-cutter/internal/app/gopher-cutter/util"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Server struct {
	Db     *mongo.Database
	Router *mux.Router
	Dao    *link.DAO
}

func NewServer(db *mongo.Database) (s *Server) {
	router := mux.NewRouter()
	server := Server{db, router, link.NewLinkDAO(db)}
	server.routes()
	return &server
}

func (s *Server) CutURLEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	url := r.URL.Query().Get("url")
	if url == "" {
		log.Panicf("Cannot cut empty string")
	}
	var result *link.Link
	result = s.Dao.FindByURL(url)
	if result == nil {
		var randId, e = util.NewRandom().RInt(10)
		if e != nil {
			log.Panic(e)
		}

		shortUrl := util.Encode(randId)

		result = s.Dao.Save(link.Link{
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
