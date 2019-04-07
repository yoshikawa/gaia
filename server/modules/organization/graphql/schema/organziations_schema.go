package schema

// OrganizationListResolver struct
type OrganizationListResolver struct {
	OrganizationsField []*OrganizationSchema
	MetaField          *MetaResolver
}

// Organizations function
func (organization *OrganizationListResolver) Organizations() []*OrganizationSchema {
	return organization.OrganizationsField
}

// Meta function
func (organization *OrganizationListResolver) Meta() *MetaResolver {
	return organization.MetaField
}
