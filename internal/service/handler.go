/*
Copyright © 2021, 2022 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import "github.com/gorilla/mux"

// Handler structure represents HTTP request handler.
type Handler struct {
	svc Interface
}

// NewHandler function constructs new HTTP request handler.
func NewHandler(svc Interface) *Handler {
	return &Handler{
		svc: svc,
	}
}

// Register function registers new handler for given endpoint URL.
func (s *Handler) Register(r *mux.Router) {
	r.Handle("/gathering_rules", gatheringRulesEndpoint(s.svc))
}
