package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/r08521610/personal_backend/graph/model"
	"github.com/r08521610/personal_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	id := mongo.GetInstance().InsertOne("ImagesHolder", "projects", bson.M{
		"name": input.Name,
	})
	project := model.Project{
		ID:   id,
		Name: input.Name,
	}
	return &project, nil
}

func (r *mutationResolver) CreateImage(ctx context.Context, input model.NewImage) (*model.Image, error) {
	// file := model.File{}
	// fileBsonBytes, _ := bson.Marshal(mongo.GetInstance().FindOne("ImagesHolder", "files", bson.M{
	// 	"_id": input.FileID,
	// }))
	// bson.Unmarshal(fileBsonBytes, &file)
	fileID, _ := primitive.ObjectIDFromHex(input.FileID)
	file := mongo.GetInstance().FindOne("ImagesHolder", "files", bson.M{
		"_id": fileID,
	}, &model.File{}).(*model.File)
	fmt.Println(file)

	// project := model.Project{}
	// projectBsonBytes, _ := bson.Marshal(mongo.GetInstance().FindOne("ImagesHolder", "projects", bson.M{
	// 	"_id": input.ProjectID,
	// }))
	// bson.Unmarshal(projectBsonBytes, &project)
	projectID, _ := primitive.ObjectIDFromHex(input.ProjectID)
	project := mongo.GetInstance().FindOne("ImagesHolder", "projects", bson.M{
		"_id": projectID,
	}, &model.Project{}).(*model.Project)

	id := mongo.GetInstance().InsertOne("ImagesHolder", "images", bson.M{
		"file": &mongo.DBRef{
			Ref: "files",
			ID: fileID,
			DB: "ImagesHolder",
		},
		"project": &mongo.DBRef{
			Ref: "projects",
			ID: projectID,
			DB: "ImagesHolder",
		},
	})
	image := model.Image{
		ID: id,
		Project: project,
		File: file,
	}
	return &image, nil
}

func (r *queryResolver) Images(ctx context.Context) ([]*model.Image, error) {
	results := mongo.GetInstance().Find("ImagesHolder", "images", bson.M{}, &mongo.Image{})
	images := make([]*model.Image, len(results))
	for i := range results {
		result := *results[i].(*mongo.Image)
		project := mongo.GetInstance().FindOne(result.Project.DB, result.Project.Ref, bson.M{
			"_id": result.Project.ID,
		}, &model.Project{}).(*model.Project)
		file := mongo.GetInstance().FindOne(result.File.DB, result.File.Ref, bson.M{
			"_id": result.File.ID,
		}, &model.File{}).(*model.File)
		images[i] = &model.Image{
			ID: result.ID,
			Project: project,
			File: file,
		}
	}
	return images, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	results := mongo.GetInstance().Find("ImagesHolder", "projects", bson.M{}, &model.Project{})
	projects := make([]*model.Project, len(results))
	for i := range results {
		projects[i] = results[i].(*model.Project)
	}
	return projects, nil
}
