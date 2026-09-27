package forecast

import "encoding/json"

// GetRaw performs an authenticated GET against the Forecast API and returns the
// undecoded JSON body. path is relative to the API URL and may carry a query string
// (for example "assignments?start_date=2026-01-01&end_date=2026-03-31"). Callers that
// cache Forecast data use it to persist every attribute, including ones this library
// has no field for.
func (api *API) GetRaw(path string) (json.RawMessage, error) {
	return get[json.RawMessage](api, path)
}

// AssignmentsRaw retrieves assignments matching filter as undecoded JSON. The body has
// the shape {"assignments": [...]}.
func (api *API) AssignmentsRaw(filter AssignmentFilter) (json.RawMessage, error) {
	return api.GetRaw("assignments" + ToParams(filter.Values()))
}
