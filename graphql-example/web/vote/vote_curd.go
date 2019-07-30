package vote

import (
	"Gopher-By-Example/graphql-example/model"
	"Gopher-By-Example/graphql-example/pkg/database"
	"time"
)

func CreateVote(params CreateVoteParams) (model.VoteSerializer, error) {
	var vote model.Vote
	var result model.VoteSerializer
	tx := database.Engine.NewSession()
	tx.Begin()

	for _, i := range params.Options {
		var one model.Option
		one = model.Option{
			Name: i.Name,
		}
		if _, dbError := tx.InsertOne(&one); dbError != nil {
			tx.Rollback()
			return result, dbError
		}
		vote.OptionIds = append(vote.OptionIds, one.Id)
	}
	if params.Deadline == "" {
		now := time.Now()
		vote.Deadline = now.AddDate(0, 0, 1)
	}
	if params.Description != "" {
		vote.Description = params.Description
	}
	vote.Title = params.Title
	vote.Class = params.Class
	if _, dbError := tx.InsertOne(&vote); dbError != nil {
		tx.Rollback()
		return result, dbError
	}
	tx.Commit()
	result = vote.Serializer()
	return result, nil

}

func GetOneVote(id int64) (model.VoteSerializer, error) {
	var vote model.Vote
	var result model.VoteSerializer
	if has, dbError := database.Engine.ID(id).Get(&vote); !has || dbError != nil {
		return result, dbError
	}
	result = vote.Serializer()
	return result, nil
}
