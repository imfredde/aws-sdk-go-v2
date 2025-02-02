// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/mwaa/document"
	smithy "github.com/aws/smithy-go"
)

// Access to the Apache Airflow Web UI or CLI has been denied due to insufficient
// permissions. To learn more, see [Accessing an Amazon MWAA environment].
//
// [Accessing an Amazon MWAA environment]: https://docs.aws.amazon.com/mwaa/latest/userguide/access-policies.html
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// InternalServerException: An internal error has occurred.
type InternalServerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// ResourceNotFoundException: The resource is not available.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception indicating that a client-side error occurred during the Apache
// Airflow REST API call.
type RestApiClientException struct {
	Message *string

	ErrorCodeOverride *string

	RestApiStatusCode *int32
	RestApiResponse   document.Interface

	noSmithyDocumentSerde
}

func (e *RestApiClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RestApiClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RestApiClientException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RestApiClientException"
	}
	return *e.ErrorCodeOverride
}
func (e *RestApiClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception indicating that a server-side error occurred during the Apache
// Airflow REST API call.
type RestApiServerException struct {
	Message *string

	ErrorCodeOverride *string

	RestApiStatusCode *int32
	RestApiResponse   document.Interface

	noSmithyDocumentSerde
}

func (e *RestApiServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RestApiServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RestApiServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RestApiServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *RestApiServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// ValidationException: The provided input is not valid.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
