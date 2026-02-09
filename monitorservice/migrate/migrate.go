package migrate

import (
	"database/sql"
	"embed"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgerrcode"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
)

const (
	dbErrorDatabaseDoesNotExist = pgerrcode.InvalidCatalogName
	migrationTable              = "migration_info"
	dialect                     = "postgres"
	defaultLimit                = -1
)

//go:embed sql/*
var sqlMigrateFS embed.FS

type statusRow struct {
	ID        string
	Migrated  bool
	Unknown   bool
	AppliedAt time.Time
}

type migrationService struct {
	dbAddress  string
	limit      int
	migrations *migrate.EmbedFileSystemMigrationSource
	db         *sql.DB
}

// 確認初始化資料是否完整
func StartUpCheck(logger *zap.Logger, db *sql.DB) {

	// startup check
	migrate.SetTable(migrationTable)
	migrate.SetIgnoreUnknown(true)

	m := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: sqlMigrateFS,
		Root:       "sql",
	}

	migrations, err := m.FindMigrations()
	if err != nil {
		logger.Fatal("Could not find migrations", zap.Error(err))
	}
	records, err := migrate.GetMigrationRecords(db, dialect)
	if err != nil {
		logger.Fatal("Could not get migration records, run `order service migrate up`", zap.Error(err))
	}

	diff := len(migrations) - len(records)
	if diff > 0 {
		// migrate up
		ms := &migrationService{
			migrations: m,
			db:         db,
		}

		ms.up(logger)
	}
	if diff < 0 {
		// update migrate file
		logger.Warn("DB schema newer, update order service", zap.Int64("migrations", int64(math.Abs(float64(diff)))))
	}
}

func Parse(args []string, tmpLogger *zap.Logger, db *sql.DB) {
	tmp := strings.Split(args[0], " ")
	if len(tmp) == 0 {
		tmpLogger.Fatal("Migrate requires a subcommand. Available commands are: 'up', 'down', 'redo', 'status'.")
	}

	command := ""
	limit := 1

	for i, v := range tmp {
		if v == "migrate" {
			// 命令
			if len(tmp) > i+1 {
				command = tmp[i+1]

				// 次數
				if len(tmp) > i+2 {
					l, err := strconv.Atoi(tmp[i+2])
					if err != nil {
						limit = -1
					} else {
						limit = l
					}
				}
				break
			}

		}
	}

	switch command {
	case "up":
	case "down":
	case "redo":
	case "status":
	}

	log.Println(command)
	log.Println(limit)
}

func (ms *migrationService) up(logger *zap.Logger) {
	if ms.limit < defaultLimit {
		ms.limit = 0
	}
	// appliedMigrations, err := migrate.ExecMax(ms.db, dialect, ms.migrations, migrate.Down, ms.limit) //all
	appliedMigrations, err := migrate.Exec(ms.db, dialect, ms.migrations, migrate.Up)
	if err != nil {
		logger.Fatal("Failed to apply migrations", zap.Int("count", appliedMigrations), zap.Error(err))
	}

	logger.Info("Successfully applied migration", zap.Int("count", appliedMigrations))
}

func (ms *migrationService) down(logger *zap.Logger) {
	if ms.limit < defaultLimit {
		ms.limit = 1
	}

	appliedMigrations, err := migrate.ExecMax(ms.db, dialect, ms.migrations, migrate.Down, ms.limit)
	if err != nil {
		logger.Fatal("Failed to migrate back", zap.Int("count", appliedMigrations), zap.Error(err))
	}

	logger.Info("Successfully migrated back", zap.Int("count", appliedMigrations))
}

func (ms *migrationService) redo(logger *zap.Logger) {
	if ms.limit > defaultLimit {
		logger.Warn("Limit is ignored when redo is invoked")
	}

	appliedMigrations, err := migrate.ExecMax(ms.db, dialect, ms.migrations, migrate.Down, 1)
	if err != nil {
		logger.Fatal("Failed to migrate back", zap.Int("count", appliedMigrations), zap.Error(err))
	}
	logger.Info("Successfully migrated back", zap.Int("count", appliedMigrations))

	appliedMigrations, err = migrate.ExecMax(ms.db, dialect, ms.migrations, migrate.Up, 1)
	if err != nil {
		logger.Fatal("Failed to apply migrations", zap.Int("count", appliedMigrations), zap.Error(err))
	}
	logger.Info("Successfully applied migration", zap.Int("count", appliedMigrations))
}

func (ms *migrationService) status(logger *zap.Logger) {
	if ms.limit > defaultLimit {
		logger.Warn("Limit is ignored when status is invoked")
	}

	migrations, err := ms.migrations.FindMigrations()
	if err != nil {
		logger.Fatal("Could not find migrations", zap.Error(err))
	}

	records, err := migrate.GetMigrationRecords(ms.db, dialect)
	if err != nil {
		logger.Fatal("Could not get migration records", zap.Error(err))
	}

	rows := make(map[string]*statusRow)

	for _, m := range migrations {
		rows[m.Id] = &statusRow{
			ID:       m.Id,
			Migrated: false,
		}
	}

	unknownMigrations := make([]string, 0)
	for _, r := range records {
		sr, ok := rows[r.Id]
		if !ok {
			// Unknown migration found in database, perhaps from a newer server version.
			unknownMigrations = append(unknownMigrations, r.Id)
			continue
		}
		sr.Migrated = true
		sr.AppliedAt = r.AppliedAt
	}

	for _, m := range migrations {
		if rows[m.Id].Migrated {
			logger.Info(m.Id, zap.String("applied", rows[m.Id].AppliedAt.Format(time.RFC822Z)))
		} else {
			logger.Info(m.Id, zap.String("applied", ""))
		}
	}
	for _, m := range unknownMigrations {
		logger.Warn(m, zap.String("applied", "unknown migration, check if database is set up for a newer server version"))
	}
}
