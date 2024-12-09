package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	SortOptions []SortOption
	Limit       int
}

type SortOption struct {
	Key   string
	Value int64
}

func (o Options) iterSortOptions() bson.D {
	var sortItems bson.D
	if len(o.SortOptions) > 0 {
		for _, sortOption := range o.SortOptions {
			sortItems = append(sortItems, bson.E{
				Key:   sortOption.Key,
				Value: sortOption.Value,
			})
		}
	}

	return sortItems
}

func (o Options) ManyEntryOptions() *options.FindOptions {
	opts := options.Find()
	sortItems := o.iterSortOptions()
	if len(sortItems) > 0 {
		opts.SetSort(sortItems)
	}

	if o.Limit > 0 {
		opts.SetLimit(int64(o.Limit))
	}

	return opts
}

func (o Options) OneEntryOptions() *options.FindOneOptions {
	opts := options.FindOne()
	sortItems := o.iterSortOptions()
	if len(sortItems) > 0 {
		opts.SetSort(sortItems)
	}

	return opts
}
