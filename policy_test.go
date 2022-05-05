package linkedca

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "empty-slice",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "single-item",
			input: []string{"x"},
			want:  []string{"x"},
		},
		{
			name:  "ok",
			input: []string{"x", "y", "x", "z", "x", "z", "y"},
			want:  []string{"x", "y", "z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_Deduplicate(t *testing.T) {
	tests := []struct {
		name     string
		policy   *Policy
		expected *Policy
	}{
		{
			name:     "nil",
			policy:   nil,
			expected: nil,
		},
		{
			name: "x509",
			policy: &Policy{
				X509: &X509Policy{
					Allow: &X509Names{
						Dns:    []string{"*.local", "*.example.com", "*.local"},
						Ips:    []string{"192.168.0.0/24", "10.10.10.0/24", "192.168.0.0/24"},
						Emails: []string{"@example.com", "@local", "@example.com"},
						Uris:   []string{"*.local", "*.example.com", "*.local"},
					},
					Deny: &X509Names{
						Dns:    []string{"*.local", "*.example.com", "*.local"},
						Ips:    []string{"192.168.0.0/24", "10.10.10.0/24", "192.168.0.0/24"},
						Emails: []string{"@example.com", "@local", "@example.com"},
						Uris:   []string{"*.local", "*.example.com", "*.local"},
					},
				},
			},
			expected: &Policy{
				X509: &X509Policy{
					Allow: &X509Names{
						Dns:    []string{"*.local", "*.example.com"},
						Ips:    []string{"192.168.0.0/24", "10.10.10.0/24"},
						Emails: []string{"@example.com", "@local"},
						Uris:   []string{"*.local", "*.example.com"},
					},
					Deny: &X509Names{
						Dns:    []string{"*.local", "*.example.com"},
						Ips:    []string{"192.168.0.0/24", "10.10.10.0/24"},
						Emails: []string{"@example.com", "@local"},
						Uris:   []string{"*.local", "*.example.com"},
					},
				},
			},
		},
		{
			name: "ssh-host",
			policy: &Policy{
				Ssh: &SSHPolicy{
					Host: &SSHHostPolicy{
						Allow: &SSHHostNames{
							Dns:        []string{"*.local", "*.example.com", "*.local"},
							Ips:        []string{"192.168.0.0/24", "10.10.10.0/24", "192.168.0.0/24"},
							Principals: []string{"localhost", "test", "localhost"},
						},
						Deny: &SSHHostNames{
							Dns:        []string{"*.local", "*.example.com", "*.local"},
							Ips:        []string{"192.168.0.0/24", "10.10.10.0/24", "192.168.0.0/24"},
							Principals: []string{"localhost", "test", "localhost"},
						},
					},
				},
			},
			expected: &Policy{
				Ssh: &SSHPolicy{
					Host: &SSHHostPolicy{
						Allow: &SSHHostNames{
							Dns:        []string{"*.local", "*.example.com"},
							Ips:        []string{"192.168.0.0/24", "10.10.10.0/24"},
							Principals: []string{"localhost", "test"},
						},
						Deny: &SSHHostNames{
							Dns:        []string{"*.local", "*.example.com"},
							Ips:        []string{"192.168.0.0/24", "10.10.10.0/24"},
							Principals: []string{"localhost", "test"},
						},
					},
				},
			},
		},
		{
			name: "ssh-user",
			policy: &Policy{
				Ssh: &SSHPolicy{
					User: &SSHUserPolicy{
						Allow: &SSHUserNames{
							Emails:     []string{"@example.com", "@local", "@example.com"},
							Principals: []string{"user", "root", "user"},
						},
						Deny: &SSHUserNames{
							Emails:     []string{"@example.com", "@local", "@example.com"},
							Principals: []string{"user", "root", "user"},
						},
					},
				},
			},
			expected: &Policy{
				Ssh: &SSHPolicy{
					User: &SSHUserPolicy{
						Allow: &SSHUserNames{
							Emails:     []string{"@example.com", "@local"},
							Principals: []string{"user", "root"},
						},
						Deny: &SSHUserNames{
							Emails:     []string{"@example.com", "@local"},
							Principals: []string{"user", "root"},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.policy.Deduplicate()
			assert.Equal(t, tt.expected, tt.policy)
		})
	}
}
