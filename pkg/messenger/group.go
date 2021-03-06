package messenger

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type GroupType int

const (
	GroupTypeSystem       GroupType = 1
	GroupTypeP2P          GroupType = 2
	GroupTypePrivateGroup GroupType = 3
)

type GroupState int

const (
	GroupStateNormal    GroupState = 1
	GroupStateDissolved GroupState = 2
)

type Group struct {
	ID             string     `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	Description    string     `json:"description" db:"description"`
	Type           GroupType  `json:"type" db:"type"`
	MaxMemberCount int        `json:"-" db:"max_member_count"`
	MemberCount    int        `json:"member_count" db:"member_count"`
	CreatorID      string     `json:"-" db:"creator_id"`
	State          GroupState `json:"state" db:"state"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
}

type GroupMember struct {
	ID        string     `json:"id" db:"id"`
	GroupID   string     `json:"group_id" db:"group_id"`
	MemberID  string     `json:"member_id" db:"member_id"`
	IsAdmin   bool       `json:"is_admin" db:"is_admin"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type FindGroupOptions struct {
	IDs             string `db:"group_ids"`
	Type            GroupType
	MemberID        string     `db:"member_id"`
	AnchorUpdatedAt *time.Time `db:"anchor_updated_at"`
	Skip            int        `db:"skip"`
	PerPage         int        `db:"per_page"`
}

type FindGroupMemberOptions struct {
	GroupID         string     `db:"group_id"`
	AnchorUpdatedAt *time.Time `db:"anchor_updated_at"`
	Skip            int        `db:"skip"`
	PerPage         int        `db:"per_page"`
}

type GroupServicer interface {
	Groups(ctx context.Context, opts *FindGroupOptions) ([]*Group, error)
	CreateGroup(ctx context.Context, group *Group, memberIDs []string) error
	DissolveGroup(ctx context.Context, groupID string) error // 解散群組
	JoinGroup(ctx context.Context, groupID string, memberID string) error
	LeaveGroup(ctx context.Context, groupID string) error
	AddGroupMember(ctx context.Context, groupID string, memberID string) error
	GroupMembers(ctx context.Context, opts *FindGroupMemberOptions) ([]*GroupMember, error)
	SetAdmin(ctx context.Context, groupID string, memberID string) error
	RemoveAdmin(ctx context.Context, groupID string, memberID string) error
	GroupAdmins(ctx context.Context, groupID string) ([]*GroupMember, error)
}

type GroupRepository interface {
	DB() *sqlx.DB
	Groups(ctx context.Context, opts *FindGroupOptions) ([]*Group, error)
	InsertTx(ctx context.Context, target *Group, tx *sqlx.Tx) error
	IsMemberInGroup(ctx context.Context, memberID, groupID string) bool
}

type GroupMemberRepository interface {
	DB() *sqlx.DB
	GroupMembers(ctx context.Context, opts *FindGroupMemberOptions) ([]*GroupMember, error)
	BatchInsertTx(ctx context.Context, members []*GroupMember, tx *sqlx.Tx) error
}
