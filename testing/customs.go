package gen

//Code generated by mulgen. DO NOT EDIT (i mean, it will all be lost in the void)

import (
	"context"
	"github.com/Vliro/humus"
	"time"
)

var _ context.Context
var _ time.Time
var _ humus.Fields

// This is a custom field that is defined by custom.toml.
var QuestionPublishedFields humus.FieldList = []humus.Field{MakeField("Post.datePublished", 0)}

// This is a custom field that is defined by custom.toml.
var QuestionMetaFields humus.FieldList = []humus.Field{MakeField("Question.id", 0), MakeField("Post.datePublished", 0)}
