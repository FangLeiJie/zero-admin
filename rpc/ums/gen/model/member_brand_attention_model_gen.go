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

type memberBrandAttentionModel interface {
	Insert(ctx context.Context, data *MemberBrandAttention) error
	FindOne(ctx context.Context, id string) (*MemberBrandAttention, error)
	Update(ctx context.Context, data *MemberBrandAttention) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultMemberBrandAttentionModel struct {
	conn *mon.Model
}

func newDefaultMemberBrandAttentionModel(conn *mon.Model) *defaultMemberBrandAttentionModel {
	return &defaultMemberBrandAttentionModel{conn: conn}
}

func (m *defaultMemberBrandAttentionModel) Insert(ctx context.Context, data *MemberBrandAttention) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultMemberBrandAttentionModel) FindOne(ctx context.Context, id string) (*MemberBrandAttention, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data MemberBrandAttention

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

func (m *defaultMemberBrandAttentionModel) Update(ctx context.Context, data *MemberBrandAttention) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultMemberBrandAttentionModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
