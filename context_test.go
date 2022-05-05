package linkedca

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdminFromContext(t *testing.T) {
	t.Parallel()

	// nil admin; expect false
	var exp *Admin
	got, ok := AdminFromContext(NewContextWithAdmin(context.Background(), exp))
	assert.Same(t, exp, got)
	assert.False(t, ok)

	// non-nil admin; expect true
	exp = new(Admin)
	got, ok = AdminFromContext(NewContextWithAdmin(context.Background(), exp))
	assert.Same(t, exp, got)
	assert.True(t, ok)
}

func TestMustAdminFromContext(t *testing.T) {
	t.Parallel()

	exp := new(Admin)
	got := MustAdminFromContext(NewContextWithAdmin(context.Background(), exp))
	assert.Same(t, exp, got)
}

func TestMustAdminFromContextPanics(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { MustAdminFromContext(context.Background()) })
}

func TestProvisionerFromContext(t *testing.T) {
	t.Parallel()

	// nil Provisioner; expect false
	var exp *Provisioner
	got, ok := ProvisionerFromContext(NewContextWithProvisioner(context.Background(), exp))
	assert.Same(t, exp, got)
	assert.False(t, ok)

	// non-nil Provisioner; expect true
	exp = new(Provisioner)
	got, ok = ProvisionerFromContext(NewContextWithProvisioner(context.Background(), exp))
	assert.Same(t, exp, got)
	assert.True(t, ok)
}

func TestMustProvisionerFromContext(t *testing.T) {
	t.Parallel()

	exp := new(Provisioner)
	got := MustProvisionerFromContext(NewContextWithProvisioner(context.Background(), exp))
	assert.Same(t, exp, got)
}

func TestMustProvisionerFromContextPanics(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { MustProvisionerFromContext(context.Background()) })
}

func TestExternalAccountKeyFromContext(t *testing.T) {
	t.Parallel()

	// nil EABKey; expect false
	var exp *EABKey
	got, ok := ExternalAccountKeyFromContext(NewContextWithExternalAccountKey(context.Background(), exp))
	assert.Same(t, exp, got)
	assert.False(t, ok)

	// non-nil EABKey; expect true
	exp = new(EABKey)
	got, ok = ExternalAccountKeyFromContext(NewContextWithExternalAccountKey(context.Background(), exp))
	assert.Same(t, exp, got)
	assert.True(t, ok)
}

func TestMustExternalAccountKeyFromContext(t *testing.T) {
	t.Parallel()

	exp := new(EABKey)
	got := MustExternalAccountKeyFromContext(NewContextWithExternalAccountKey(context.Background(), exp))
	assert.Same(t, exp, got)
}

func TestExternalAccountKeyFromContextPanics(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { MustExternalAccountKeyFromContext(context.Background()) })
}
