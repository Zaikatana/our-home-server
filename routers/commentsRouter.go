package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"our-home-server/db"
)

func getCommentsForItem(c *gin.Context) {
	itemId, err := uuid.Parse(c.Param("itemId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item is invalid"})
		return
	}

	comments, err := db.GetCommentsForItem(itemId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to load comments"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"comments": comments})
}

func addComment(c *gin.Context) {
	var newComment *db.Comment
	if err := c.BindJSON(&newComment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid comment body"})
		return
	}

	addedComment, err := db.AddComment(newComment)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to add comment"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"commentId": addedComment.CommentID})
}

func updateComment(c *gin.Context) {
	var updatedComment *db.Comment
	commentId, err := uuid.Parse(c.Param("commentId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Comment is invalid"})
		return
	}

	if err := c.BindJSON(&updatedComment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid comment body"})
		return
	}

	updatedCommentRes := db.UpdateComment(commentId, updatedComment)
	if updatedCommentRes != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to update comment"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"commentId": commentId})
}

func deleteComment(c *gin.Context) {
	commentId, err := uuid.Parse(c.Param("commentId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Comment is invalid"})
		return
	}

	deleteCommentResult := db.DeleteComment(commentId)
	if deleteCommentResult != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to delete comment"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Comment was deleted"})
}

func InitCommentsRouter(r *gin.Engine) {
	items := r.Group("/api/comments")
	{
		items.GET("/item/:itemId", getCommentsForItem)
		items.DELETE("/:commentId", deleteComment)
		items.POST("/add", addComment)
		items.PUT("/:commentId", updateComment)
	}
}
