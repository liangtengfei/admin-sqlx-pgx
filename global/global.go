package global

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/config"
	"study.com/demo-sqlx-pgx/pkg/token"
)

var (
	Config     config.Config
	Log        *zap.Logger
	TokenMaker token.Maker
	Enforcer   *casbin.Enforcer
)
