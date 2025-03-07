// This file was auto-generated by Fern from our API Definition.

package zep

type CreateDocumentCollectionRequest struct {
	Description *string                `json:"description,omitempty" url:"-"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" url:"-"`
}

type GetDocumentListRequest struct {
	DocumentIDs []string `json:"document_ids,omitempty" url:"-"`
	UUIDs       []string `json:"uuids,omitempty" url:"-"`
}

type DocumentSearchPayload struct {
	// Limit the number of returned documents
	Limit *int `json:"-" url:"limit,omitempty"`
	// Document metadata to filter on.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"-"`
	MinScore *float64               `json:"min_score,omitempty" url:"-"`
	// The lambda parameter for the MMR Reranking Algorithm.
	MmrLambda *float64 `json:"mmr_lambda,omitempty" url:"-"`
	// The type of search to perform. Defaults to "similarity". Must be one of "similarity" or "mmr".
	SearchType *SearchType `json:"search_type,omitempty" url:"-"`
	// The search text.
	Text *string `json:"text,omitempty" url:"-"`
}

type UpdateDocumentCollectionRequest struct {
	Description *string                `json:"description,omitempty" url:"-"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" url:"-"`
}

type UpdateDocumentRequest struct {
	DocumentID *string                `json:"document_id,omitempty" url:"-"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"-"`
}
