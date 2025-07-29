package virtualclinic

type (
	Config struct {
		GVCID                   string            `toml:"gvc_id"`
		GVCOptionsSel           string            `toml:"gvc_more_options_sel"`
		ChromeUserDir           string            `toml:"chrome_user_dir"`
		ChromeBin               string            `toml:"chrome_bin"`
		WhatsAppRecipientSuffix string            `toml:"whatsapp_recipient_suffix"`
		Tunnels                 map[string]Tunnel `toml:"tunnel"`
		OtoCam                  int               `toml:"otocam_num"`
		WebCam                  int               `toml:"webcam_num"`
		OtoCamDevice            string            `toml:"otocam_dev"`
		PrinterDevice           string            `toml:"printer_dev"`
		CheckURL                string            `toml:"check_url"`
		BrowserWindowState      string            `toml:"browser_window_state"`
	}

	Tunnel struct {
		LocalPort  int    `toml:"local_port"`
		RemotePort int    `toml:"remote_port"`
		LocalHost  string `toml:"local_host"`
		RemoteHost string `toml:"remote_host"`
		User       string `toml:"user"`
	}
)

const (
	RPCTunID string = "rpc"
	SSHTunID string = "ssh"
	RDPTunID string = "rdp"
)
