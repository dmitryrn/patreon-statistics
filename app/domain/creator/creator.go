package creator

type Creator struct {
	Id        int    `json:"id,omitempty"`
	PatreonId string `json:"patreonId,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}
