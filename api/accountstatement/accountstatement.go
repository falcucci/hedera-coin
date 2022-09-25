package accountstatement

import (
	"net/http"

	"github.com/falcucci/maga-coin-api/dtos"
	"github.com/falcucci/maga-coin-api/models"
	"github.com/falcucci/maga-coin-api/utils/pagination"
	"github.com/falcucci/maga-coin-api/utils/response"
	"github.com/jinzhu/gorm"

	utils "github.com/falcucci/maga-coin-api/utils/database"
)

const (
	// IN : Input coin transaction
	IN = 1
	// OUT : Output coin transaction
	OUT = 2
)

// GetAccountStatement : Get accountStatement of transactions
func GetAccountStatement(w http.ResponseWriter, r *http.Request) {
	limit, offset := pagination.Pagination(r)
	var accountStatements []models.AccountStatement
	if err := utils.DB.Limit(limit).Offset(offset).Find(&accountStatements).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.GenerateHTTPResponse(w, http.StatusNoContent, nil)
			return
		}
		response.GenerateHTTPResponse(w, http.StatusInternalServerError, response.GenerateErrorResponse(response.InternalServerError,
			"Error at Find accountStatement", "Was encountered an error when processing your request. We apologize for the inconvenience."))
		return
	}

	var accountStatementsDto []dtos.AccountStatement
	for _, item := range accountStatements {
		var accountStatement dtos.AccountStatement
		accountStatement.MapModelToDto(item)
		accountStatementsDto = append(accountStatementsDto, accountStatement)
	}

	response.GenerateHTTPResponse(
		w, http.StatusOK,
		response.GenerateSuccessResponse(accountStatementsDto, limit, offset, len(accountStatementsDto)))
}

// CreateAccountStatement : create accountStatement for a transaction
func CreateAccountStatement(accountStatement models.AccountStatement) (*models.AccountStatement, error) {
	if err := utils.DB.Create(&accountStatement).Error; err != nil {
		return nil, err
	}
	return &accountStatement, nil
}
