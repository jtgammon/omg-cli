package health_monitor 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Smtp struct {

	/*Password - Descr: Password for SMTP Authentication (optional, use in conjuction with hm.smtp.auth) Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Port - Descr: Port of the SMTP server to connect to (e.g. 25, 465, or 587) Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Auth - Descr: SMTP Authentication type (optional, only "plain" is supported) Default: <nil>
*/
	Auth interface{} `yaml:"auth,omitempty"`

	/*User - Descr: User for SMTP Authentication (optional, use in conjuction with hm.smtp.auth) Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*Domain - Descr: SMTP EHLO domain (typically server's fully qualified domain name, e.g. "bosh.cfapps.io") Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Tls - Descr: Use STARTTLS (optional) Default: <nil>
*/
	Tls interface{} `yaml:"tls,omitempty"`

	/*From - Descr: Email of sender, e.g. "bosh-director@cfapps.io" Default: <nil>
*/
	From interface{} `yaml:"from,omitempty"`

	/*Host - Descr: Address of the SMTP server to connect to (e.g. "smtp.gmail.com") Default: <nil>
*/
	Host interface{} `yaml:"host,omitempty"`

}