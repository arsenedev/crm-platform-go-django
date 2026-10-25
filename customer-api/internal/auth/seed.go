package auth

import (
	"context"
	"log"
	"os"

	"company/customer-api/pkg/logger"

	"company/customer-api/internal/user"
)

func SeedAdmin(
	ctx context.Context,
	userService *user.Service,
) error {
	logger.Log.Info(
		"admin_seed_started",
	)
	users, err := userService.List(
		ctx,
	)

	if err != nil {
		return err
	}

	for _, u := range users {

		if u.Email == os.Getenv("ADMIN_EMAIL") {

			logger.Log.Info(
				"admin_already_exists",
			)

			return nil
		}
	}

	err = userService.Create(
		ctx,
		user.CreateUserRequest{
			Email:     os.Getenv("ADMIN_EMAIL"),
			Password:  os.Getenv("ADMIN_PASSWORD"),
			Firstname: os.Getenv("ADMIN_FIRSTNAME"),
			Lastname:  os.Getenv("ADMIN_LASTNAME"),
			Role:      "admin",
		},
	)

	if err != nil {
		logger.Log.Info(
			"admin_already_exists",
		)
		return err
	}

	log.Println(
		"[SEED] admin created",
	)
	logger.Log.Info(
		"admin_created",
	)

	return nil
}
