package content

var getstartedTOC = TOC{
	{
		Name: "Unpack & Setup",
		Slug: "unpack-and-setup",
		Sections: []Section{
			{Name: "Check Content", ID: "check"},
			{Name: "Setting Up The Machine", ID: "setup"},
		},
	},
	{
		Name: "Install The Software",
		Slug: "install-the-software",
		Sections: []Section{
			{Name: "Download", ID: "download"},
			{Name: "Verify Image", ID: "verify"},
			{Name: "Write Image", ID: "write"},
			{Name: "Start The Controller", ID: "start"},
		},
	},
	{
		Name: "Load Seed & Wallet",
		Slug: "load-seed-and-wallet",
		Sections: []Section{
			{Name: "Load Seed", ID: "input"},
			{Name: "Scan Output Descriptor or Public Key", ID: "scan-output-descriptor-or-public-key"},
		},
	},
	{
		Name: "Engrave Plate",
		Slug: "engrave-plate",
		Sections: []Section{
			{Name: "Onscreen Instructions", ID: "instructions"},
			{Name: "In-app Tips", ID: "tips"},
		},
	},
	{
		Name: "Verify Backup",
		Slug: "verify-backup",
		Sections: []Section{
			{Name: "Verify Singlesig", ID: "singlesig"},
			{Name: "Verify Multisig", ID: "multisig"},
		},
	},
	{
		Name: "Perform Recovery",
		Slug: "perform-recovery",
		Sections: []Section{
			{Name: "Recover Singlesig", ID: "singlesig"},
			{Name: "Recover Multisig", ID: "multisig"},
		},
	},
}

var manualTOC = TOC{
	{
		Name: "Firmware Upgrade",
		Slug: "firmware-upgrade",
		Sections: []Section{
			{Name: "Requirements", ID: "requirements"},
			{Name: "Download", ID: "download"},
			{Name: "Connect", ID: "connect"},
			{Name: "Upload", ID: "upload"},
		},
	},
	{
		Name: "Engrave Descriptor (Android)",
		Slug: "engrave-descriptor-android",
		Sections: []Section{
			{Name: "Install NFC Tools", ID: "install"},
			{Name: "Transfer Descriptor", ID: "transfer"},
		},
	},
	{
		Name: "Engrave Descriptor (iOS)",
		Slug: "engrave-descriptor-ios",
		Sections: []Section{
			{Name: "Install NFC Tools", ID: "install"},
			{Name: "Transfer Descriptor", ID: "transfer"},
		},
	},
}

var aboutTOC = TOC{
	{Name: "How It Started", Slug: "how-it-started", Sections: []Section{}},
	{Name: "Our Goal", Slug: "our-goal", Sections: []Section{}},
	{Name: "The SeedHammer Ethos", Slug: "the-seedhammer-ethos", Sections: []Section{}},
	{Name: "Attitude to Privacy", Slug: "attitude-to-privacy", Sections: []Section{}},
	{Name: "Contact", Slug: "contact", Sections: []Section{}},
}

type Nav struct {
	Name   string
	URL    string
	Target string
}

var topNav = []Nav{
	{Name: "Shop", URL: "/shop"},
	{Name: "Guide", URL: "/doc/manual"},
	{Name: "FAQ", URL: "/faq"},
	{Name: "Articles", URL: "/article/"},
	{Name: "Software", URL: "https://github.com/seedhammer/seedhammer/releases/latest", Target: "_blank"},
	{Name: "About", URL: "/about/"},
}
