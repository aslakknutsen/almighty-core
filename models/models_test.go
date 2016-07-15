package models

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

func TestMigration(t *testing.T) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	//migration.Perform(db)

	//createInitialSetup(db)
	//findIdentity(db)
	//findTeam(db)
	//findTeamByIdentity(db)

	userDB := NewUserDB(db)
	users, err := userDB.Query(ByEmails([]string{"aslak@4fs.no"}), WithIdentity())
	if err != nil {
		panic(err)
	}
	fmt.Println(users[0])

}

func findTeamByIdentity(db *gorm.DB) {
	var user User
	db.Preload("Identity").Where(&User{Email: "aslak@4fs.no"}).First(&user)

	var teams []Team
	db.Debug().Model(&user.Identity).Association("Identities").Find(&teams)
	fmt.Println("Teams: ", teams)
}

func findTeam(db *gorm.DB) {

	var team Team
	db.Debug().Preload("Permissions").Preload("Projects").Preload("Identities").Where(&Team{Name: "Core"}).First(&team)
	fmt.Println("Permissions: ", team.Permissions)
	fmt.Println("Projects: ", team.Projects)

}

func findIdentity(db *gorm.DB) {

	var user User
	db.Debug().Preload("Identity.Teams").Where(&User{Email: "aslak@4fs.no"}).First(&user)
	fmt.Println("A: ", user.Identity)
}

func createInitialSetup(db *gorm.DB) {
	users := NewUserDB(db)
	identities := NewIdentityDB(db)
	projects := NewProjectDB(db)
	teams := NewTeamDB(db)
	permissions := NewPermissionDB(db)
	ctx := context.Background()

	me := &Identity{FullName: "Aslak Knutsen"}
	err := identities.Add(ctx, me)
	if err != nil {
		panic(err)
	}

	firstEmail := &User{Email: "aslak@4fs.no", Identity: *me}
	err = users.Add(ctx, firstEmail)
	if err != nil {
		panic(err)
	}

	project := &Project{Name: "ALMighty"}
	err = projects.Add(ctx, project)
	if err != nil {
		panic(err)
	}

	var storedPermissions []Permission

	permissionStr := []string{"create.workitem", "read.workitem", "update.workitem", "delete.workitem"}
	for _, permission := range permissionStr {
		perm := &Permission{Name: permission}
		err := permissions.Add(ctx, perm)
		if err != nil {
			panic(err)
		}
		storedPermissions = append(storedPermissions, *perm)
	}

	team := &Team{Name: "Core", Permissions: storedPermissions, Identities: []Identity{*me}, Projects: []Project{*project}}
	err = teams.Add(ctx, team)
	if err != nil {
		panic(err)
	}

	fmt.Println(firstEmail)
	fmt.Println(team)
}
