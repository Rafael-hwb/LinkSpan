package main

import (
	"context"
	"log"
	"time"

	"linkspan/admin/bootstrap"
	"linkspan/admin/internal/repository"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := bootstrap.OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	fieldCipher, err := bootstrap.OpenFieldCipher()
	if err != nil {
		return err
	}

	userRepository := repository.NewUserRepository(
		db,
		fieldCipher,
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Minute,
	)
	defer cancel()

	if err := userRepository.MigrateSensitiveFields(ctx); err != nil {
		return err
	}

	log.Println("用户手机号和邮箱迁移完成")
	return nil
}
