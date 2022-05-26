package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/utils"
	"time"
)

func userCreateSQL(req request.UserCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameUser).
		Columns("real_name", "email", "dept_id", "user_name", "remark", "create_time", "create_by", "status", "password", "posts", "avatar", "sex", "mobile").
		Values(req.RealName, req.Email, req.DeptID, req.UserName, req.Remark, time.Now(), username, "0", req.Password, utils.Int64Join(req.PostIds), req.Avatar, req.Sex, req.Mobile).
		Suffix("RETURNING \"id\"").ToSql()
}

func (store *SQLStore) UserCreate(ctx context.Context, req request.UserCreateRequest, username string) (int64, error) {
	sql, args, err := userCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	// 事务执行
	tx, err := store.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	// 执行操作
	var id int64
	err = tx.QueryRowx(sql, args...).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	// 关联角色
	if len(req.RoleIds) > 0 {
		err = userRoleRelate(tx, id, req.RoleIds)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()

	return id, err
}

func userUpdateSQL(req request.UserUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameUser)
	if req.RealName != "" {
		sql = sql.Set("real_name", req.RealName)
	}
	if req.Mobile != "" {
		sql = sql.Set("mobile", req.Mobile)
	}
	if req.Email != "" {
		sql = sql.Set("email", req.Email)
	}
	if req.Sex != "" {
		sql = sql.Set("sex", req.Sex)
	}
	if req.Avatar != "" {
		sql = sql.Set("avatar", req.Avatar)
	}
	if req.Status != "" {
		sql = sql.Set("status", req.Status)
	}
	if len(req.PostIds) > 0 {
		sql = sql.Set("posts", utils.Int64Join(req.PostIds))
	}
	if req.DeptID > 0 {
		sql = sql.Set("dept_id", req.DeptID)
	}
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) UserUpdate(ctx context.Context, req request.UserUpdateRequest, username string) (int64, error) {
	sql, args, err := userUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	// 事务执行
	tx, err := store.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	// 执行操作
	var rows int64
	result, err := tx.Exec(sql, args...)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	rows, err = result.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	// 关联角色
	if len(req.RoleIds) > 0 {
		err = userRoleRelate(tx, req.ID, req.RoleIds)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()

	return rows, err
}

func userRoleRelate(tx *sqlx.Tx, id int64, relateIds []int64) error {
	return relateDataReset(tx, id, relateIds, "user_id", "role_id", TBNameUserRole)
}

func userFindByUsernameSQL(username string) (string, []interface{}, error) {
	return baseQuerySQLBuilder(TBNameUser).Where(sq.Eq{"user_name": username}).Limit(1).ToSql()
}

func (store *SQLStore) UserFindByUsername(username string) (AgoUser, error) {
	var result AgoUser

	sql, args, err := userFindByUsernameSQL(username)
	if err != nil {
		return result, err
	}

	err = store.db.Get(&result, sql, args...)
	return result, err
}

func userFindByMobileSQL(mobile string) (string, []interface{}, error) {
	return baseQuerySQLBuilder(TBNameUser).Where(sq.Eq{"mobile": mobile}).Limit(1).ToSql()
}

func (store *SQLStore) UserFindByMobile(mobile string) (AgoUser, error) {
	var result AgoUser

	sql, args, err := userFindByMobileSQL(mobile)
	if err != nil {
		return result, err
	}

	err = store.db.Get(&result, sql, args...)
	return result, err
}

func (store *SQLStore) UserCountByMobile(mobile string) (int64, error) {
	var total int64

	sql, args, err := SQLBuilder().Select("count(*)").From(TBNameUser).Where(sq.Eq{"mobile": mobile}).ToSql()
	if err != nil {
		return 0, err
	}

	err = store.db.Get(&total, sql, args...)

	return total, err
}

func userFindByIdSQL(id int64) (string, []interface{}, error) {
	return baseQuerySQLBuilder(TBNameUser).Where(sq.Eq{"id": id}).Limit(1).ToSql()
}

func (store *SQLStore) UserFindById(id int64) (AgoUser, error) {
	var result AgoUser
	sql, args, err := userFindByIdSQL(id)
	if err != nil {
		return result, err
	}

	err = store.db.Get(&result, sql, args...)
	return result, err
}

func userPageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameUser).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"user_name": fmt.Sprint("%", req.Keyword, "%")},
			sq.Like{"real_name": fmt.Sprint("%", req.Keyword, "%")},
			sq.Like{"mobile": fmt.Sprint("%", req.Keyword, "%")},
		})
	}

	// 此处截取COUNT SQL
	countSQL, _, err = sql.ToSql()
	if err != nil {
		return
	}
	countSQL = SQLCount(countSQL)

	sql = sql.Offset(req.GetOffset()).Limit(req.GetLimit())
	sql = sql.OrderBy("create_time DESC")

	querySQL, args, err = sql.ToSql()

	return querySQL, countSQL, args, err
}

func (store *SQLStore) UserPageAndKeyword(ctx context.Context, req request.PaginationRequest) (int64, []AgoUser, error) {
	var result []AgoUser
	var total int64

	fail := func(err error) (int64, []AgoUser, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := userPageAndKeywordSQL(req)
	if err != nil {
		return fail(err)
	}

	err = store.db.GetContext(ctx, &total, countSQL, args...)
	if err != nil {
		return fail(err)
	}

	err = store.db.SelectContext(ctx, &result, pageSQL, args...)
	if err != nil {
		return fail(err)
	}

	return total, result, nil
}

func (store *SQLStore) UserDetail(req request.ByIdRequest) (AgoUser, error) {
	var result AgoUser

	sql, args, err := DetailSQLBuilder(TBNameUser, req.Id)
	if err != nil {
		return result, err
	}

	err = store.db.Get(&result, sql, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (store *SQLStore) UserDeleteId(id int64, username string) error {
	sql, args, err := DeleteFakeSQLBuilder(TBNameUser, id, username)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(sql, args...)
	return err
}
