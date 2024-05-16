package ws

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
)

type Server struct {
	Topics map[string]*Topic
}

func NewServer() *Server {
	return &Server{
		Topics: make(map[string]*Topic),
	}
}

func (s *Server) NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws/sub/{topic}", s.HandleSubscribe)
	mux.HandleFunc("/ws/pub/{topic}", s.HandlePublish)
	return mux
}

func (s *Server) HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	topicName, err := GetTopicName(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t := s.GetOrCreateTopic(topicName)
	t.ServeHTTP(w, r)
}

func (s *Server) HandlePublish(w http.ResponseWriter, r *http.Request) {
	topicName, err := GetTopicName(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t := s.GetOrCreateTopic(topicName)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body:", err)
		return
	}
	t.Forward <- body
}

func GetTopicName(r *http.Request) (string, error) {
	v := r.PathValue("topic")
	if v == "" {
		return "", errors.New("missing topic in path")
	}
	return v, nil
}

func (s *Server) GetOrCreateTopic(t string) *Topic {
	if _, ok := s.Topics[t]; !ok {
		s.Topics[t] = NewTopic()
	}
	return s.Topics[t]
}
