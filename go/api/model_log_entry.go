/*
 * k8l
 *
 * Kubernetes light logs API
 *
 * API version: 1.0
 * Contact: mogui83@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"time"
)

// LogEntry - A Log Entry
type LogEntry struct {

	Timestamp time.Time `json:"timestamp,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Container string `json:"container,omitempty"`

	Pod string `json:"pod,omitempty"`

	Message string `json:"message,omitempty"`
}
