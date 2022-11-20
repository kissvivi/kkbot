package kkbot

// Type defines the message types.
const (
	TypeRoom     = "room"
	TypeCreator  = "creator"
	TypeGift     = "gift"
	TypeMessage  = "message"
	TypeNotify   = "notify"
	TypeMember   = "member"
	TypeChannel  = "channel"
	TypeQuestion = "question"
	TypeNoble    = "noble"
	TypePK       = "pk"
	TypeSuperFan = "super_fan"
)

// Event defines the message events.
const (
	EventSend         = "send"         // gift send.
	EventNew          = "new"          // new message received.
	EventStatistic    = "statistics"   // statistics of the live room.
	EventJoin         = "join"         // connect to the live room channel.
	EventJoinQueue    = "join_queue"   // members join the live room.
	EventFollowed     = "followed"     // user followed the room creator.
	EventOpen         = "open"         // the live room opened.
	EventClose        = "close"        // the live room closed.
	EventNewRank      = "new_rank"     // the new rank information of the live room.
	EventLeave        = "leave"        // user leaved the live room.
	EventAddAdmin     = "add_admin"    // add a room admin.
	EventRemoveAdmin  = "remove_admin" // Remove a room admin.
	EventRemoveMute   = "remove_mute"  // unmute a user in the live room.
	EventAsk          = "ask"          // ask a question.
	EventAnswer       = "answer"       // answer a question.
	EventConnect      = "connect"      // connect to the live room.
	EventUpdate       = "update"
	EventStart        = "start"
	EventFinish       = "finish"
	EventLastHourRank = "last_hour_rank" // the last hour rank information of the live room.
	EventRenewal      = "renewal"        // 续费（超粉、贵族）事件
	EventRegistration = "registration"   //
	EventHorn         = "horn"           // horn message
)

const (
	TitleLevel       = "level"        // 用户等级
	TitleNoble       = "noble"        // 贵族
	TitleMedal       = "medal"        // 粉丝牌
	TitleStaff       = "staff"        // 猫耳职员
	TitleBadge       = "badge"        // 徽章
	TitleAvatarFrame = "avatar_frame" // 头像框
	TitleHighness    = "highness"     // 上神贵族
)

type (
	// FmMessage represents the Websocket message from the live room.
	FmMessage struct {
		Type       string        `json:"type"`
		Event      string        `json:"event"`
		NotifyType string        `json:"notify_type"`
		RoomID     int64         `json:"room_id"`
		Message    string        `json:"message"`
		MessageID  string        `json:"msg_id"`
		User       FmUser        `json:"user"`
		Queue      []FmJoin      `json:"queue"`
		Noble      *fmNoble      `json:"noble"`
		SuperFan   *fmSuperFan   `json:"super_fan"`
		Info       *fmInfo       `json:"info"`
		Gift       *FmGift       `json:"gift"`
		Lucky      *FmGift       `json:"lucky"`
		Target     *FmTarget     `json:"target"`
		Statistics *fmStatistics `json:"statistics"`
		PK         *fmPK         `json:"pk"`
		Question   *FmQuestion   `json:"question"`
	}

	fmInfo struct {
		Room struct {
			Status struct {
				Open int64 `json:"open"`
			} `json:"status"`
		} `json:"room"`
	}

	// FmUser represents the information of a user.
	FmUser struct {
		IconUrl  string    `json:"iconurl"`
		Titles   []FmTitle `json:"titles"`
		UserID   int64     `json:"user_id"`
		Username string    `json:"username"`
	}

	// FmJoin represents basic information of the user who is joining.
	FmJoin struct {
		Contribution int64     `json:"contribution"`
		IconUrl      string    `json:"iconurl"`
		Titles       []FmTitle `json:"titles"`
		UserID       int64     `json:"user_id"`
		Username     string    `json:"username"`
	}

	FmTitle struct {
		Level int64  `json:"level"`
		Name  string `json:"name"`
		Type  string `json:"type"`
		Color string `json:"color"`
	}

	// FmGift represents the information of gift.
	FmGift struct {
		GiftID       int64  `json:"gift_id"`
		Name         string `json:"name"`
		Price        int64  `json:"price"`
		Number       int64  `json:"num"`
		EffectURL    string `json:"effect_url"`
		WebEffectURL string `json:"web_effect_url"`
	}

	FmTarget struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}

	fmPK struct {
		ID        string `json:"pk_id"`
		Status    int64  `json:"status"`
		StartTime int64  `json:"start_time"`
		Result    int64  `json:"result"` // PK 结果，0 -> 失败，1 -> 胜利
	}

	FmQuestion struct {
		Username string `json:"username"`
		Question string `json:"question"`
		Price    int64  `json:"price"`
	}

	fmNoble struct {
		Name         string `json:"name"`
		Level        int64  `json:"level"`
		Status       int64  `json:"status"`
		Price        int64  `json:"price"`
		Contribution int64  `json:"contribution"`
	}

	fmSuperFan struct {
		Num int64 `json:"num"`
	}
)

