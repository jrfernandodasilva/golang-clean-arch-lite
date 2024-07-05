package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/jrfernandodasilva/golang-clean-arch-lite/domain/entity"
	"github.com/jrfernandodasilva/golang-clean-arch-lite/domain/usecase"
	"github.com/jrfernandodasilva/golang-clean-arch-lite/infra/persistence"
)

func main() {
	repo := persistence.NewMemoryUserRepository()

	handle(
		usecase.NewCreateUser(repo),
		usecase.NewGetUser(repo),
		usecase.NewListUsersUseCase(repo),
	)
}

func handle(
	ucCreateUser *usecase.CreateUserUseCase,
	ucGetUser *usecase.GetUserUserCase,
	ucListUsers *usecase.ListUsersUseCase,
) {
	handleCreateUser(ucCreateUser)
	handleGetUser(ucGetUser)
	handleListUsers(ucListUsers)
}

func handleCreateUser(uc *usecase.CreateUserUseCase) {
	var input usecase.CreateUserInputDTO
	for _, user := range fakeUsers(3) {
		input.ID = user.ID
		input.Name = user.Name
		input.Email = user.Email

		output, err := uc.Execute(input) // execute use case "create user"
		if err != nil {
			fmt.Println("Error creating user:", err)
			return
		}

		fmt.Println("[uc - User created] -", output)
	}
}

func handleGetUser(uc *usecase.GetUserUserCase) {
	input := usecase.GetUserInputDTO{ID: 97}
	output, err := uc.Execute(input) // execute use case "get user"
	if err != nil {
		fmt.Println("Error getting user:", err)
		return
	}
	fmt.Println("[uc - Get user]", input.ID, "-", output)
}

func handleListUsers(uc *usecase.ListUsersUseCase) {
	output, err := uc.Execute() // execute use case "list users"
	if err != nil {
		fmt.Println("Error listing users:", err)
		return
	}

	jsonOutput, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("Error on output JSON convertion:", err)
		return
	}

	fmt.Println("[uc - List users]", "-", string(jsonOutput))
}

func fakeUsers(numUsers int) []*entity.User {
	users := make([]*entity.User, 0, numUsers)

	for i := 0; i < numUsers; i++ {
		var user entity.User
		if err := faker.FakeData(&user); err != nil {
			fmt.Println("Error on fake data:", err)
			return users
		}
		users = append(users, &user)
	}

	return users
}
