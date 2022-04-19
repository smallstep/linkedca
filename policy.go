package linkedca

// Deduplicate removes duplicate values from the Policy
func (p *Policy) Deduplicate() {
	if p == nil {
		return
	}
	if p.GetX509() != nil {
		if p.GetX509().GetAllow() != nil {
			allow := p.GetX509().GetAllow()
			if len(allow.Dns) > 1 {
				allow.Dns = removeDuplicates(allow.Dns)
			}
			if len(allow.Ips) > 1 {
				allow.Ips = removeDuplicates(allow.Ips)
			}
			if len(allow.Emails) > 1 {
				allow.Emails = removeDuplicates(allow.Emails)
			}
			if len(allow.Uris) > 1 {
				allow.Uris = removeDuplicates(allow.Uris)
			}
		}
		if p.GetX509().GetDeny() != nil {
			deny := p.GetX509().GetDeny()
			if len(deny.Dns) > 1 {
				deny.Dns = removeDuplicates(deny.Dns)
			}
			if len(deny.Ips) > 1 {
				deny.Ips = removeDuplicates(deny.Ips)
			}
			if len(deny.Emails) > 1 {
				deny.Emails = removeDuplicates(deny.Emails)
			}
			if len(deny.Uris) > 1 {
				deny.Uris = removeDuplicates(deny.Uris)
			}
		}
	}
	if p.GetSsh() != nil {
		if p.GetSsh().GetHost() != nil {
			if p.GetSsh().GetHost().GetAllow() != nil {
				allow := p.GetSsh().GetHost().GetAllow()
				if len(allow.Dns) > 1 {
					allow.Dns = removeDuplicates(allow.Dns)
				}
				if len(allow.Ips) > 1 {
					allow.Ips = removeDuplicates(allow.Ips)
				}
				if len(allow.Principals) > 1 {
					allow.Principals = removeDuplicates(allow.Principals)
				}
			}
			if p.GetSsh().GetHost().GetDeny() != nil {
				deny := p.GetSsh().GetHost().GetDeny()
				if len(deny.Dns) > 1 {
					deny.Dns = removeDuplicates(deny.Dns)
				}
				if len(deny.Ips) > 1 {
					deny.Ips = removeDuplicates(deny.Ips)
				}
				if len(deny.Principals) > 1 {
					deny.Principals = removeDuplicates(deny.Principals)
				}
			}
		}
		if p.GetSsh().GetUser() != nil {
			if p.GetSsh().GetUser().GetAllow() != nil {
				allow := p.GetSsh().GetUser().GetAllow()
				if len(allow.Emails) > 1 {
					allow.Emails = removeDuplicates(allow.Emails)
				}
				if len(allow.Principals) > 1 {
					allow.Principals = removeDuplicates(allow.Principals)
				}
			}
			if p.GetSsh().GetUser().GetDeny() != nil {
				deny := p.GetSsh().GetUser().GetDeny()
				if len(deny.Emails) > 1 {
					deny.Emails = removeDuplicates(deny.Emails)
				}
				if len(deny.Principals) > 1 {
					deny.Principals = removeDuplicates(deny.Principals)
				}
			}
		}
	}
}

// removeDuplicates returns a new slice with duplicate values
// in a slice of strings removed. It retains the order of elements
// in the source slice.
func removeDuplicates(strSlice []string) []string {
	result := []string{}
	if len(strSlice) == 0 {
		return result
	}
	keys := make(map[string]bool)
	for _, item := range strSlice {
		if _, value := keys[item]; !value && item != "" { // skip empty constraints
			keys[item] = true
			result = append(result, item)
		}
	}
	return result
}
