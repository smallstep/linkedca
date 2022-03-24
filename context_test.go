package linkedca

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdminFromContext(t *testing.T) {
	t.Parallel()

	exp := new(Admin)

	got := AdminFromContext(NewContextWithAdmin(context.Background(), exp))
	assert.Same(t, exp, got)
}

func TestAdminFromContextPanics(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { AdminFromContext(context.Background()) })
}

func TestProvisionerFromContext(t *testing.T) {
	t.Parallel()

	exp := new(Provisioner)

	got := ProvisionerFromContext(NewContextWithProvisioner(context.Background(), exp))
	assert.Same(t, exp, got)
}

func TestProvisionerFromContextPanics(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { ProvisionerFromContext(context.Background()) })
}
