package constants

const (
	RoleAdmin  = "admin"
	RoleAuthor = "author"
	RoleViewer = "viewer"
)

func GetRoles() []string {
	return []string{RoleAdmin, RoleAuthor, RoleViewer}
}
