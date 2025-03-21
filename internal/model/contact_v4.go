package model

// CREATE TABLE contact(
// id INTEGER PRIMARY KEY,
// username TEXT,
// local_type INTEGER,
// alias TEXT,
// encrypt_username TEXT,
// flag INTEGER,
// delete_flag INTEGER,
// verify_flag INTEGER,
// remark TEXT,
// remark_quan_pin TEXT,
// remark_pin_yin_initial TEXT,
// nick_name TEXT,
// pin_yin_initial TEXT,
// quan_pin TEXT,
// big_head_url TEXT,
// small_head_url TEXT,
// head_img_md5 TEXT,
// chat_room_notify INTEGER,
// is_in_chat_room INTEGER,
// description TEXT,
// extra_buffer BLOB,
// chat_room_type INTEGER
// )
type ContactV4 struct {
	UserName  string `json:"username"`
	Alias     string `json:"alias"`
	Remark    string `json:"remark"`
	NickName  string `json:"nick_name"`
	LocalType int    `json:"local_type"` // 2 群聊; 3 群聊成员(非好友); 5,6 企业微信;

	// ID                  int    `json:"id"`

	// EncryptUserName     string `json:"encrypt_username"`
	// Flag                int    `json:"flag"`
	// DeleteFlag          int    `json:"delete_flag"`
	// VerifyFlag          int    `json:"verify_flag"`
	// RemarkQuanPin       string `json:"remark_quan_pin"`
	// RemarkPinYinInitial string `json:"remark_pin_yin_initial"`
	// PinYinInitial       string `json:"pin_yin_initial"`
	// QuanPin             string `json:"quan_pin"`
	// BigHeadUrl          string `json:"big_head_url"`
	// SmallHeadUrl        string `json:"small_head_url"`
	// HeadImgMd5          string `json:"head_img_md5"`
	// ChatRoomNotify      int    `json:"chat_room_notify"`
	// IsInChatRoom        int    `json:"is_in_chat_room"`
	// Description         string `json:"description"`
	// ExtraBuffer         []byte `json:"extra_buffer"`
	// ChatRoomType        int    `json:"chat_room_type"`
}

func (c *ContactV4) Wrap() *Contact {
	return &Contact{
		UserName: c.UserName,
		Alias:    c.Alias,
		Remark:   c.Remark,
		NickName: c.NickName,
		IsFriend: c.LocalType != 3,
	}
}
