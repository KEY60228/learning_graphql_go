package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"gql/graph/model"
)

type Repository struct {
	DB *mongo.Client
}

func NewRepository(db *mongo.Client) (*Repository, error) {
	if err := db.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return &Repository{DB: db}, nil
}

func (r *Repository) PostPhoto(
	id int64,
	url string,
	name string,
	description string,
	category string,
	postedByUserID string,
	taggedUserIDs []string,
	now time.Time,
) error {
	photoColl := r.DB.Database("graphql").Collection("photos")
	photoDoc := bson.D{
		{"ID", id},
		{"URL", url},
		{"Name", name},
		{"Description", description},
		{"Category", category},
		{"PostedByUserID", postedByUserID},
		{"Created", now},
	}
	_, err := photoColl.InsertOne(context.TODO(), photoDoc)
	if err != nil {
		return err
	}

	tagColl := r.DB.Database("graphql").Collection("tags")
	tagDocs := make([]interface{}, len(taggedUserIDs))
	for i, taggedUserID := range taggedUserIDs {
		tagDocs[i] = bson.D{{"PhotoID", id}, {"UserID", taggedUserID}}
	}
	_, err = tagColl.InsertMany(context.TODO(), tagDocs)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) TotalPhotos() int {
	coll := r.DB.Database("graphql").Collection("photos")
	count, _ := coll.CountDocuments(context.TODO(), bson.D{})
	return int(count)
}

func (r *Repository) AllPhotos() []*model.Photo {
	photoColl := r.DB.Database("graphql").Collection("photos")
	cursor, err := photoColl.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var photoRes []struct {
		ID             int64
		URL            string
		Name           string
		Description    string
		Category       string
		PostedByUserID string
		Created        time.Time
	}
	cursor.All(context.TODO(), &photoRes)

	photos := make([]*model.Photo, len(photoRes))

	userColl := r.DB.Database("graphql").Collection("users")
	tagColl := r.DB.Database("graphql").Collection("tags")
	for i, r := range photoRes {
		var userRes struct {
			GithubLogin string
			Name        string
			AvatarUrl   string
		}
		userColl.FindOne(context.TODO(), bson.D{{"GithubLogin", r.PostedByUserID}}).Decode(&userRes)
		postedByUser, _ := model.NewUser(userRes.GithubLogin, userRes.Name, userRes.AvatarUrl)

		var tagRes []struct {
			PhotoID int64
			UserID  string
		}
		cursor, err := tagColl.Find(context.TODO(), bson.D{{"PhotoID", r.ID}})
		if err != nil {
			log.Fatal(err)
		}
		cursor.All(context.TODO(), &tagRes)

		taggedUsers := make([]*model.User, len(tagRes))
		for j, tag := range tagRes {
			var userRes struct {
				GithubLogin string
				Name        string
				AvatarUrl   string
			}
			userColl.FindOne(context.TODO(), bson.D{{"GithubLogin", tag.UserID}}).Decode(&userRes)
			taggedUsers[j], _ = model.NewUser(userRes.GithubLogin, userRes.Name, userRes.AvatarUrl)
		}

		photos[i], _ = model.NewPhoto(photoRes[i].ID, photoRes[i].URL, photoRes[i].Name, photoRes[i].Description, photoRes[i].Category, postedByUser, taggedUsers, photoRes[i].Created)
	}

	return photos
}

func (r *Repository) UserByID(id string) *model.User {
	var user *model.User
	coll := r.DB.Database("graphql").Collection("users")
	coll.FindOne(context.TODO(), bson.D{{"ID", id}}).Decode(&user)
	return user
}

func (r *Repository) UsersByIDs(ids []string) []*model.User {
	users := make([]*model.User, len(ids))
	for _, id := range ids {
		users = append(users, r.UserByID(id))
	}
	return users
}
