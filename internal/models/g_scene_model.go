package models

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GSceneModel = (*customGSceneModel)(nil)

type (
	// GSceneModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGSceneModel.
	GSceneModel interface {
		gSceneModel
		withSession(session sqlx.Session) GSceneModel
		FindListByGameIdAndParentId(ctx context.Context, gameId, parentId int64) ([]*GScene, error)
	}

	customGSceneModel struct {
		*defaultGSceneModel
	}
)

// NewGSceneModel returns a model for the database table.
func NewGSceneModel(conn sqlx.SqlConn) GSceneModel {
	return &customGSceneModel{
		defaultGSceneModel: newGSceneModel(conn),
	}
}

func (m *customGSceneModel) withSession(session sqlx.Session) GSceneModel {
	return NewGSceneModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customGSceneModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*GScene, error) {

	builder = builder.Columns(gSceneRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*GScene

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customGSceneModel) FindListByGameIdAndParentId(ctx context.Context, gameId, parentId int64) ([]*GScene, error) {
	sb := squirrel.Select().From(m.table).Where("game_id = ?", gameId).
		Where("parent_id = ?", parentId)
	return m.FindAll(ctx, sb, "id asc")
}
