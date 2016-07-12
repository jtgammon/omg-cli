package blobstore 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Director struct {

	/*Password - Descr: Password director must use to access blobstore via HTTP Basic Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*User - Descr: Username director must use to access blobstore via HTTP Basic Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

}