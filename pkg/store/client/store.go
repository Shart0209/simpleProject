package store

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"
)

type Executor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Store interface {
	Stop() error
	GetLogger() zerolog.Logger
	GetExecutor() (Executor, error)
	GetRepository(ex Executor) Repository
}

//
//func main() {
//	cfg, _ := api.NewConfig()
//	ctx := context.Background()
//
//	conn, err := NewClient(ctx, cfg)
//
//	if err = conn.Ping(context.Background()); err != nil {
//		fmt.Println("Ping nil")
//		conn.Close()
//	}
//	fmt.Println("Ping OK")
//
//	rows, _ := conn.Query(ctx, "SELECT file_id, file_uuid FROM files WHERE file_id=$1", 1)
//	defer rows.Close()
//
//	for rows.Next() {
//		var file_id int64
//		var file_uuid uuid.UUID
//		_ = rows.Scan(&file_id, &file_uuid)
//		fmt.Printf("id: %v name: %v", file_id, file_uuid)
//	}
//}
