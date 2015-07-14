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
		tagService := NewTagService(this.params)
		if tags, err := tagService.Tags(); err == nil {
			for _, item := range *tags {
				searchResults := SearchResult{
					Type:     "tag",
					Resource: item,
				}
				collection = append(collection, searchResults)
			}
		}

		confService := NewConferenceService(this.params)
		if conferences, err := confService.Conferences(); err == nil {
			for _, item := range *conferences {
				searchResults := SearchResult{
					Type:     "conference",
					Resource: item,
				}
				collection = append(collection, searchResults)
			}
		}

		eventService := NewEventService(this.params)
		if events, err := eventService.Events(); err == nil {
			for _, item := range *events {
				searchResults := SearchResult{
					Type:     "event",
					Resource: item,
				}
				collection = append(collection, searchResults)
			}
		}

		videoService := NewVideoService(this.params)
		if videos, err := videoService.Videos(); err == nil {
			for _, item := range *videos {
				searchResults := SearchResult{
					Type:     "video",
					Resource: item,
				}
				collection = append(collection, searchResults)
			}
		}
	}
	return &collection, err
}
