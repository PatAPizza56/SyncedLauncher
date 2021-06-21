package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber"
	_ "github.com/jackc/pgx/stdlib"
)

var (
	DB  *sql.DB
	err error
)

func Open() {
	DB, err = sql.Open("pgx", "host=database-do-user-9142180-0.b.db.ondigitalocean.com port=25060 dbname=defaultdb user=doadmin password=mmzued8avl109eeo sslmode=require")
	if err != nil {
		fmt.Println(err)
	}
}

func Close() {
	DB.Close()
}

func Insert(table string, fields []string, values []string, id *int) (error, int) {
	if len(fields) != len(values) {
		return errors.New("Failed to write to database, unequal number of fields and values"), fiber.StatusBadRequest
	}

	request := fmt.Sprintf(`INSERT INTO "%v" (`, table)
	for i := 0; i < len(fields); i++ {
		if fields[i] == "" {
			return errors.New("Failed to create request, all fields must not be empty"), fiber.StatusBadRequest
		}
		request += fmt.Sprintf(`"%v", `, fields[i])
	}
	request = strings.TrimSuffix(request, ", ") + ") VALUES ("
	for i := 0; i < len(fields); i++ {
		if values[i] == "" {
			return errors.New("Failed to create request, all values must not be empty"), fiber.StatusBadRequest
		}
		request += fmt.Sprintf(`'%v', `, values[i])
	}
	request = strings.TrimSuffix(request, ", ") + `) RETURNING "ID";`

	err := DB.QueryRow(request).Scan(id)
	if err != nil {
		return errors.New("Failed to execute database command, please recheck all fields of information. (Some fields may need to be unique)"), fiber.StatusBadRequest
	}

	return nil, 0
}

func Select(table string, param string, value string, objects ...interface{}) (error, int) {
	request := fmt.Sprintf(`SELECT * FROM "%v" WHERE "%v" = '%v';`, table, param, value)

	err := DB.QueryRow(request).Scan(objects...)
	if err != nil {
		return errors.New(fmt.Sprintf(`Failed to locate "%v", be sure get value and method are valid, and not null`, value)), fiber.StatusNotFound
	}

	return nil, 0
}

func Update(table string, param string, value string, fields []string, values []string) (error, int) {
	if len(fields) != len(values) {
		return errors.New("Failed to write to database, unequal number of fields and values"), fiber.StatusBadRequest
	}

	request := fmt.Sprintf(`UPDATE "%v" SET `, table)
	for i := 0; i < len(fields); i++ {
		if fields[i] == "" {
			return errors.New("Failed to create request, all fields must not be empty"), fiber.StatusBadRequest
		} else if values[i] == "" {
			return errors.New("Failed to create request, all values must not be empty"), fiber.StatusBadRequest
		}
		request += fmt.Sprintf(`"%v" = '%v', `, fields[i], values[i])
	}
	request = strings.TrimSuffix(request, ", ") + fmt.Sprintf(` WHERE "%v" = '%v';`, param, value)

	_, err := DB.Exec(request)

	if err != nil {
		return errors.New(fmt.Sprintf(`Failed to locate "%v", be sure get value and method are valid, and not null`, value)), fiber.StatusNotFound
	}

	return nil, 0
}

func Delete(table string, param string, value string) (error, int) {
	request := fmt.Sprintf(`DELETE FROM "%v" WHERE "%v" = '%v';`, table, param, value)

	_, err := DB.Exec(request)
	if err != nil {
		return errors.New(fmt.Sprintf(`Failed to locate "%v", be sure get value and method are valid, and not null`, value)), fiber.StatusNotFound
	}

	return nil, 0
}
