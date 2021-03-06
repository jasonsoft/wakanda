package identity

import (
	"context"
	"time"
)

type Account struct {
	ID                    string     `json:"id,omitempty" db:"id"`
	App                   string     `json:"app,omitempty" db:"app"`
	Username              string     `json:"username,omitempty" db:"username"`
	PasswordHash          string     `json:"-" db:"password_hash"`
	FirstName             string     `json:"first_name" db:"first_name"`
	LastName              string     `json:"last_name" db:"last_name"`
	Avatar                string     `json:"avatar,omitempty" db:"avatar"`
	Email                 string     `json:"email,omitempty" db:"email"`
	Mobile                string     `json:"mobile,omitempty" db:"mobile"`
	ExternalID            string     `json:"external_id,omitempty" db:"external_id"`
	IsLockedOut           bool       `json:"is_locked_out,omitempty" db:"is_locked_out"`
	FailedPasswordAttempt int        `json:"failed_password_attempt_count,omitempty" db:"failed_password_attempt_count"`
	Roles                 []string   `json:"roles,omitempty"`
	ClientIP              string     `json:"client_ip,omitempty" db:"client_ip"`
	UserAgent             string     `json:"user_agent,omitempty" db:"user_agent"`
	Extension             string     `json:"extension,omitempty" db:"extension"`
	LastLoginAt           *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
	CreatedAt             *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt             *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type FindAccountOptions struct {
	ID               string `json:"id" db:"id"`
	ExternalID       string `json:"external_id" db:"external_id"`
	App              string `json:"app" db:"app"`
	Username         string `json:"username" db:"username"`
	Email            string `json:"email,omitempty" db:"email"`
	Mobile           string `json:"mobile,omitempty" db:"mobile"`
	Role             string `json:"role"`
	IsLockedOut      int    `json:"is_locked_out" db:"is_locked_out"`
	Skip             int    `db:"skip"`
	Take             int    `db:"take"`
	SortBy           string `db:"sortby"`
	Sort             string
	CreatedTimeStart *time.Time `db:"created_start_time"`
	CreatedTimeEnd   *time.Time `db:"created_end_time"`
	LoginTimeStart   *time.Time `db:"login_start_time"`
	LoginTimeEnd     *time.Time `db:"login_end_time"`
}

type AccountServicer interface {
	Account(ctx context.Context, accountID string) (*Account, error)
	Accounts(ctx context.Context, opt *FindAccountOptions) ([]*Account, error)
	AccountCount(ctx context.Context, opt *FindAccountOptions) (int, error)
	CreateAccount(ctx context.Context, account *Account) error
	UpdateAccountPassword(ctx context.Context, accountID string, newPassword string) error
	LockAccount(ctx context.Context, app, accountID string) error
	UnlockAccount(ctx context.Context, app, accountID string) error
}
