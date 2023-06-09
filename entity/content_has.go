package entity

type ContentHasTag struct {
	ID        int64    `bun:",pk"`
	TagID     int64    `bun:",pk"`
	Tag       *Tag     `bun:"rel:belongs-to,join:tag_id=id"`
	ContentID int64    `bun:",pk"`
	Content   *Content `bun:"rel:belongs-to,join:content_id=id"`
}

type ContentHasTopic struct {
	ID        int64    `bun:",pk"`
	TopicID   int64    `bun:",pk"`
	Topic     *Topic   `bun:"rel:belongs-to,join:topic_id=id"`
	ContentID int64    `bun:",pk"`
	Content   *Content `bun:"rel:belongs-to,join:content_id=id"`
}

type ContentHasReporter struct {
	ID         int64     `bun:",pk"`
	ReporterID int64     `bun:",pk"`
	Reporter   *Reporter `bun:"rel:belongs-to,join:reporter_id=id"`
	ContentID  int64     `bun:",pk"`
	Content    *Content  `bun:"rel:belongs-to,join:content_id=id"`
}
