package mongo

import (
	"context"
	"time"

	"github.com/talento90/imgart/errors"
	"github.com/talento90/imgart/imgart"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type profileRepository struct {
	collection string
	client     *Client
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if err == mongo.ErrNilDocument {
		return errors.ENotExists("Profile does not exists", err)
	}

	return errors.EInternal("Error occurred", err)
}

// NewProfileRepository returns a profile mongo repository
func NewProfileRepository(client *Client) imgart.ProfileRepository {
	return &profileRepository{
		collection: "profiles",
		client:     client,
	}
}

func (r *profileRepository) GetAll(limit int64, skip int64) (*[]imgart.Profile, error) {
	db := r.client.Client.Database(r.client.Database)
	col := db.Collection(r.collection)
	opt := options.Find().SetSkip(skip).SetLimit(limit)
	ctx := context.Background()

	c, err := col.Find(ctx, bson.D{{}}, opt)

	if err != nil {
		return nil, handleError(err)
	}

	profiles := make([]imgart.Profile, 0, limit)

	err = c.All(ctx, &profiles)

	return &profiles, err
}

func (r *profileRepository) Get(id string) (*imgart.Profile, error) {
	db := r.client.Client.Database(r.client.Database)
	col := db.Collection(r.collection)
	opt := options.FindOne()

	result := col.FindOne(context.Background(), bson.M{"_id": id}, opt)

	profile := &imgart.Profile{}
	err := result.Decode(profile)

	return profile, handleError(err)
}

func (r *profileRepository) Create(profile *imgart.Profile) error {
	db := r.client.Client.Database(r.client.Database)
	col := db.Collection(r.collection)

	profile.Created = time.Now().UTC()
	profile.Updated = time.Now().UTC()

	_, err := col.InsertOne(context.Background(), profile)

	return handleError(err)
}

func (r *profileRepository) Update(profile *imgart.Profile) error {
	db := r.client.Client.Database(r.client.Database)
	col := db.Collection(r.collection)

	profile.Updated = time.Now().UTC()

	_, err := col.ReplaceOne(context.Background(), bson.M{"_id": profile.ID}, profile)

	return handleError(err)
}

func (r *profileRepository) Delete(id string) error {
	db := r.client.Client.Database(r.client.Database)
	col := db.Collection(r.collection)

	_, err := col.DeleteOne(context.Background(), bson.M{"_id": id})

	return handleError(err)
}
