package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go_bank/db"
	"net/http"
)

// CreateAccountRequest represents the request body for creating an account
// Note the json tags for the fields, which are used for JSON unmarshalling
// Binding is used to validate the request body
type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required,min=1,max=30"`
	Currency string `json:"currency" binding:"required,oneof=EUR USD"`
}

// createAccount creates a new account in the database
func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest

	// Check if the request body is valid (JSON)
	// Besides from unmarshalling, ShouldBindJSON also checks the binding property to validate the body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create a new account
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return the account as a response
	ctx.JSON(http.StatusCreated, account)
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getAccount returns an account from the database by its ID
func (server *Server) getAccount(ctx *gin.Context) {
	var req GetAccountRequest

	// Check if the request parameters are valid (uri params)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get the account from the database
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}

		return
	}

	// Return the account as a response
	ctx.JSON(http.StatusOK, account)
}

type ListAccountRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

// listAccounts returns a list of accounts from the database
func (server *Server) listAccounts(ctx *gin.Context) {
	var req ListAccountRequest

	// Check if the request parameters are valid (query uri params)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get the accounts from the database
	accounts, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: int32((req.PageID - 1) * int64(req.PageSize)),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return the accounts as a response
	ctx.JSON(http.StatusOK, accounts)
}
