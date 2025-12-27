package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
}

type Genre struct {
	GenreId   int    `bson:"genre_id" json:"genre_id" validate:"required"`
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required,min=2,max=100"`
}

type Movie struct {
	ID          bson.ObjectID `bson:"_id" json:"_id" gorm:"primaryKey"`
	ImdbID      string        `bson:"imdb_id" json:"imdb_id" gorm:"uniqueIndex" validate:"required"`
	Title       string        `bson:"title" json:"title" validate:"required,min=2,max=500" `
	PosterPath  string        `bson:"poster_path" json:"poster_path" validate:"required,url"`
	YoutubeID   string        `bson:"youtube_id" json:"youtube_id" validate:"required"`
	Genre       []Genre       `bson:"genre" json:"genre" validate:"required,dive,required"`
	AdminReview string        `bson:"admin_review" json:"admin_review" validate:"required,min=10,max=5000"`
	Ranking     Ranking       `json:"ranking" gorm:"foreignKey:RankingID" validate:"required"`
}
