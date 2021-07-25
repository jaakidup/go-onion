package service

import "net/http"

type Service struct {
	Name string
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
