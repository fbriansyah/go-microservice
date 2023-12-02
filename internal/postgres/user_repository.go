package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fbriansyah/go-microservice/internal/application/domain"
	"github.com/fbriansyah/go-microservice/internal/application/domain/user"
	"github.com/google/uuid"
)

type UserRepo struct {
	db        *sql.DB
	tableName string
}

func (r UserRepo) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}

// Save persists the given user to the database.
func (r *UserRepo) Save(ctx context.Context, user *user.User) error {
	query := `INSERT INTO %s (name, email, password, status, created_at, update_at) VALUES ($1, $2 , $3, $4, now(), now())`
	_, err := r.db.ExecContext(ctx, r.table(query), user.Name, user.Email, user.Password, user.Status)
	return err
}

// Update updates the status of a user to "deleted"
func (r *UserRepo) Deleted(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE %s SET status = $1, update_at = now() WHERE id =$2`
	_, err := r.db.ExecContext(ctx, r.table(query), int(domain.Deleted), id)
	return err
}

// GetAll retrieves all users from the database.
func (r *UserRepo) GetAll(ctx context.Context) ([]*user.User, error) {
	// Define the SQL query to retrieve all users.
	query := `SELECT id, name, email, status FROM %s`

	// Execute the SQL query and retrieve the results.
	rows, err := r.db.QueryContext(ctx, r.table(query))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to store the retrieved users.
	var users []*user.User

	// Iterate over the rows and scan the data into the user struct.
	for rows.Next() {
		var u user.User
		var status int
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &status); err != nil {
			return nil, err
		}
		// Append the user to the slice.
		u.Status = domain.Status(status)
		users = append(users, &u)
	}

	// Return the retrieved users.
	return users, nil
}

// GetByID retrieves a user by its ID.
func (r *UserRepo) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	// Define the SQL query to retrieve a user by its ID.
	query := `SELECT id, name, email, status FROM %s WHERE id = $1`

	// Execute the SQL query and retrieve the result.
	row := r.db.QueryRowContext(ctx, r.table(query), id)

	// Create a new user struct to store the retrieved data.
	var u user.User

	// Scan the data into the user struct.
	var status int
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &status); err != nil {
		return nil, err
	}

	// Set the user status.
	u.Status = domain.Status(status)

	// Return the retrieved user.
	return &u, nil
}

// Update updates the user details in the database.
func (r *UserRepo) Update(ctx context.Context, user *user.User) error {
	// Define the SQL query to update the user details.
	query := `UPDATE %s SET name = $1, email = $2, password = $3, status = $4, update_at = now() WHERE id =$5`

	// Execute the SQL query with the user details.
	_, err := r.db.ExecContext(ctx, r.table(query), user.Name, user.Email, user.Password, int(user.Status), user.ID)
	return err
}

func NewUserRepo(db *sql.DB, tableName string) *UserRepo {
	return &UserRepo{
		db:        db,
		tableName: tableName,
	}
}

var _ user.Repository = (*UserRepo)(nil)
