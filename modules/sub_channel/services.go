package sub_channel

import (
	"context"
	"github.com/zakariawahyu/go-echo-news/entity"
)

type SubChannelServices interface {
	GetAllSubChannel(ctx context.Context) (subChannels []entity.SubChannelResponse)
	GetSubChannelBySlugOrId(ctx context.Context, slug string) entity.SubChannelResponse
}
