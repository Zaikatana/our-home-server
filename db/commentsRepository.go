package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentID uuid.UUID `json:"commentId" gorm:"primaryKey"`
	ItemID    uuid.UUID `json:"itemId"`
	Item      Item      `json:"-"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetCommentsForItem(itemId uuid.UUID) ([]Comment, error) {
	var comments []Comment

	res := db.Preload("Item").Find(&comments, "item_id = ?", itemId)
	if res.Error != nil {
		return nil, fmt.Errorf("error fetching comments for item: %s", itemId)
	}

	return comments, nil
}

func AddComment(comment *Comment) (*Comment, error) {
	comment.CommentID = uuid.New()
	comment.CreatedAt = time.Now().UTC()
	comment.UpdatedAt = time.Now().UTC()

	res := db.Create(comment)
	if res.Error != nil {
		return nil, res.Error
	}

	return comment, nil
}

func UpdateComment(commentId uuid.UUID, updateComment *Comment) error {
	updateComment.UpdatedAt = time.Now().UTC()
	res := db.Model(&Comment{}).Where("comment_id = ?", commentId).Updates(updateComment)
	if res.RowsAffected == 0 {
		return errors.New("comment not updated")
	}

	return nil
}

func DeleteComment(commentId uuid.UUID) error {
	var deletedComment Comment
	res := db.Delete(&deletedComment, "comment_id = ?", commentId)
	if res.RowsAffected == 0 {
		return errors.New("comment not deleted")
	}

	return nil
}
