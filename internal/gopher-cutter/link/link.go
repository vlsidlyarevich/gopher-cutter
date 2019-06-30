package link

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Link struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ShortURL string             `bson:"short_url,omitempty"`
	URL      string             `bson:"url,omitempty"`
}

type DAO struct {
	Database *mongo.Database
}

const (
	COLLECTION = "urls"
)

func NewLinkDAO(db *mongo.Database) (dao *DAO) {
	return &DAO{db}
}

func (dao *DAO) FindByURL(url string) (link *Link) {
	var result Link
	collection := dao.Database.Collection(COLLECTION)
	err := collection.FindOne(context.TODO(), Link{URL: url}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func (dao *DAO) FindByShortURL(url string) (link *Link) {
	var result Link
	collection := dao.Database.Collection(COLLECTION)
	err := collection.FindOne(context.TODO(), Link{ShortURL: url}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func (dao *DAO) Save(link Link) (id *Link) {
	////Вынести в декоратор который обернет обработку запросов и будет класть конекшн в контекст
	collection := dao.Database.Collection(COLLECTION)
	_, _ = collection.InsertOne(context.TODO(), link)
	return &link
}
