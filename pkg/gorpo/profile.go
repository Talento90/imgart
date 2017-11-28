package gorpo

type Filter struct {
	Id     string                 `json:"id"`
	Params map[string]interface{} `json:"params"`
}

type Profile struct {
	Id          string   `json:"id"`
	Filters     []Filter `json:"array"`
	FallbackUri string   `json:"fallback_uri,omitempty"`
}

func (p *Profile) Merge(profile Profile) {

}
