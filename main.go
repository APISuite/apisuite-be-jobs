package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

type Config struct {
	DatabaseURI           string
	ActivationCleanerCron string
	ActivationCleanerTTL  string
	PasswordRecoveryCron  string
	PasswordRecoveryTTL   string
}

type cleanerConfig struct {
	db *sql.DB
	ttl string
}

func main() {
	c := setupConfig()

	execute(c)
}

func execute(config *Config) {
	db, err := openDB(config.DatabaseURI)
	if err != nil {
		log.Fatal("Could not init db connection: ", err)
	}

	c := cron.New()
	c.AddFunc(config.ActivationCleanerCron, setupActivationCleaner(&cleanerConfig{
		db: db,
		ttl: config.ActivationCleanerTTL,
	}))
	c.AddFunc(config.PasswordRecoveryCron, setupRecoveryCleaner(&cleanerConfig{
		db: db,
		ttl: config.PasswordRecoveryTTL,
	}))
	c.Start()

	select {} // block it
}

func setupConfig() *Config {
	return &Config{
		DatabaseURI:           getEnv("APISUITE_JOBS_DB", "postgres://apisuite:m00se@localhost:5432/apisuite?sslmode=disable"),
		ActivationCleanerCron: getEnv("APISUITE_JOBS_ACTV_CRON", "*/30 * * * *"),
		ActivationCleanerTTL:  getEnv("APISUITE_JOBS_ACTV_TTL", "12"),
		PasswordRecoveryCron:  getEnv("APISUITE_JOBS_RECOV_CRON", "*/30 * * * *"),
		PasswordRecoveryTTL:   getEnv("APISUITE_JOBS_RECOV_TTL", "2"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func openDB(databaseURI string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURI)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func setupActivationCleaner(conf *cleanerConfig) func() {
	return func() {
		log.Println("[ACTIVATION] starting job")

		stmt := `DELETE FROM users 
			WHERE activation_token IS NOT NULL
			AND created_at + interval '1h' * $1 < now()`
		res, err := conf.db.Exec(stmt, conf.ttl)
		if err != nil {
			log.Println("[ACTIVATION] error: ", err)
		}

		count, err := res.RowsAffected()
		if err != nil {
			log.Println("[ACTIVATION] error: ", err)
		}
		log.Println("[ACTIVATION] deleted rows: ", count)

		log.Println("[ACTIVATION] finished job")
	}
}

func setupRecoveryCleaner(conf *cleanerConfig) func() {
	return func() {
		log.Println("[RECOVERY] starting job")

		stmt := `DELETE FROM password_recovery 
			WHERE created_at + interval '1h' * $1 < now()`
		res, err := conf.db.Exec(stmt, conf.ttl)
		if err != nil {
			log.Println("[RECOVERY] error: ", err)
		}

		count, err := res.RowsAffected()
		if err != nil {
			log.Println("[RECOVERY] error: ", err)
		}
		log.Println("[RECOVERY] deleted rows: ", count)

		log.Println("[RECOVERY] finished job")
	}
}
