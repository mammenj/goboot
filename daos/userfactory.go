package daos

// UserFactoryDao select the sql engine for running
import "log"

func UserFactoryDao(e string) UserDao {
	var dao UserDao
	switch e {
	case "mysql":
		dao = UserImplMysql{}
	default:
		dao = StaticUserImpl{}
		log.Fatalf("Errorr %s", e)
		return nil
	}
	return dao
}
