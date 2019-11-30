package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DataStore interface {
	Find(string) (*FeatureFlag, error)
	Save(*FeatureFlag) (*FeatureFlag, error)
	Update(*FeatureFlag) (*FeatureFlag, error)
	FindAll() ([]FeatureFlag, error)
	Delete(string) error
}

type MySQLDataStore struct {
	db *sql.DB
}

var (
	createDatabase string = `CREATE DATABASE IF NOT EXISTS rocketlauncherff;`
	query          string = `CREATE TABLE IF NOT EXISTS rocketlauncherff.flags (
id BIGINT AUTO_INCREMENT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NULL,
		enabled  BOOL DEFAULT FALSE,
		created DATETIME,
		updated DATETIME
	    );`
	insertQuery    string = `INSERT INTO rocketlauncherff.flags (name, description, created, updated) VALUES (?, ?, ?, ?)`
	selectAllQuery string = `SELECT name, description, enabled, id FROM rocketlauncherff.flags`
	selectQuery    string = `SELECT name, description, enabled, id FROM rocketlauncherff.flags WHERE name = ?`
	updateQuery    string = `UPDATE rocketlauncherff.flags SET name=?, description=?, enabled=?, updated=? WHERE id = ?`
	deleteQuery    string = `DELETE FROM rocketlauncherff.flags WHERE id=?`
)

func NewDataStore(dbUrl string) (*MySQLDataStore, error) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createDatabase)
	if err != nil {
		return nil, err
	}
	return &MySQLDataStore{db: db}, nil
}

func (datastore MySQLDataStore) Find(name string) (*FeatureFlag, error) {
	ff := FeatureFlag{}
	err := datastore.db.QueryRow(selectQuery, name).Scan(&ff.Name, &ff.Description, &ff.Enabled, &ff.Id)
	if err != nil {
		return nil, err
	}
	return &ff, nil
}

func (datastore MySQLDataStore) Save(ff *FeatureFlag) (*FeatureFlag, error) {
	_, err := datastore.db.Exec(insertQuery, ff.Name, ff.Description, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	return ff, nil
}

func (datastore MySQLDataStore) Update(ff *FeatureFlag) (*FeatureFlag, error) {
	_, err := datastore.db.Exec(updateQuery, ff.Name, ff.Description, ff.Enabled, time.Now(), ff.Id)
	if err != nil {
		return nil, err
	}
	return ff, nil
}

func (datastore MySQLDataStore) Delete(id string) error {
	_, err := datastore.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (datastore MySQLDataStore) FindAll() ([]FeatureFlag, error) {
	var flags []FeatureFlag
	rows, err := datastore.db.Query(selectAllQuery)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ff FeatureFlag
		err := rows.Scan(&ff.Name, &ff.Description, &ff.Enabled, &ff.Id)
		if err != nil {
			return nil, err
		}
		flags = append(flags, ff)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return flags, nil
}
