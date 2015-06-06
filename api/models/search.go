package models

import ()

type SearchResult struct {
	Type     string      `json:"type"`
	Resource interface{} `json:"resource"`
}

type searchService struct {
	Service
}

func NewSearchService(params map[string]interface{}) *searchService {
	s := new(searchService)
	s.params = params
	return s
}

func (this *searchService) Search() (*[]SearchResult, error) {
	var (
		collection []SearchResult = make([]SearchResult, 0)
		err        error
	)
	if _, ok := this.params["query"]; ok {
		confService := NewConferenceService(this.params)
		if conferences, err := confService.Conferences(); err == nil {
			for _, item := range *conferences {
				searchResult := SearchResult{
					Type:     "conference",
					Resource: item,
				}
				collection = append(collection, searchResult)
			}
		}

		eventService := NewEventService(this.params)
		if events, err := eventService.Events(); err == nil {
			for _, item := range *events {
				searchResult := SearchResult{
					Type:     "event",
					Resource: item,
				}
				collection = append(collection, searchResult)
			}
		}

		videoService := NewVideoService(this.params)
		if videos, err := videoService.Videos(); err == nil {
			for _, item := range *videos {
				searchResult := SearchResult{
					Type:     "video",
					Resource: item,
				}
				collection = append(collection, searchResult)
			}
		}
	}
	return &collection, err
}
