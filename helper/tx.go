package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	//cek apakah eror atau tidak
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
