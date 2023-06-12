package entity

import (
	"github.com/uptrace/bun"
	"reflect"
	"time"
)

type Content struct {
	bun.BaseModel    `bun:"table:contents_new"`
	ID               int64                   `bun:",pk" json:"id"`
	Type             string                  `json:"type"`
	TypeID           *int64                  `json:"type_id"`
	TypeChildID      *int64                  `json:"type_child_id"`
	SuplemenID       *int64                  `json:"suplemen_id"`
	UpperTitle       string                  `bun:"uppertitle" json:"upper_title"`
	Title            string                  `json:"title"`
	Slug             string                  `json:"slug"`
	Excerpt          string                  `json:"excerpt"`
	Content          string                  `json:"content"`
	Caption          *string                 `json:"caption"`
	Image            *string                 `json:"image"`
	Thumbnail        *string                 `json:"thumbnail"`
	HeadlineType     int64                   `json:"headline_type"`
	IsActive         bool                    `json:"is_active"`
	IsFeatured       bool                    `json:"is_featured"`
	IsEditorChoice   bool                    `json:"is_editor_choice"`
	IsNational       bool                    `json:"is_national"`
	Token            string                  `json:"token"`
	AdsPosition      int                     `json:"ads_position"`
	AdsExpiredDate   time.Time               `json:"ads_expired_date"`
	UrlVideo         *string                 `json:"url_video"`
	IsPopular        bool                    `json:"is_popular"`
	IsAdult          bool                    `json:"is_adult"`
	IsComment        bool                    `json:"is_comment"`
	LocationCity     string                  `json:"location_city"`
	LocationDistrict string                  `bun:"location_citydistrict" json:"location_district"`
	LocationSuburb   string                  `json:"location_suburb"`
	CreatedBy        int64                   `json:"created_by"`
	User             *UserResponse           `bun:"rel:has-one,join:created_by=id" json:"user"`
	Region           *RegionResponse         `bun:"rel:has-one,join:type_id=id" json:"region"`
	Channel          *ContentChannelResponse `bun:"rel:has-one,join:type_id=id" json:"channel"`
	SubChannel       *SubChannelResponse     `bun:"rel:has-one,join:type_child_id=id" json:"sub_channel"`
	Tags             []Tag                   `bun:"m2m:content_has_tags,join:Content=Tag" json:"tags"`
	Topics           []Topic                 `bun:"m2m:content_has_topics,join:Content=Topic" json:"topics"`
	Reporters        []Reporter              `bun:"m2m:content_has_reporters,join:Content=Reporter" json:"reporters"`
	SubPhotos        []*SubPhoto             `bun:"rel:has-many,join:id=content_id" json:"sub_photos"`
	PublishedDate    time.Time               `json:"published_date"`
	CreatedAt        time.Time               `bun:"created" json:"created_at"`
	UpdatedAt        bun.NullTime            `bun:"modified" json:"updated_at"`
}

type ContentResponse struct {
	bun.BaseModel `bun:"table:contents_new"`
	ID            int64                   `bun:",pk" json:"id"`
	Type          string                  `json:"type"`
	TypeID        *int64                  `json:"type_id"`
	TypeChildID   *int64                  `json:"type_child_id"`
	SuplemenID    *int64                  `json:"suplemen_id"`
	UpperTitle    string                  `bun:"uppertitle" json:"upper_title"`
	Title         string                  `json:"title"`
	Slug          string                  `json:"slug"`
	Excerpt       string                  `json:"excerpt"`
	Content       string                  `json:"content"`
	Caption       *string                 `json:"caption"`
	Image         *string                 `json:"image"`
	Thumbnail     *string                 `json:"thumbnail"`
	UrlVideo      *string                 `json:"url_video"`
	IsPopular     bool                    `json:"is_popular"`
	IsAdult       bool                    `json:"is_adult"`
	IsComment     bool                    `json:"is_comment"`
	CreatedBy     int64                   `json:"created_by"`
	User          *UserResponse           `bun:"rel:has-one,join:created_by=id" json:"user"`
	Region        *RegionResponse         `bun:"rel:has-one,join:type_id=id" json:"region"`
	Channel       *ContentChannelResponse `bun:"rel:has-one,join:type_id=id" json:"channel"`
	SubChannel    *SubChannelResponse     `bun:"rel:has-one,join:type_child_id=id" json:"sub_channel"`
	Tags          []Tag                   `bun:"m2m:content_has_tags,join:Content=Tag" json:"tags"`
	Topics        []Topic                 `bun:"m2m:content_has_topics,join:Content=Topic" json:"topics"`
	Reporters     []Reporter              `bun:"m2m:content_has_reporters,join:Content=Reporter" json:"reporters"`
	SubPhotos     []*SubPhoto             `bun:"rel:has-many,join:id=content_id" json:"sub_photos"`
	PublishedDate time.Time               `json:"published_date"`
	UpdatedAt     bun.NullTime            `bun:"modified" json:"updated_at"`
}

func NewContentResponse(content *Content) ContentResponse {
	return ContentResponse{
		ID:            content.ID,
		Type:          content.Type,
		TypeID:        content.TypeID,
		TypeChildID:   content.TypeChildID,
		SuplemenID:    content.SuplemenID,
		UpperTitle:    content.UpperTitle,
		Title:         content.Title,
		Slug:          content.Slug,
		Excerpt:       content.Excerpt,
		Content:       content.Content,
		Caption:       content.Caption,
		Image:         content.Image,
		Thumbnail:     content.Thumbnail,
		UrlVideo:      content.UrlVideo,
		IsPopular:     content.IsPopular,
		IsAdult:       content.IsAdult,
		IsComment:     content.IsComment,
		CreatedBy:     content.CreatedBy,
		User:          content.User,
		Region:        content.Region,
		Channel:       content.Channel,
		SubChannel:    content.SubChannel,
		Tags:          content.Tags,
		Topics:        content.Topics,
		Reporters:     content.Reporters,
		SubPhotos:     content.SubPhotos,
		PublishedDate: content.PublishedDate,
		UpdatedAt:     content.UpdatedAt,
	}
}

func (c *Content) IsEmpty() bool {
	return reflect.DeepEqual(c, Content{})
}

func (c *Content) TagNameArray() []string {
	var tags []string

	for _, tag := range c.Tags {
		tags = append(tags, tag.Name)
	}

	return tags
}
