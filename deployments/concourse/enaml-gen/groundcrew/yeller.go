package groundcrew 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Yeller struct {

	/*EnvironmentName - Descr: Environment name you wish to group errors under in yeller.
 Default: 
*/
	EnvironmentName interface{} `yaml:"environment_name,omitempty"`

	/*ApiKey - Descr: API key to output errors from Concourse to Yeller.
 Default: 
*/
	ApiKey interface{} `yaml:"api_key,omitempty"`

}