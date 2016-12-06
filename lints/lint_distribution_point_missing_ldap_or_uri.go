// lint_distribution_point_missing_ldap_or_uri.go
/************************************************
RFC 5280: 4.2.1.13
When present, DistributionPointName SHOULD include at least one LDAP or HTTP URI.
************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
	"strings"
)

type distribNoLDAPorURI struct {
	// Internal data here
}

func (l *distribNoLDAPorURI) Initialize() error {
	return nil
}

func (l *distribNoLDAPorURI) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsExtInCert(c, util.CrlDistOID)
}

func (l *distribNoLDAPorURI) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, point := range c.CRLDistributionPoints {
		if point = strings.ToLower(point); strings.HasPrefix(point, "http://") || strings.HasPrefix(point, "ldap://") {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: Warn}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "distribution_point_missing_ldap_or_uri",
		Description:   "When present in the CRLDistributionPoints extension, DistributionPointName SHOULD include at least one LDAP or HTTP URI.",
		Providence:    "RFC 5280: 4.2.1.13",
		EffectiveDate: util.RFC5280Date,
		Test:          &distribNoLDAPorURI{}})
}