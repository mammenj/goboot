package daos

import "github.com/mammenj/goboot/models"

type UserImplMysql struct {
}

func (dao UserImplMysql) Create(u *models.User) error {
	query := "INSERT INTO allusers (name, gender, age) VALUES(?, ?, ?)"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Name, u.Gender, u.Age)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.Id = int(id)
	return nil
}

func (dao UserImplMysql) GetAll() ([]models.User, error) {
	query := "SELECT id, name, gender, age FROM allusers"
	users := make([]models.User, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return users, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.Id, &row.Name, &row.Gender, &row.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, row)
	}
	return users, nil
}

func (dao UserImplMysql) Delete(id int) error {
	query := "DELETE FROM allusers WHERE id = ?"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (dao UserImplMysql) Get(id int) (models.User, error) {
	query := "SELECT id, name, gender, age FROM allusers WHERE id = ?"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return models.User{}, err
	}
	var row models.User
	for rows.Next() {
		err := rows.Scan(&row.Id, &row.Name, &row.Gender, &row.Age)
		if err != nil {
			return models.User{}, err
		}
	}
	return row, err
}

func (dao UserImplMysql) Update(u *models.User) error {
	query := "UPDATE allusers SET name=?, gender=?, age=? WHERE id=?"

	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Name, u.Gender, u.Age, u.Id)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(id)
	return nil
}
