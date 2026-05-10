// Package security racchiude controlli di sicurezza cognitiva (capitolo 13, Volume 2).
package security

import (
	"errors"
	"fmt"
)

var (
	// ErrTenantLeak quando policy vieta retrieval cross‑tenant.
	ErrTenantLeak = errors.New("security: cross-tenant access denied")
)

// TenantScope verifica isolamento tra tenant per query simulate.
func TenantScope(requestingTenant, resourceTenant string, isolate bool) error {
	if !isolate {
		return nil
	}
	if requestingTenant == "" || resourceTenant == "" {
		return fmt.Errorf("security: missing tenant scope")
	}
	if requestingTenant != resourceTenant {
		return ErrTenantLeak
	}
	return nil
}
