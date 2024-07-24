package base_repo

import (
	customContext "github.com/balireddypraveen/allen/internal/pkg/context"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBaseRepo interface {
	GetTransaction() *gorm.DB
	Rollback(rCtx customContext.ReqCtx, tx *gorm.DB)
	GetFirstRecord(rCtx customContext.ReqCtx, model interface{}, tableName string) error
	GetRecords(rCtx customContext.ReqCtx, model interface{}, filterModel interface{}, tableName string) error
	GetRecordsWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}, filterModel interface{}, tableName string) error
	GetRecordsByCondition(rCtx customContext.ReqCtx, model interface{}, whereClause string, tableName string) error
	Update(rCtx customContext.ReqCtx, model interface{}, tableName string) error
	UpdateWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}) error
	UpdateWhere(reqCtx customContext.ReqCtx, model interface{}, tableName string, conditions interface{}) error
	UpdateWhereWithQuery(rCtx customContext.ReqCtx, model interface{}, tableName, query string, args []interface{}) error
	UpdateWhereWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}, tableName string, conditions interface{}) error
	UpdateWhereReturningRowsAffected(reqCtx customContext.ReqCtx, model interface{}, tableName string, conditions interface{}, throwErrorOnZeroRowsAffected bool) (int64, error)
	Create(rCtx customContext.ReqCtx, model interface{}, tableName string) error
	CreateWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}) error
	GetRecordsWithFilterAndLimitAndOrderByAndOffset(rCtx customContext.ReqCtx, model interface{}, filterModel interface{}, limit int, orderBy string, tableName string, offset int, fields string) error
	GetRecordsWithFilterAndLimitAndOrderByAndWhereClause(rCtx customContext.ReqCtx, model interface{}, filterModel interface{}, limit int, orderBy string, tableName string, whereClause string, fields string) error
	GetRecordsWithFilterAndLimitAndOrderByAndWhereClauseArgs(rCtx customContext.ReqCtx, model, filterModel interface{},
		limit int, orderBy, tableName, whereClause string, args []interface{}, fields string) error
	Delete(rCtx customContext.ReqCtx, model interface{}, tableName string) error
	GetRecordsWithTxnAndWithSkipLock(rCtx customContext.ReqCtx, tx interface{}, model interface{}, filterModel interface{}) error
	Save(rCtx customContext.ReqCtx, value interface{}) (err error)
	SaveWithTxn(rCtx customContext.ReqCtx, tx *gorm.DB, value interface{}) (err error)
	BulkSave(rCtx customContext.ReqCtx, values interface{}, tableName string) (err error)
	BulkSaveWithTxn(rCtx customContext.ReqCtx, tx *gorm.DB, values interface{}, tableName string) (err error)
	FirstOrCreate(rCtx customContext.ReqCtx, model interface{}, tableName string) error
	FirstOrCreateWithFilters(rCtx customContext.ReqCtx, model, filterModel, attrs interface{}) error
	FirstOrCreateWithTxnAndFilters(rCtx customContext.ReqCtx, txn *gorm.DB, model, filterModel, attrs interface{}) error
	CreateOrUpdateWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}, uniqueConstraint []clause.Column) error
	GetSumOfColumn(rCtx customContext.ReqCtx, tableName string, columnName string, whereClause string, sum *decimal.Decimal) error
}
