package models

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"

	"github.com/almighty/almighty-core/migration"
	"github.com/almighty/almighty-core/models"
	"github.com/jinzhu/gorm"
)

func TestMigration(t *testing.T) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	migration.Perform(db)

	//createInitialSetup(db)
	//findIdentity(db)
	//findTeam(db)
	findTeamByIdentity(db)
}

func findTeamByIdentity(db *gorm.DB) {
	var user models.User
	db.Preload("Identity").Where(&models.User{Email: "aslak@4fs.no"}).First(&user)

	var teams []models.Team
	db.Debug().Model(&user.Identity).Association("Identities").Find(&teams)
	fmt.Println("Teams: ", teams)
}

func findTeam(db *gorm.DB) {

	var team models.Team
	db.Debug().Preload("Permissions").Preload("Projects").Preload("Identities").Where(&models.Team{Name: "Core"}).First(&team)
	fmt.Println("Permissions: ", team.Permissions)
	fmt.Println("Projects: ", team.Projects)

}

func findIdentity(db *gorm.DB) {

	var user models.User
	db.Debug().Preload("Identity.Teams").Where(&models.User{Email: "aslak@4fs.no"}).First(&user)
	fmt.Println("A: ", user.Identity)
}

func createInitialSetup(db *gorm.DB) {
	users := models.NewUserDB(*db)
	identities := models.NewIdentityDB(*db)
	projects := models.NewProjectDB(*db)
	teams := models.NewTeamDB(*db)
	permissions := models.NewPermissionDB(*db)
	ctx := context.Background()

	me, err := identities.Add(ctx, &models.Identity{FullName: "Aslak Knutsen"})
	if err != nil {
		panic(err)
	}

	firstEmail, err := users.Add(ctx, &models.User{Email: "aslak@4fs.no", Identity: *me})
	if err != nil {
		panic(err)
	}

	project, err := projects.Add(ctx, &models.Project{Name: "ALMighty"})
	if err != nil {
		panic(err)
	}

	var storedPermissions []models.Permission

	permissionStr := []string{"create.workitem", "read.workitem", "update.workitem", "delete.workitem"}
	for _, permission := range permissionStr {
		perm, err := permissions.Add(ctx, &models.Permission{Name: permission})
		if err != nil {
			panic(err)
		}
		storedPermissions = append(storedPermissions, *perm)
	}

	team, err := teams.Add(ctx, &models.Team{Name: "Core", Permissions: storedPermissions, Identities: []models.Identity{*me}, Projects: []models.Project{*project}})
	if err != nil {
		panic(err)
	}

	fmt.Println(firstEmail)
	fmt.Println(team)
}
