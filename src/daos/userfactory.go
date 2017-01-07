package daos

import "log"

func UserFactoryDao(e string) UserDao {
	var dao UserDao
	switch e {
	case "mysql":
		dao = UserImplMysql{}
	default:
		log.Fatalf("Errorr ", e)
		return nil
	}
	return dao
}