package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// deleteCallback used to delete data from database or set deleted_at to current time (when using with soft delete)
// ref: https://github.com/jinzhu/gorm/blob/master/callback_delete.go
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		if !scope.Search.Unscoped {
			if isDeletedField, ok := scope.FieldByName("IsDeleted"); ok {
				if deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt"); hasDeletedAtField {
					scope.Raw(fmt.Sprintf(
						"UPDATE %v SET %v=%v,%v=%v%v%v",
						scope.QuotedTableName(),
						scope.Quote(isDeletedField.DBName),
						scope.AddToVars(1),
						scope.Quote(deletedAtField.DBName),
						scope.AddToVars(gorm.NowFunc()),
						addExtraSpaceIfExist(scope.CombinedConditionSql()),
						addExtraSpaceIfExist(extraOption),
					)).Exec()
				} else {
					scope.Raw(fmt.Sprintf(
						"UPDATE %v SET %v=%v%v%v",
						scope.QuotedTableName(),
						scope.Quote(isDeletedField.DBName),
						scope.AddToVars(1),
						addExtraSpaceIfExist(scope.CombinedConditionSql()),
						addExtraSpaceIfExist(extraOption),
					)).Exec()
				}

				return
			}
		}

		scope.Raw(fmt.Sprintf(
			"DELETE FROM %v%v%v",
			scope.QuotedTableName(),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)).Exec()
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
