package post

import "HexMaster/api/handler/group"

type Post struct {
	Id        string       `json:"id" db:"id"`
	CreatedAt string       `json:"created-at" db:"createdAt"`
	Creator   group.Member `json:"creator" db:"creator"`
	Group     group.Group  `json:"group" db:"group"`
	Title     string       `json:"title" db:"title"`
	Content   string       `json:"content" db:"content"`
	Type      ContentType  `json:"type" db:"type"`
	Parent    string       `json:"parent" db:"parent"`
}

type ContentType string

const (
	CONTENT_POST    ContentType = "post"
	CONTENT_COMMENT ContentType = "comment"
	CONTENT_LIKE    ContentType = "like"
)

func (m ContentType) String() string {
	return string(m)
}

func contentTypeFromString(role string) (ContentType, bool) {
	switch role {
	case CONTENT_POST.String():
		return CONTENT_POST, true
	case CONTENT_COMMENT.String():
		return CONTENT_COMMENT, true
	case CONTENT_LIKE.String():
		return CONTENT_LIKE, true
	default:
		return ContentType(""), false
	}
}
