// cli/menu.go
package cli

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/charmbracelet/huh"

	"gitlab.com/lidsol-ng/bitacora/database"
)

var (
	ErrInvalidUserType      = errors.New("invalid user type")
	ErrInvalidAccountNumber = errors.New("invalid account number")
	ErrInvalidRFC           = errors.New("invalid RFC")
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidSurName       = errors.New("invalid surname")
	ErrUserExists           = errors.New("user already exists")
	ErrUserNotExists        = errors.New("user does not exist")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrProjectNotExists     = errors.New("project does not exist")
	ErrInvalidMenuOption    = errors.New("invalid menu option")
)

func MainMenu(db *sql.DB) (bool, error) {
	var (
		option      int
		userID      string
		name        string
		surname     string
		uType       string
		email       string
		projectName string
		projectID   int
		manager     string
		description string
		err         error
	)

	huh.NewSelect[int]().
		Title("Bienvenido a la bitacora!\nSelecciona una opción.").
		Options(
			huh.NewOption("Registrar visita", 1),
			huh.NewOption("Registrar nuevo usuario", 2),
			huh.NewOption("Registrar nuevo proyecto", 3),
			huh.NewOption("Salir", 0),
		).
		Value(&option).
		Run()

	switch option {
	case 0:
		return true, nil
	case 1:
		userID = ""
		projectName = ""

		form := huh.NewForm(
			huh.NewGroup(
				// Ask for user ID
				huh.NewInput().
					Title("Ingrese su número de cuenta o RFC:").
					Value(&userID).
					Validate(func(str string) error {
						str = CleanString(str)
						if !database.GetUserExistence(db, str) {
							return ErrUserNotExists
						}
						return nil
					}),

				// Ask for project name
				huh.NewInput().
					Title("Ingrese a que proyecto viene:").
					Value(&projectName).
					Validate(func(str string) error {
						str = CleanString(str)
						projectID, err = database.GetProjectID(db, str)
						if err != nil {
							return ErrProjectNotExists
						}
						return nil
					}),
			),
		)

		// Run form
		err = form.Run()
		if err != nil {
			return false, err
		}

		// Add log
		userID = CleanString(userID)
		if err := database.AddLog(db, database.CreateLog(userID, projectID)); err != nil {
			return false, err
		}

		SaveLog(fmt.Sprintf("User %s visited project %s", userID, projectName))
	case 2:
		uType = ""
		userID = ""
		name = ""
		surname = ""
		email = ""

		// Ask for user type
		huh.NewSelect[string]().
			Title("Selecciona un tipo de usuario.").
			Options(
				huh.NewOption("Estudiante", "STUDENT"),
				huh.NewOption("Trabajador", "WORKER"),
			).
			Value(&uType).
			Run()

		// Ask user identifier
		if uType == "STUDENT" {
			huh.NewInput().
				Title("Ingrese su número de cuenta:").
				Value(&userID).
				Validate(func(str string) error {
					if !IsValidAccountNumber(str) {
						return ErrInvalidAccountNumber
					}
					if database.GetUserExistence(db, str) {
						return ErrUserExists
					}
					return nil
				}).
				Run()
		} else if uType == "WORKER" {
			huh.NewInput().
				Title("Ingrese su RFC:").
				Value(&userID).
				Validate(func(str string) error {
					str = CleanString(str)
					if !IsValidRFC(str) {
						return ErrInvalidRFC
					}
					if database.GetUserExistence(db, str) {
						return ErrUserExists
					}
					return nil
				}).
				Run()
		}

		form := huh.NewForm(
			huh.NewGroup(
				// Ask for name
				huh.NewInput().
					Title("Ingrese su nombre(s):").
					Value(&name).
					Validate(func(str string) error {
						if !IsValidName(str) {
							return ErrInvalidName
						}
						return nil
					}),

				// Ask for surname
				huh.NewInput().
					Title("Ingrese su(s) apellido(s):").
					Value(&surname).
					Validate(func(str string) error {
						if !IsValidName(str) {
							return ErrInvalidSurName
						}
						return nil
					}),

				// Ask for email
				huh.NewInput().
					Title("Ingrese su correo electrónico:").
					Value(&email).
					Validate(func(str string) error {
						if !IsValidEmail(str) {
							return ErrInvalidEmail
						}
						return nil
					}),
			),
		)

		// Run form
		err = form.Run()
		if err != nil {
			return false, err
		}

		// Add user
		userID = CleanString(userID)
		name = CleanString(name)
		surname = CleanString(surname)
		email = CleanString(email)
		if err := database.AddUser(db, database.CreateUser(userID, name, surname, uType, email)); err != nil {
			return false, err
		}

		SaveLog(fmt.Sprintf("User %s added to the database", userID))
	case 3:
		name = ""
		manager = ""
		description = ""

		form := huh.NewForm(
			huh.NewGroup(
				// Ask for project name
				huh.NewInput().
					Title("Ingrese el nombre del proyecto:").
					Value(&name).
					Validate(func(str string) error {
						str = CleanString(str)
						if database.GetProjectExistence(db, str) {
							return ErrUserExists
						}
						return nil
					}),
				
				// Ask for manager
				huh.NewInput().
					Title("Ingrese el identificador(no. cuenta / RFC) del responsable del proyecto:").
					Value(&manager).
					Validate(func(str string) error {
						str = CleanString(str)
						if !database.GetUserExistence(db, str) {
							return ErrUserNotExists
						}
						return nil
					}),
				
				// Ask for description
				huh.NewInput().
					Title("Ingrese la descripción del proyecto:").
					Value(&description),
			),
		)

		// Run form
		err = form.Run()
		if err != nil {
			return false, err
		}

		// Add project
		name = CleanString(name)
		manager = CleanString(manager)
		if err := database.AddProject(db, database.CreateProject(name, manager, description)); err != nil {
			return false, err
		}

		SaveLog(fmt.Sprintf("Project %s added to the database", name))
	default:
		fmt.Println("Opción no válida.")

	}

	return false, nil
}
