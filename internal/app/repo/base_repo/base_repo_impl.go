package base_repo

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"

	customContext "github.com/balireddypraveen/allen/internal/pkg/context"

	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepo struct {
	db *gorm.DB
}

func NewBaseRepo() *BaseRepo {
	return &BaseRepo{}
}

// NewRelicDbWrapper What is explicit use of this?
func (baseRepo *BaseRepo) NewRelicDbWrapper(rCtx customContext.ReqCtx) (*gorm.DB, *newrelic.Segment) {
	gormTransactionTrace := rCtx.NewRelicTxn
	gormTransactionContext := newrelic.NewContext(rCtx.Context, gormTransactionTrace)
	tracedDB := baseRepo.db.WithContext(gormTransactionContext)
	gormTransactionTraceSegment := gormTransactionTrace.StartSegment("DB Query")
	return tracedDB, gormTransactionTraceSegment
}

func (baseRepo *BaseRepo) NewRelicTxn(rCtx customContext.ReqCtx, txn *gorm.DB) (*gorm.DB, *newrelic.Segment) {
	gormTransactionTrace := rCtx.NewRelicTxn
	gormTransactionContext := newrelic.NewContext(rCtx.Context, gormTransactionTrace)
	tracedDB := txn.WithContext(gormTransactionContext)
	gormTransactionTraceSegment := gormTransactionTrace.StartSegment("DB Query")
	return tracedDB, gormTransactionTraceSegment
}

