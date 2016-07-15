package migration

import (
	"github.com/almighty/almighty-core/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// Perform executes the required migration of the database on startup
func Perform(db *gorm.DB) {

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.AutoMigrate(
		&models.Project{},
		&models.Identity{},
		&models.User{},
		&models.Permission{},
		&models.Team{})

	createDefaults(db)
}

func createDefaults(db *gorm.DB) {
	projects := models.NewProjectDB(db)
	teams := models.NewTeamDB(db)
	permissions := models.NewPermissionDB(db)
	ctx := context.Background()

	project := &models.Project{Name: "ALMighty"}
	err := projects.Add(ctx, project)
	if err != nil {
		panic(err)
	}

	var storedPermissions []models.Permission

	permissionStr := []string{"create.workitem", "read.workitem", "update.workitem", "delete.workitem"}
	for _, permission := range permissionStr {
		perm := &models.Permission{Name: permission}
		err := permissions.Add(ctx, perm)
		if err != nil {
			panic(err)
		}
		storedPermissions = append(storedPermissions, *perm)
	}

	err = teams.Add(ctx, &models.Team{Name: "Core", Permissions: storedPermissions, Identities: []models.Identity{}, Projects: []models.Project{*project}})
	if err != nil {
		panic(err)
	}
}
