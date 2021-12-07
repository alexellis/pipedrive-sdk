package pipedrivesdk

type CreateOrgResponse struct {
	Success bool `json:"success"`
	Data    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type CreatePersonResponse struct {
	Success bool `json:"success"`
	Data    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type SearchResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Items []struct {
			ResultScore float64 `json:"result_score"`
			Item        struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"item"`
		} `json:"items"`
	} `json:"data"`
}

type SearchPersonResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Items []struct {
			ResultScore float64 `json:"result_score"`
			Item        struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"item"`
		} `json:"items"`
	} `json:"data"`
}
