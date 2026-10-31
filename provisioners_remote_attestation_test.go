package linkedca

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// TestACMERemoteAttestation_roundTrip checks that an ACME provisioner with the
// remote-attest-01 challenge and its configuration survives protobuf and
// protojson, as remote management stores and serves it.
func TestACMERemoteAttestation_roundTrip(t *testing.T) {
	p := &Provisioner{
		Type: Provisioner_ACME,
		Name: "attested",
		Details: &ProvisionerDetails{Data: &ProvisionerDetails_ACME{ACME: &ACMEProvisioner{
			Challenges: []ACMEProvisioner_ChallengeType{ACMEProvisioner_REMOTE_ATTEST_01},
			RemoteAttestation: &ACMERemoteAttestation{
				Require:     true,
				Formats:     []string{"tdx", "sev-snp", "confidential-space"},
				ClaimsHint:  []string{"hwmodel"},
				TdxRoots:    []byte("tdx roots"),
				SevSnpRoots: []byte("sev-snp roots"),
				ConfidentialSpace: &ACMEConfidentialSpace{
					Roots:                    []byte("roots"),
					Audience:                 "https://ca.example.com/acme/attested/directory",
					MaxTokenAge:              "5m",
					ImageDigests:             []string{"sha256:00"},
					ImageSignerKeyIds:        []string{"c0ffee"},
					RequireSupportAttributes: []string{"STABLE"},
					Hwmodels:                 []string{"GCP_INTEL_TDX"},
					Swnames:                  []string{"CONFIDENTIAL_SPACE"},
					ProjectIds:               []string{"my-project"},
					ServiceAccounts:          []string{"workload@my-project.iam.gserviceaccount.com"},
				},
			},
		}}},
	}

	b, err := proto.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	got := new(Provisioner)
	if err := proto.Unmarshal(b, got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(p, got) {
		t.Errorf("protobuf round trip: got %v, want %v", got, p)
	}

	j, err := protojson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	got = new(Provisioner)
	if err := protojson.Unmarshal(j, got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(p, got) {
		t.Errorf("protojson round trip: got %v, want %v", got, p)
	}

	if name := ACMEProvisioner_REMOTE_ATTEST_01.String(); name != "REMOTE_ATTEST_01" {
		t.Errorf("challenge name: got %q", name)
	}
}
