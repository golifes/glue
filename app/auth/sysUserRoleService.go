package auth

func findRoleIDByUser(userID int64) ([]int64, error) {
	return findRoleIDByUserID(userID)
}
