package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"time"
)

func operationLogCreateSQL(req request.OperationLogCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameOperationLog).
		Columns("business_type", "business_title", "invoke_method", "request_method", "request_url", "client_type", "client_ip", "client_location",
			"client_param", "operation_type", "operation_result", "error_msg", "status", "create_by", "create_time", "remark", "dept_name").
		Values(req.BusinessType, req.BusinessTitle, req.InvokeMethod, req.RequestMethod, req.RequestUrl, req.ClientType, req.ClientIp,
			req.ClientLocation, req.ClientParam, req.OperationType, req.OperationResult, req.ErrorMsg, req.Status, username,
			time.Now(), req.Remark, req.DeptName).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) OperationLogCreate(ctx context.Context, req request.OperationLogCreateRequest, username string) (int64, error) {
	sql, args, err := operationLogCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func operationLogPageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameOperationLog)
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"business_type": fmt.Sprint("%", req.Keyword, "%")},
			sq.Like{"business_title": fmt.Sprint("%", req.Keyword, "%")},
			sq.Like{"invoke_method": fmt.Sprint("%", req.Keyword, "%")},
		})
	}

	// 此处截取COUNT SQL
	countSQL, _, err = sql.ToSql()
	if err != nil {
		return
	}
	countSQL = SQLCount(countSQL)

	//分页
	sql = sql.Offset(req.GetOffset()).Limit(req.GetLimit())

	//排序
	if req.SortField != "" && req.SortOrder != "" {
		sql = sql.OrderBy(req.SortField + " " + req.SortOrder)
	} else {
		sql = sql.OrderBy("create_time DESC")
	}

	querySQL, args, err = sql.ToSql()

	return querySQL, countSQL, args, err
}

func (store *SQLStore) OperationLogPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoOperationLog, error) {
	var result []AgoOperationLog
	var total int64

	fail := func(err error) (int64, []AgoOperationLog, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := operationLogPageAndKeywordSQL(req)
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
