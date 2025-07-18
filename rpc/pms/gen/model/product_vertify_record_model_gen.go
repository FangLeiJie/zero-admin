// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4

package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productVertifyRecordModel interface {
	Insert(ctx context.Context, data *ProductVertifyRecord) error
	FindOne(ctx context.Context, id string) (*ProductVertifyRecord, error)
	Update(ctx context.Context, data *ProductVertifyRecord) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultProductVertifyRecordModel struct {
	conn *mon.Model
}

func newDefaultProductVertifyRecordModel(conn *mon.Model) *defaultProductVertifyRecordModel {
	return &defaultProductVertifyRecordModel{conn: conn}
}

func (m *defaultProductVertifyRecordModel) Insert(ctx context.Context, data *ProductVertifyRecord) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultProductVertifyRecordModel) FindOne(ctx context.Context, id string) (*ProductVertifyRecord, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data ProductVertifyRecord

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductVertifyRecordModel) Update(ctx context.Context, data *ProductVertifyRecord) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultProductVertifyRecordModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
