package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetAnnouncementResponse struct {
	Id         string    `json:"id"`          // uuid
	Type       int       `json:"type"`        // 公告類型(1: 系統, 2: 活動)
	Subject    string    `json:"subject"`     // 主旨
	Content    string    `json:"content"`     // 內文
	CreateTime time.Time `json:"create_time"` // 創建時間
	UpdateTime time.Time `json:"update_time"` // 更新時間
}

func NewEmptyGetAnnouncementResponse() *GetAnnouncementResponse {
	return &GetAnnouncementResponse{}
}

/* ----------------------------------------------------------- */

type GetAnnouncementRequest struct {
	Id string `json:"id"`
}

func NewEmptyGetAnnouncementRequest() *GetAnnouncementRequest {
	return &GetAnnouncementRequest{}
}

func (p *GetAnnouncementRequest) CheckParams() bool {
	return p.Id != ""
}

/* ----------------------------------------------------------- */

type CreateAnnouncementRequest struct {
	Type    int    `json:"type"`    // 公告類型(1: 系統, 2: 活動)
	Subject string `json:"subject"` // 主旨(20長度)
	Content string `json:"content"` // 內文(800長度)
}

func NewEmptyCreateAnnouncementRequest() *CreateAnnouncementRequest {
	return &CreateAnnouncementRequest{}
}

func (p *CreateAnnouncementRequest) CheckParams() bool {
	if p.Type <= definition.ANNOUNCEMENT_TYPE_NONE || p.Type >= definition.ANNOUNCEMENT_TYPE_COUNT ||
		p.Subject == "" || utils.WordLength(p.Subject) > 20 ||
		p.Content == "" || utils.WordLength(p.Content) > 800 {
		return false
	}
	return true
}

/* ----------------------------------------------------------- */

type UpdateAnnouncementRequest struct {
	Id      string `json:"id"`      // uuid
	Type    int    `json:"type"`    // 公告類型(1: 系統, 2: 活動)
	Subject string `json:"subject"` // 主旨(20長度)
	Content string `json:"content"` // 內文(800長度)
}

func NewEmptyUpdateAnnouncementRequest() *UpdateAnnouncementRequest {
	return &UpdateAnnouncementRequest{}
}

func (p *UpdateAnnouncementRequest) CheckParams() bool {
	if p.Type <= definition.ANNOUNCEMENT_TYPE_NONE || p.Type >= definition.ANNOUNCEMENT_TYPE_COUNT ||
		p.Subject == "" || utils.WordLength(p.Subject) > 20 ||
		p.Content == "" || utils.WordLength(p.Content) > 800 ||
		p.Id == "" {
		return false
	}
	return true
}

/* ----------------------------------------------------------- */

type DeleteAnnouncementRequest struct {
	Id string `json:"id"` // uuid
}

func NewEmptyDeleteAnnouncementRequest() *DeleteAnnouncementRequest {
	return &DeleteAnnouncementRequest{}
}

func (p *DeleteAnnouncementRequest) CheckParams() bool {
	return p.Id != ""
}
