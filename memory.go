// This file was auto-generated by Fern from our API Definition.

package zep

import (
	fmt "fmt"
)

type CreateSessionRequest struct {
	Metadata  map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	SessionID string                 `json:"session_id" url:"session_id"`
	// Must be a pointer to allow for null values
	UserID *string `json:"user_id,omitempty" url:"user_id,omitempty"`
}

type ClassifySessionRequest struct {
	Classes     []string `json:"classes,omitempty" url:"classes,omitempty"`
	Instruction *string  `json:"instruction,omitempty" url:"instruction,omitempty"`
	LastN       *int     `json:"last_n,omitempty" url:"last_n,omitempty"`
	Name        string   `json:"name" url:"name"`
	Persist     *bool    `json:"persist,omitempty" url:"persist,omitempty"`
}

type ModelsExtractDataRequest struct {
	LastNMessages  *int                  `json:"last_n_messages,omitempty" url:"last_n_messages,omitempty"`
	ZepDataClasses []*ModelsZepDataClass `json:"zep_data_classes,omitempty" url:"zep_data_classes,omitempty"`
}

type MemoryGetRequest struct {
	// memoryType: perpetual or message_window
	MemoryType *MemoryGetRequestMemoryType `json:"-" url:"memoryType,omitempty"`
	// Last N messages. Overrides memory_window configuration
	Lastn *int `json:"-" url:"lastn,omitempty"`
}

type MemoryListSessionsRequest struct {
	// Page number for pagination, starting from 1
	PageNumber *int `json:"-" url:"page_number,omitempty"`
	// Number of sessions to retrieve per page
	PageSize *int `json:"-" url:"page_size,omitempty"`
	// Field to order the results by: created_at, updated_at, user_id, session_id
	OrderBy *string `json:"-" url:"order_by,omitempty"`
	// Order direction: true for ascending, false for descending
	Asc *bool `json:"-" url:"asc,omitempty"`
}

type MemorySearchPayload struct {
	// Limit the number of results returned
	Limit *int `json:"-" url:"limit,omitempty"`
	// Metadata Filter
	Metadata    map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	MinScore    *float64               `json:"min_score,omitempty" url:"min_score,omitempty"`
	MmrLambda   *float64               `json:"mmr_lambda,omitempty" url:"mmr_lambda,omitempty"`
	SearchScope *SearchScope           `json:"search_scope,omitempty" url:"search_scope,omitempty"`
	SearchType  *SearchType            `json:"search_type,omitempty" url:"search_type,omitempty"`
	Text        *string                `json:"text,omitempty" url:"text,omitempty"`
}

type MemorySynthesizeQuestionRequest struct {
	// Last N messages
	LastNMessages *int `json:"-" url:"lastNMessages,omitempty"`
}

type MemoryGetRequestMemoryType string

const (
	MemoryGetRequestMemoryTypePerpetual        MemoryGetRequestMemoryType = "perpetual"
	MemoryGetRequestMemoryTypeSummaryRetriever MemoryGetRequestMemoryType = "summary_retriever"
	MemoryGetRequestMemoryTypeMessageWindow    MemoryGetRequestMemoryType = "message_window"
)

func NewMemoryGetRequestMemoryTypeFromString(s string) (MemoryGetRequestMemoryType, error) {
	switch s {
	case "perpetual":
		return MemoryGetRequestMemoryTypePerpetual, nil
	case "summary_retriever":
		return MemoryGetRequestMemoryTypeSummaryRetriever, nil
	case "message_window":
		return MemoryGetRequestMemoryTypeMessageWindow, nil
	}
	var t MemoryGetRequestMemoryType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (m MemoryGetRequestMemoryType) Ptr() *MemoryGetRequestMemoryType {
	return &m
}

type ModelsMessageMetadataUpdate struct {
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
}

type UpdateSessionRequest struct {
	// The metadata to update
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
}
