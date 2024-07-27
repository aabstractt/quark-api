package source

type Permissible interface {
    // HasPermission returns true if the user has the given permission
    HasPermission(perm string) bool

    // SetPermissions sets the user's permissions
    SetPermissions(perms []string)

    // AddPermission adds the permission to the user
    AddPermission(perm string)

    // Permissions returns the user's permissions
    Permissions() []string
}
