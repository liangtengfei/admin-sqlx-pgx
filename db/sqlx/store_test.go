package db

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSQLBuilderIN(t *testing.T) {
	id := []int64{1, 2, 3}
	var questions = AttachQuestions(len(id))
	subSQL := fmt.Sprintf("SELECT menu_id FROM ago_role_menu WHERE role_id = IN (%s)", questions)
	sql, args, err := SQLBuilder("pgx").Select("*").From(TBNameMenu).
		Where("id IN ("+subSQL+")", id).
		ToSql()
	t.Log("生成的SQL：", sql, args)
	require.NoError(t, err)
	require.Equal(t, len(args), len(id))
	require.NotEmpty(t, sql)

}