func (baseRepo *BaseRepo) GetFirstRecord(rCtx customContext.ReqCtx, model interface{}, tableName string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	if tx := db.Table(tableName).First(model); tx.Error != nil {
		rCtx.Log.Errorf("failed to get entity error: " + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) GetRecords(rCtx customContext.ReqCtx, model interface{}, filterModel interface{}, tableName string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	tx := db.Table(tableName).Find(model, filterModel)
	if tx.Error != nil {
		rCtx.Log.Errorf("failed to get records error: " + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) GetRecordsWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}, filterModel interface{}, tableName string) error {
	txn, seg := baseRepo.NewRelicTxn(rCtx, txn)
	defer seg.End()

	tx := txn.Table(tableName).Where(filterModel).Find(model)
	if tx.Error != nil {
		rCtx.Log.Errorf("failed to get records error:" + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) GetRecordsByCondition(rCtx customContext.ReqCtx, model interface{}, whereClause string, tableName string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	tx := db.Table(tableName).Where(whereClause).Find(model)
	if tx.Error != nil {
		rCtx.Log.Errorf("failed to get records error:" + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) Update(rCtx customContext.ReqCtx, model interface{}, tableName string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	if tx := db.Table(tableName).Updates(model); tx.Error != nil {
		rCtx.Log.Error("failed to update entity error: " + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) UpdateWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}) error {
	txn, seg := baseRepo.NewRelicTxn(rCtx, txn)
	defer seg.End()

	if tx := txn.Updates(model); tx.Error != nil {
		rCtx.Log.Error("failed to update entity error:" + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) UpdateWhereWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}, tableName string, conditions interface{}) error {
	txn, seg := baseRepo.NewRelicTxn(rCtx, txn)
	defer seg.End()

	result := txn.Table(tableName).Where(conditions).Updates(model)
	if result.Error != nil {
		rCtx.Log.Error("failed to update entity error: " + result.Error.Error())
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows to update, please check the filter query: %v on table: %v", conditions, tableName)
	}

	return nil
}

func (baseRepo *BaseRepo) UpdateWhere(rCtx customContext.ReqCtx, model interface{}, tableName string, conditions interface{}) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	result := db.Table(tableName).Where(conditions).Updates(model)
	if result.Error != nil {
		rCtx.Log.Error("failed to update entity error: " + result.Error.Error())
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows to update, please check the filter query: %v on table: %v", conditions, tableName)
	}

	return nil
}

func (baseRepo *BaseRepo) UpdateWhereWithQuery(rCtx customContext.ReqCtx, model interface{}, tableName, query string, args []interface{}) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	result := db.Table(tableName).Where(query, args...).Updates(model)
	if result.Error != nil {
		rCtx.Log.Error("failed to update entity error: " + result.Error.Error())
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows to update, please check the filter query: %v args %+v on table: %v", query, args, tableName)
	}

	return nil
}

func (baseRepo *BaseRepo) UpdateWhereReturningRowsAffected(rCtx customContext.ReqCtx, model interface{}, tableName string, conditions interface{}, throwErrorOnZeroRowsAffected bool) (int64, error) {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	result := db.Table(tableName).Where(conditions).Updates(model)
	if result.Error != nil {
		rCtx.Log.Error("failed to update entity error:" + result.Error.Error())
		return 0, result.Error
	}

	if throwErrorOnZeroRowsAffected && result.RowsAffected == 0 {
		return 0, fmt.Errorf("no rows to update, please check the filter query: %v on table: %v", conditions, tableName)
	}

	return result.RowsAffected, nil
}

func (baseRepo *BaseRepo) Create(rCtx customContext.ReqCtx, model interface{}, tableName string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	if tx := db.Table(tableName).Create(model); tx.Error != nil {
		rCtx.Log.Error("failed to create entity error: " + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) CreateWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}) error {
	txn, seg := baseRepo.NewRelicTxn(rCtx, txn)
	defer seg.End()

	if tx := txn.Create(model); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) GetRecordsWithFilterAndLimitAndOrderByAndOffset(rCtx customContext.ReqCtx, model interface{}, filterModel interface{}, limit int, orderBy string, tableName string, offset int, fields string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	if fields == "" {
		fields = "*"
	}
	tx := db.Table(tableName).
		Select(fields).
		Order(orderBy).
		Offset(offset).
		Limit(limit).
		Find(model, filterModel)

	if tx.Error != nil {
		rCtx.Log.Error("failed to get records error: " + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) GetRecordsWithFilterAndLimitAndOrderByAndWhereClause(rCtx customContext.ReqCtx, model interface{}, filterModel interface{}, limit int, orderBy string, tableName string, whereClause string, fields string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	if fields == "" {
		fields = "*"
	}
	tx := db.Table(tableName).
		Select(fields).
		Order(orderBy).
		Where(whereClause).
		Limit(limit).
		Find(model, filterModel)

	if tx.Error != nil {
		rCtx.Log.Error("failed to get records error: " + tx.Error.Error())
		return tx.Error
	}

	//rCtx.Log.InfofCf("GetRecordsWithFilterAndLimitAndOrderByAndWhereClause transaction Statement:- %+v, tx.Statement.TableExpr.SQL:- %+v", tx.Statement, tx.Statement.SQL.String())

	return nil
}

func (baseRepo *BaseRepo) GetRecordsWithFilterAndLimitAndOrderByAndWhereClauseArgs(rCtx customContext.ReqCtx, model, filterModel interface{},
	limit int, orderBy, tableName, whereClause string, args []interface{}, fields string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	if fields == "" {
		fields = "*"
	}
	tx := db.Table(tableName).
		Select(fields).
		Order(orderBy).
		Where(whereClause, args...)
	if limit != 0 {
		tx = tx.Limit(limit)
	}
	tx = tx.Find(model, filterModel)
	if tx.Error != nil {
		rCtx.Log.Error("failed to get records error: " + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) Delete(rCtx customContext.ReqCtx, model interface{}, tableName string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	//Passing nil here as we want to delete based on condition filter not value
	if tx := db.Table(tableName).Unscoped().Delete(nil, model); tx.Error != nil {
		rCtx.Log.Errorf("failed to delete entity error: " + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) GetTransaction() *gorm.DB {
	return baseRepo.db.Begin()
}

func (baseRepo *BaseRepo) Rollback(rCtx customContext.ReqCtx, tx *gorm.DB) {
	if tx != nil {
		tx.Rollback()
	}
}

func (baseRepo *BaseRepo) GetRecordsWithTxnAndWithSkipLock(rCtx customContext.ReqCtx, tx interface{}, model interface{}, filterModel interface{}) error {
	db := tx.(*gorm.DB)
	db, seg := baseRepo.NewRelicTxn(rCtx, db)
	defer seg.End()

	result := db.Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}).Find(model, filterModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (baseRepo *BaseRepo) Save(rCtx customContext.ReqCtx, value interface{}) (err error) {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	tx := db.Save(value)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) SaveWithTxn(rCtx customContext.ReqCtx, tx *gorm.DB, value interface{}) (err error) {
	tx, seg := baseRepo.NewRelicTxn(rCtx, tx)
	defer seg.End()

	tx.Save(value)
	commitError := tx.Commit().Error
	if commitError != nil {
		rCtx.Log.Errorf("Commit error: %v", commitError)
		return commitError
	}

	return nil
}

func (baseRepo *BaseRepo) BulkSave(rCtx customContext.ReqCtx, values interface{}, tableName string) (err error) {
	// Begin a database transaction
	txn := baseRepo.GetTransaction()
	defer baseRepo.Rollback(rCtx, txn)
	return baseRepo.BulkSaveWithTxn(rCtx, txn, values, tableName)
}

func (baseRepo *BaseRepo) BulkSaveWithTxn(rCtx customContext.ReqCtx, tx *gorm.DB, values interface{}, tableName string) (err error) {
	tx, seg := baseRepo.NewRelicTxn(rCtx, tx)
	defer seg.End()

	for _, value := range values.([]interface{}) {
		tx.Table(tableName).Save(value)
	}

	commitError := tx.Commit().Error
	if commitError != nil {
		rCtx.Log.Errorf("Commit error: %v", commitError)
		return commitError
	}

	return nil
}

func (baseRepo *BaseRepo) FirstOrCreate(rCtx customContext.ReqCtx, model interface{}, tableName string) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	if tx := db.Table(tableName).FirstOrCreate(model); tx.Error != nil {
		rCtx.Log.Error("failed to create entity error: " + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) FirstOrCreateWithFilters(rCtx customContext.ReqCtx, model, filterModel, attrs interface{}) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	tx := db.Where(filterModel).Attrs(attrs).FirstOrCreate(model)
	if tx.Error != nil {
		rCtx.Log.Error("failed to create entity error:" + tx.Error.Error())
		return tx.Error
	}
	//rCtx.Log.InfofCf("statement %+v", tx)

	return nil
}

func (baseRepo *BaseRepo) FirstOrCreateWithTxnAndFilters(rCtx customContext.ReqCtx, txn *gorm.DB, model, filterModel, attrs interface{}) error {
	txn, seg := baseRepo.NewRelicTxn(rCtx, txn)
	defer seg.End()

	if tx := txn.Where(filterModel).Attrs(attrs).FirstOrCreate(model); tx.Error != nil {
		rCtx.Log.Error("failed to create entity error:" + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) CreateOrUpdateWithTxn(rCtx customContext.ReqCtx, txn *gorm.DB, model interface{}, uniqueConstraint []clause.Column) error {
	txn, seg := baseRepo.NewRelicTxn(rCtx, txn)
	defer seg.End()

	if tx := txn.Clauses(clause.OnConflict{
		UpdateAll: true,
		Columns:   uniqueConstraint,
	}).Create(model); tx.Error != nil {
		rCtx.Log.Error("failed to create entity error:" + tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (baseRepo *BaseRepo) GetSumOfColumn(rCtx customContext.ReqCtx, tableName string, columnName string, whereClause string, sum *decimal.Decimal) error {
	db, seg := baseRepo.NewRelicDbWrapper(rCtx)
	defer seg.End()

	var dbSum sql.NullFloat64
	err := db.Table(tableName).Select(fmt.Sprintf("sum(%s)", columnName)).Where(whereClause).Scan(&dbSum).Error
	//rCtx.Log.InfofCf("error %+v", err.Error())
	if err != nil {
		return err
	}

	// Convert to decimal
	if dbSum.Valid {
		*sum = decimal.NewFromFloat(dbSum.Float64)
	} else {
		*sum = decimal.Zero
	}

	return nil
}

func (baseRepo *BaseRepo) SetDb(db *gorm.DB) {
	baseRepo.db = db
}
