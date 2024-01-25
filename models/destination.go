package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Destination struct {
    ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Name        string             `json:"name,omitempty" bson:"name,omitempty"`
    Description string             `json:"description,omitempty" bson:"description,omitempty"`
    City        string             `json:"city,omitempty" bson:"city,omitempty"`
    Country     string             `json:"country,omitempty" bson:"country,omitempty"`
    Landmarks   []string           `json:"landmarks,omitempty" bson:"landmarks,omitempty"`
    Rating      float64            `json:"rating,omitempty" bson:"rating,omitempty"`
}
