package daos

// UserFactoryDao select the sql engine for running
func UserFactoryDao(e string) UserDao {
	var dao UserDao
	switch e {
	case "mysql":
		dao = UserImplMysql{}
	default:
		dao = StaticUserImpl{}
	}
	return dao
}
