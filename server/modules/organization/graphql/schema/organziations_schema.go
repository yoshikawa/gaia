package schema

// OrganizationListResolver struct
type OrganizationListResolver struct {
	OrganizationField []*OrganizationSchema
}

// Organizations function
func (organization *OrganizationListResolver) Organizations() []*OrganizationSchema {
	return organization.OrganizationField
}
