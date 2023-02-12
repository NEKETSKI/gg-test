// Package repository provides an interface that describes the way to interact with the database
// required for the service to work.
package repository

// Repository defines the database methods required for the service to work.
type Repository interface {
	StoreStatusCode(requestUrl string, code int) error
	Close() error
}
