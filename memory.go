// This file was auto-generated by Fern from our API Definition.

package zep

type AddMemoryRequest struct {
	// Additional instruction for generating the facts. Zep Cloud Only, will be ignored on Community Edition.
	FactInstruction *string `json:"fact_instruction,omitempty" url:"-"`
	// A list of message objects, where each message contains a role and content.
	Messages []*Message `json:"messages,omitempty" url:"-"`
	// Additional instruction for generating the summary. Zep Cloud Only, will be ignored on Community Edition.
	SummaryInstruction *string `json:"summary_instruction,omitempty" url:"-"`
}

type CreateSessionRequest struct {
	// Optional instruction to use for fact rating.
	FactRatingInstruction *FactRatingInstruction `json:"fact_rating_instruction,omitempty" url:"-"`
	// The metadata associated with the session.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"-"`
	// The unique identifier of the session.
	SessionID string `json:"session_id" url:"-"`
	// The unique identifier of the user associated with the session
	UserID string `json:"user_id" url:"-"`
}

type AddFactsRequest struct {
	Facts []*NewFact `json:"facts,omitempty" url:"-"`
}

type EndSessionRequest struct {
	Classify    *ClassifySessionRequest `json:"classify,omitempty" url:"-"`
	Instruction *string                 `json:"instruction,omitempty" url:"-"`
}

type EndSessionsRequest struct {
	Instruction *string  `json:"instruction,omitempty" url:"-"`
	SessionIDs  []string `json:"session_ids,omitempty" url:"-"`
}

type ExtractDataRequest struct {
	// Your current date and time in ISO 8601 format including timezone. This is used for determining relative dates.
	CurrentDateTime *string `json:"current_date_time,omitempty" url:"-"`
	// The number of messages in the chat history from which to extract data
	LastN int `json:"last_n" url:"-"`
	// The schema describing the data to be extracted. See Zep's SDKs for more details.
	ModelSchema string `json:"model_schema" url:"-"`
	// Validate that the extracted data is present in the dialog and correct per the field description.
	// Mitigates hallucination, but is slower and may result in false negatives.
	Validate *bool `json:"validate,omitempty" url:"-"`
}

type MemoryGetRequest struct {
	// The number of most recent memory entries to retrieve.
	Lastn *int `json:"-" url:"lastn,omitempty"`
	// The minimum rating by which to filter facts
	MinRating *float64 `json:"-" url:"minRating,omitempty"`
}

type MemoryGetSessionFactsRequest struct {
	// Minimum rating by which to filter facts (Zep Cloud only)
	MinRating *float64 `json:"-" url:"minRating,omitempty"`
}

type MemoryGetSessionMessagesRequest struct {
	// Limit the number of results returned
	Limit *int `json:"-" url:"limit,omitempty"`
	// Cursor for pagination
	Cursor *int `json:"-" url:"cursor,omitempty"`
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
	// The maximum number of search results to return. Defaults to None (no limit).
	Limit *int `json:"-" url:"limit,omitempty"`
	// Metadata Filter
	Metadata      map[string]interface{} `json:"metadata,omitempty" url:"-"`
	MinFactRating *float64               `json:"min_fact_rating,omitempty" url:"-"`
	MinScore      *float64               `json:"min_score,omitempty" url:"-"`
	MmrLambda     *float64               `json:"mmr_lambda,omitempty" url:"-"`
	SearchScope   *SearchScope           `json:"search_scope,omitempty" url:"-"`
	SearchType    *SearchType            `json:"search_type,omitempty" url:"-"`
	Text          *string                `json:"text,omitempty" url:"-"`
}

type SessionSearchQuery struct {
	// The maximum number of search results to return. Defaults to None (no limit).
	Limit *int `json:"-" url:"limit,omitempty"`
	// The minimum fact rating to filter on. Only supported on cloud. Will be ignored on Community Edition.
	MinFactRating *float64 `json:"min_fact_rating,omitempty" url:"-"`
	// The minimum score for search results. Only supported on cloud. Will be ignored on Community Edition.
	MinScore *float64 `json:"min_score,omitempty" url:"-"`
	// The lambda parameter for the MMR Reranking Algorithm. Only supported on cloud. Will be ignored on Community Edition.
	MmrLambda *float64 `json:"mmr_lambda,omitempty" url:"-"`
	// Record filter on the metadata. Only supported on cloud. Will be ignored on Community Edition.
	RecordFilter map[string]interface{} `json:"record_filter,omitempty" url:"-"`
	// Search scope. Only supported on cloud. On Community Edition the search scope is always "facts".
	SearchScope *SearchScope `json:"search_scope,omitempty" url:"-"`
	// Search type. Only supported on cloud. Will be ignored on Community Edition.
	SearchType *SearchType `json:"search_type,omitempty" url:"-"`
	// the session ids to search
	SessionIDs []string `json:"session_ids,omitempty" url:"-"`
	// The search text.
	Text string `json:"text" url:"-"`
	// User ID used to determine which sessions to search. Required on Community Edition.
	UserID *string `json:"user_id,omitempty" url:"-"`
}

type MemorySynthesizeQuestionRequest struct {
	// The number of messages to use for question synthesis.
	LastNMessages *int `json:"-" url:"lastNMessages,omitempty"`
}

type ModelsMessageMetadataUpdate struct {
	// The metadata to update
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"-"`
}

type UpdateSessionRequest struct {
	// Optional instruction to use for fact rating.
	// Fact rating instructions can not be unset.
	FactRatingInstruction *FactRatingInstruction `json:"fact_rating_instruction,omitempty" url:"-"`
	// The metadata to update
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"-"`
}
