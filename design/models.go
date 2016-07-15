package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var sg = StorageGroup("ALMigthy Core Storage", func() {
	Description("The ALM Core storage group")
	Store("Core", gorma.Postgres, func() {
		Description("Describes the core Project/user/Team/Item store")
		Model("Project", func() {
			Description("This is the Project model")
			Field("id", gorma.UUID, func() {
				PrimaryKey()
				SQLTag("type:uuid default uuid_generate_v4()")
				Description("This is the ID PK field")
			})
			Field("name", gorma.String, func() {
				SQLTag("index:idx_project_name")

			})
		})
		Model("Identity", func() {
			Description("Describes a unique Person with the ALM")
			Field("id", gorma.UUID, func() {
				PrimaryKey()
				SQLTag("type:uuid default uuid_generate_v4()")
				Description("This is the ID PK field")
			})
			Field("full_name", gorma.String, func() {
				Description("The fullname of the Identity")
			})
			Field("image_url", gorma.String, func() {
				Description("The image URL for this Identity")
				DatabaseFieldName("image_url")
			})
			HasMany("emails", "User")
		})
		Model("User", func() {
			Description("Describes a User(single email) in any system")
			Field("id", gorma.UUID, func() {
				PrimaryKey()
				SQLTag("type:uuid default uuid_generate_v4()")
				Description("This is the ID PK field")
			})
			Field("email", gorma.String, func() {
				SQLTag("unique_index")
				Description("This is the unique email field")
			})
			BelongsTo("Identity", "type:uuid")
		})
		Model("Team", func() {
			Description("Describes a Team and how users share e.g. Permissions")
			Field("id", gorma.UUID, func() {
				PrimaryKey()
				SQLTag("type:uuid default uuid_generate_v4()")
				Description("This is the ID PK field")
			})
			Field("name", gorma.String, func() {
				Description("The display name of the team")
			})
			ManyToMany("Project", "team_project")
			ManyToMany("Permission", "team_permission")
			ManyToMany("Identity", "team_identity")
		})
		Model("Permission", func() {
			Description("Describes a single permissions and it's relation to a team")
			Field("id", gorma.UUID, func() {
				PrimaryKey()
				SQLTag("type:uuid default uuid_generate_v4()")
				Description("This is the ID PK field")
			})
			Field("name", gorma.String, func() {
				Description("The string value/name of the permission used in auth token")
			})
		})
	})
})
