package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Env struct {

	/*HttpsProxy - Descr: The https_proxy across the VMs Default: <nil>
*/
	HttpsProxy interface{} `yaml:"https_proxy,omitempty"`

	/*HttpProxy - Descr: The http_proxy across the VMs Default: <nil>
*/
	HttpProxy interface{} `yaml:"http_proxy,omitempty"`

	/*NoProxy - Descr: Set No_Proxy across the VMs Default: <nil>
*/
	NoProxy interface{} `yaml:"no_proxy,omitempty"`

}