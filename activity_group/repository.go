package activity_group

import (
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	FindAll() ([]ActivityGroup, error)
	FindByID(ID string) (ActivityGroup, error)
	Create(activityGroup ActivityGroup) (ActivityGroup, error)
	Delete(ID string) (bool, error)
	Update(activityGroup ActivityGroup) (ActivityGroup, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]ActivityGroup, error) {
	activityGroups := []ActivityGroup{}

	results, err := r.db.Query("SELECT id, title, email, created_at, updated_at, deleted_at FROM activity_group WHERE deleted_at IS NULL")

	if err != nil {
		return activityGroups, err
	}

	for results.Next() {
		activityGroup := ActivityGroup{}
		err := results.Scan(&activityGroup.ID, &activityGroup.Title, &activityGroup.Email, &activityGroup.CreatedAt, &activityGroup.UpdatedAt, &activityGroup.DeletedAt)

		if err != nil {
			return activityGroups, err
		}

		activityGroups = append(activityGroups, activityGroup)
	}

	return activityGroups, nil
}

func (r *repository) FindByID(ID string) (ActivityGroup, error) {
	activityGroup := ActivityGroup{}

	err := r.db.QueryRow("SELECT id, title, email, created_at, updated_at, deleted_at FROM activity_group WHERE id=?", ID).Scan(&activityGroup.ID, &activityGroup.Title, &activityGroup.Email, &activityGroup.CreatedAt, &activityGroup.UpdatedAt, &activityGroup.DeletedAt)

	if err != nil {
		return activityGroup, err
	}

	return activityGroup, nil
}

func (r *repository) Create(activityGroup ActivityGroup) (ActivityGroup, error) {
	insert, err := r.db.Query("INSERT INTO activity_group(title, email) VALUES (?,?)", activityGroup.Title, activityGroup.Email)
	if err != nil {
		return activityGroup, err
	}

	defer insert.Close()

	return activityGroup, nil
}

func (r *repository) Delete(ID string) (bool, error) {
	sql, err := r.db.Exec("DELETE FROM activity_group WHERE id=?", ID)
	if err != nil {
		return false, err
	}

	rows, err := sql.RowsAffected()

	if err != nil {
		return false, err
	}

	return rows > 0, nil
}

func (r *repository) Update(activityGroup ActivityGroup) (ActivityGroup, error) {
	sql, err := r.db.Exec("UPDATE activity_group SET title=?, email=?, updated_at=NOW() WHERE id=?", activityGroup.Title, activityGroup.Email, activityGroup.ID)

	if err != nil {
		return activityGroup, err
	}

	rows, err := sql.RowsAffected()

	if err != nil {
		return activityGroup, err
	}

	if rows < 1 {
		return activityGroup, errors.New(fmt.Sprintf("Activity with ID %d Not Found", activityGroup.ID))
	}

	return activityGroup, nil
}
