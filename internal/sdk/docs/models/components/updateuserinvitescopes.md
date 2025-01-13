# UpdateUserInviteScopes

Scopes can only be removed or added if a user has those scopes.


## Supported Types

### UserRole

```go
updateUserInviteScopes := components.CreateUpdateUserInviteScopesUserRole(components.UserRole{/* values here */})
```

### 

```go
updateUserInviteScopes := components.CreateUpdateUserInviteScopesArrayOfScope([]components.Scope{/* values here */})
```

