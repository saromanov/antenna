package server

import structs "github.com/saromanov/antenna/structs/v1"

func mapAggregateResponse(resp *structs.AggregateSearchResponse) *AggregateResponse {
	return &AggregateResponse{
		Count: resp.Count,
	}
}
