package v1

import (
	"context"

	"github.com/tencent-connect/botgo/dto"
)

// GuildMute 频道禁言
func (o *openAPI) GuildMute(ctx context.Context, guildID string, mute *dto.UpdateGuildMute) error {
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetBody(mute).
		Patch(getURL(guildMuteURI, o.sandbox))
	if err != nil {
		return err
	}
	return nil
}

// MemberMute 频道指定成员禁言
func (o *openAPI) MemberMute(ctx context.Context, guildID, userID string,
	mute *dto.UpdateGuildMute) error {
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("user_id", userID).
		SetBody(mute).
		Patch(getURL(guildMembersMuteURI, o.sandbox))
	if err != nil {
		return err
	}
	return nil
}
