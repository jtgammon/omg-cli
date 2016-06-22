package metron_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type MetronAgent struct {

	/*EnableBuffer - Descr: DEPRECATED Default: false
*/
	EnableBuffer interface{} `yaml:"enable_buffer,omitempty"`

	/*DropsondeIncomingPort - Descr: Incoming port for dropsonde log messages Default: 3457
*/
	DropsondeIncomingPort interface{} `yaml:"dropsonde_incoming_port,omitempty"`

	/*MetronAgent - Descr: DEPRECATED Default: 10000
*/
	MetronAgent *MetronAgent `yaml:"metron_agent,omitempty"`

	/*Loggregator - Descr: CA root required for key/cert verification Default: 
*/
	Loggregator *Loggregator `yaml:"loggregator,omitempty"`

	/*SyslogDaemonConfig - Descr: Custom rule for syslog forward daemon Default: 
*/
	SyslogDaemonConfig *SyslogDaemonConfig `yaml:"syslog_daemon_config,omitempty"`

	/*Debug - Descr: boolean value to turn on verbose mode Default: false
*/
	Debug interface{} `yaml:"debug,omitempty"`

	/*BufferSize - Descr: DEPRECATED Default: 10000
*/
	BufferSize interface{} `yaml:"buffer_size,omitempty"`

	/*MetronEndpoint - Descr: Shared secret used to verify cryptographically signed dropsonde messages Default: <nil>
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

	/*PreferredProtocol - Descr: Preferred protocol to droppler (udp|tls) Default: udp
*/
	PreferredProtocol interface{} `yaml:"preferred_protocol,omitempty"`

	/*Tcp - Descr: The maximum time that a message can stay in the batching buffer before being flushed Default: 100
*/
	Tcp *Tcp `yaml:"tcp,omitempty"`

	/*Tls - Descr: TLS client key Default: 
*/
	Tls *Tls `yaml:"tls,omitempty"`

	/*Deployment - Descr: Name of deployment (added as tag on all outgoing metrics) Default: <nil>
*/
	Deployment interface{} `yaml:"deployment,omitempty"`

	/*Logrotate - Descr: The frequency in minutes which logrotate will rotate VM logs Default: 5
*/
	Logrotate *Logrotate `yaml:"logrotate,omitempty"`

	/*Zone - Descr: Availability zone where this agent is running Default: <nil>
*/
	Zone interface{} `yaml:"zone,omitempty"`

}