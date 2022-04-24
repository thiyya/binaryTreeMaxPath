package handlers

import (
	"binaryTreeMaxPath/model"
	"binaryTreeMaxPath/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleBinaryTreeMaxPath(ctx *gin.Context) {
	var input model.Input
	if err := ctx.BindJSON(&input); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	root := utils.InputToBinaryTree(input)
	utils.FindMaxPath(root)
	max := root.Left.MaxPath + root.Right.MaxPath + root.Value
	ctx.IndentedJSON(http.StatusOK, model.Result{MaxPathValue: max})
}
