package dao

import "github.com/jmoiron/sqlx"

type Options struct {
	Tx         *sqlx.Tx
	FromMaster bool
}

type Option func(*Options)

func WithTx(tx *sqlx.Tx) Option {
	return func(o *Options) {
		o.Tx = tx
	}
}

func WithMaster() Option {
	return func(o *Options) {
		o.FromMaster = true
	}
}
