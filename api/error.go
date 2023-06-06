package api

type InfoError struct {
    Data  any    `json:"data,omitempty"`
    Error string `json:"error"`
}
