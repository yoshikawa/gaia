package schema

// UserListResolver struct
type UserListResolver struct {
	UsersField []*UserSchema
	MetaField  *MetaResolver
}

// Users function
func (user *UserListResolver) Users() []*UserSchema {
	return user.UsersField
}

// Meta function
func (user *UserListResolver) Meta() *MetaResolver {
	return user.MetaField
}
