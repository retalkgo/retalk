package store

import (
	"fmt"

	"github.com/retalkgo/retalk/internal/gravatar"
	"github.com/retalkgo/retalk/internal/model"
	"gorm.io/gorm"
)

type CommentStore struct {
	db         *gorm.DB
	storeCache *StoreCache
}

func NewCommentsStore(db *gorm.DB, storeCache *StoreCache) *CommentStore {
	return &CommentStore{
		db:         db,
		storeCache: storeCache,
	}
}

func (s *CommentStore) Create(comment *model.Comment) error {
	err := s.db.Create(comment).Error
	if err != nil {
		return err
	}

	s.storeCache.CommentCacheSet(comment)

	return nil
}

func (s *CommentStore) FindByID(id uint) (*model.Comment, error) {
	comment, err := QueryWithCache(s.storeCache, fmt.Sprintf(CommentByIDKey, id), func() (*model.Comment, error) {
		var comment model.Comment
		err := s.db.First(&comment, id).Error

		return &comment, err
	})

	return comment, err
}

func (s *CommentStore) FindByDomainAndPath(domain string, path string, pageIndex int, pageSize int) ([]*model.Comment, int64, error) {
	var comments []*model.Comment

	var count int64
	offset := (pageIndex - 1) * pageSize

	err := s.db.Model(&model.Comment{}).
		Joins("JOIN sites ON sites.id = comments.site_id").
		Where("sites.domain = ? AND comments.path = ?", domain, path).
		Preload("User").
		Limit(pageSize).
		Order("comments.created_at DESC").
		Offset(offset).
		Find(&comments).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, err
	}

	return comments, count, nil
}

func (s *CommentStore) Update(comment *model.Comment) error {
	err := s.db.Save(comment).Error
	if err != nil {
		return err
	}

	s.storeCache.CommentCacheSet(comment)

	return nil
}

func (s *CommentStore) Delete(comment *model.Comment) error {
	err := s.db.Delete(comment).Error
	if err != nil {
		return err
	}

	s.storeCache.CommentCacheDelete(comment)

	return nil
}

func (s *CommentStore) Cook(comment *model.Comment) (*model.CookedComment, error) {
	var userID *uint
	var nickname string
	var email string
	var website string

	if comment.UserID != nil {
		userID = comment.UserID
		nickname = comment.User.Nickname
		email = comment.User.HashedEmail
		website = comment.User.Website
	}

	if comment.UserID == nil {
		nickname = comment.GuestName
		email = comment.GuestHashedEmail
		website = comment.GuestWebsite
	}

	appConfig, err := AppConfig.GetAll()
	if err != nil {
		return nil, err
	}

	avatar := gravatar.GetGravatarURL(appConfig.Gravatar.BaseURL, email)

	cookedComment := &model.CookedComment{
		BaseModel:   comment.BaseModel,
		UserID:      userID,
		Content:     comment.Content,
		HashedEmail: email,
		NickName:    nickname,
		Website:     website,
		Avatar:      avatar,
	}

	return cookedComment, nil
}
