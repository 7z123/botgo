package testcase

import (
	"testing"

	"github.com/tencent-connect/botgo/dto"
)

func TestChannel(t *testing.T) {
	t.Run("channel list", func(t *testing.T) {
		list, err := api.Channels(ctx, testGuildID)
		if err != nil {
			t.Error(err)
		}
		for _, channel := range list {
			t.Logf("%+v", channel)
		}
		t.Logf(api.TraceID())
	})
	t.Run("create and modify and delete", func(t *testing.T) {
		channel, err := api.PostChannel(ctx, testGuildID, &dto.ChannelValueObject{
			Name:     "机器人创建的频道",
			Type:     dto.ChannelTypeText,
			Position: 0,   // 默认是当前时间戳，如果传递，则要避免和其他频道的 position 重复，否则会报错
			ParentID: "0", // 父ID，正常应该找到一个分组ID，如果传0，就不归属在任何一个分组中
		})
		if err != nil {
			t.Error(err)
		}
		t.Log(channel)
		channelNew, err := api.PatchChannel(ctx, channel.ID, &dto.ChannelValueObject{
			Name: "机器人修改的频道-修改",
		})
		if err != nil {
			t.Error(err)
		}
		t.Log(channelNew)
		if channelNew.Name == channel.Name {
			t.Error("channel name not modified")
		}
		err = api.DeleteChannel(ctx, channelNew.ID)
		if err != nil {
			t.Error(err)
		}
	})
}
