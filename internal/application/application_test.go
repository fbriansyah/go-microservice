package application

import (
	"context"
	"reflect"
	"testing"

	"github.com/fbriansyah/go-microservice/internal/application/commands"
	"github.com/fbriansyah/go-microservice/internal/application/queries"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		configs []ApplicationConfig
		want    *Application
		wantErr bool
	}{
		{
			name: "success",
			configs: []ApplicationConfig{
				WithUserRepo(userRepo),
				WithLogger(logger),
			},
			want: &Application{
				userRepo: userRepo,
				logger:   logger,
				appCommands: appCommands{
					CreateUserHandler: commands.NewUserHandler(userRepo, logger),
				},
				appQueries: appQueries{
					FindUserByEmailHandler: queries.NewFindUserByEmailHandler(userRepo),
				},
			},
			wantErr: false,
		},
		{
			name: "failure - missing user repo",
			configs: []ApplicationConfig{
				WithLogger(logger),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.configs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreateUserCommand(t *testing.T) {
	err := app.appCommands.CreateUser(context.Background(), commands.CreateUserCmd{
		Name:     "Test User",
		Email:    "test@test.com",
		Password: "123qweasdzxc",
	})

	if err != nil {
		t.Errorf("CreateUserCommand() error = %v, wantErr %v", err, nil)
	}
}

func Test_FindByEmailQuery(t *testing.T) {
	q := queries.FindUserByEmailQuery{
		Email:    "dummy@example.com",
		Password: "123qweasdzxc",
	}
	user, err := app.appQueries.FindUserByEmail(context.Background(), q)
	if err != nil {
		t.Errorf("FindByEmailQuery() error = %v, wantErr %v", err, nil)
	}
	if user.Email != q.Email {
		t.Errorf("FindByEmailQuery() = %v, want %v", user.Email, q.Email)
	}
}
