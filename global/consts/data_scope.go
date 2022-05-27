package consts

const ScopeDataKey = "scope_req"

// Scope 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
type Scope int

const (
	ScopeDataAll Scope = iota + 1
	ScopeCustom
	ScopeDept
	ScopeDeptChild
	ScopeSelf
)

func (s Scope) String() string {
	switch s {
	case ScopeDataAll:
		return "1"
	case ScopeCustom:
		return "2"
	case ScopeDept:
		return "3"
	case ScopeDeptChild:
		return "4"
	case ScopeSelf:
		return "5"
	}
	return "-1"
}

func (s Scope) Label() string {
	switch s {
	case ScopeDataAll:
		return "全部数据权限"
	case ScopeCustom:
		return "自定数据权限"
	case ScopeDept:
		return "本部门数据权限"
	case ScopeDeptChild:
		return "本部门及以下数据权限"
	case ScopeSelf:
		return "仅本人权限"
	}
	return "未知类型"
}

func (s Scope) Value() int {
	switch s {
	case ScopeDataAll:
		return 1
	case ScopeCustom:
		return 2
	case ScopeDept:
		return 3
	case ScopeDeptChild:
		return 4
	case ScopeSelf:
		return 5
	}
	return -1
}
