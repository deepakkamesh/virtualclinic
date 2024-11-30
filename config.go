package virtualclinic

type (
	Config struct {
		GVCID                   string           `toml:"gvc_id"`
		GVCJoinNowSelector      string           `toml:"gvc_join_now_selector"`
		ChromeUserDir           string           `toml:"chrome_user_dir"`
		ChromeBin               string           `toml:"chrome_bin"`
		WhatsAppRecipientSuffix string           `toml:"whatsapp_recipient_suffix"`
		RelayServerAddress      string           `toml:"relay_server_address"`
		PortForwards            map[string]ports `toml:"portfwd"`
	}

	ports struct {
		LocalPort  int `toml:"local_port"`
		RemotePort int `toml:"remote_port"`
	}
)