type (
	FmResp struct {
		Code int64  `json:"code"`
		Info FmInfo `json:"info"`
	}

	FmInfo struct {
		Creator    fmCreator `json:"creator"`     // 主播
		Room       fmRoom    `json:"room"`        // 直播间
		User       FmUser    `json:"user"`        // 用户
		Medal      FmMedal   `json:"medal"`       // 粉丝牌
		OwnerCount int64     `json:"owner_count"` //
	}

	FmMedal struct {
		CreatorID   int64  `json:"creator_id"`
		CreatorName string `json:"creator_username"`
		Name        string `json:"name"`
		RoomID      int64  `json:"room_id"`
	}

	fmCreator struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}

	fmRoom struct {
		RoomID       int64        `json:"room_id"`      // 直播间ID
		Name         string       `json:"name"`         // 直播间名
		Announcement string       `json:"announcement"` // 公告
		Members      fmMembers    `json:"members"`      // 直播间成员
		Statistics   fmStatistics `json:"statistics"`   // 统计数据
		Status       fmStatus     `json:"status"`       // 状态信息
		CatalogID    int64        `json:"catalog_id"`   // 子分区ID
		GuildID      int64        `json:"guild_id"`     // 公会ID
		Medal        struct {
			Name string `json:"name"` // 粉丝牌名
		} `json:"medal"` // 粉丝牌
		Background struct {
			Enable   bool    `json:"enable"`
			ImageURL string  `json:"image_url"`
			Opacity  float64 `json:"opacity"`
		} `json:"background"` // 背景图
	}

	fmMembers struct {
		Admin []FmAdmin `json:"admin"` // 管理员
	}

	FmAdmin struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}

	// fmStatistics represents the statistics of the live room.
	fmStatistics struct {
		Accumulation   int64 `json:"accumulation"`    // 累计人数
		Vip            int64 `json:"vip"`             // 贵宾数量
		Score          int64 `json:"score"`           // 分数（热度）
		Revenue        int64 `json:"revenue"`         // 收益
		Online         int64 `json:"online"`          // 在线
		AttentionCount int64 `json:"attention_count"` // 关注数
	}

	fmStatus struct {
		Open     int64     `json:"open"`
		OpenTime int64     `json:"open_time"`
		Channel  fmChannel `json:"channel"`
	}

	fmChannel struct {
		Event    string `json:"event"`
		Platform string `json:"platform"`
		Time     int64  `json:"time"`
		Type     string `json:"type"`
	}

	FmRankMember struct {
		Revenue       int       `json:"revenue"`
		Rank          int       `json:"rank"`
		RankInvisible bool      `json:"rank_invisible"`
		UserId        int       `json:"user_id"`
		Username      string    `json:"username"`
		IconURL       string    `json:"iconurl"`
		Contribution  int       `json:"contribution"`
		Titles        []FmTitle `json:"titles"`
	}
)
