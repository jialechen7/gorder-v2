package domain

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

// DomainError represents a business domain error with gRPC code
type DomainError interface {
	error
	GRPCCode() codes.Code
}

// StockError implements DomainError for stock-related errors
type StockError struct {
	message  string
	grpcCode codes.Code
}

func (e StockError) Error() string {
	return e.message
}

func (e StockError) GRPCCode() codes.Code {
	return e.grpcCode
}

func NewItemNotFoundError(itemID string) DomainError {
	return StockError{
		message:  fmt.Sprintf("item %s not found", itemID),
		grpcCode: codes.NotFound,
	}
}

func NewItemsNotFoundError(itemIDs []string) DomainError {
	return StockError{
		message:  fmt.Sprintf("items not found: %v", itemIDs),
		grpcCode: codes.NotFound,
	}
}

func NewInsufficientStockError(itemID string, requested, available int32) DomainError {
	return StockError{
		message:  fmt.Sprintf("item %s: requested %d, available %d", itemID, requested, available),
		grpcCode: codes.FailedPrecondition,
	}
}

func NewInvalidArgumentError(message string) DomainError {
	return StockError{
		message:  message,
		grpcCode: codes.InvalidArgument,
	}
}
