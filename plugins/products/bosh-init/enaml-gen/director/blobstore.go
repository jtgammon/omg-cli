package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Blobstore struct {

	/*Provider - Descr: Provider of the blobstore used by director and agent (dav|simple|s3) Default: dav
*/
	Provider interface{} `yaml:"provider,omitempty"`

	/*SecretAccessKey - Descr: AWS secret_access_key used by s3 blobstore plugin Default: <nil>
*/
	SecretAccessKey interface{} `yaml:"secret_access_key,omitempty"`

	/*S3Region - Descr: Region of the blobstore used by s3 blobstore plugin Default: <nil>
*/
	S3Region interface{} `yaml:"s3_region,omitempty"`

	/*UseSsl - Descr: Whether the simple blobstore plugin should use SSL to connect to the blobstore server Default: true
*/
	UseSsl interface{} `yaml:"use_ssl,omitempty"`

	/*Address - Descr: Address of blobstore server used by simple blobstore plugin Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*BucketName - Descr: AWS S3 Bucket used by s3 blobstore plugin Default: <nil>
*/
	BucketName interface{} `yaml:"bucket_name,omitempty"`

	/*Agent - Descr: Password agent uses to connect to blobstore used by simple blobstore plugin Default: <nil>
*/
	Agent *BlobstoreAgent `yaml:"agent,omitempty"`

	/*CredentialsSource - Descr: AWS Credential Source (static / env_or_profile) Default: static
*/
	CredentialsSource interface{} `yaml:"credentials_source,omitempty"`

	/*SslVerifyPeer - Descr: Verify the SSL certificate used on the blobstore? Default: true
*/
	SslVerifyPeer interface{} `yaml:"ssl_verify_peer,omitempty"`

	/*SseKmsKeyId - Descr: AWS KMS key ID to use for object encryption. All GET and PUT requests for an object protected by AWS KMS will fail if not made via SSL or using SigV4. Default: <nil>
*/
	SseKmsKeyId interface{} `yaml:"sse_kms_key_id,omitempty"`

	/*Director - Descr: Username director uses to connect to blobstore used by simple blobstore plugin Default: <nil>
*/
	Director *BlobstoreDirector `yaml:"director,omitempty"`

	/*Port - Descr: Port of blobstore server used by simple blobstore plugin Default: 25250
*/
	Port interface{} `yaml:"port,omitempty"`

	/*ServerSideEncryption - Descr: Server-side encryption algorithm used when storing blobs in S3 (Optional - "AES256"|"aws:kms") Default: <nil>
*/
	ServerSideEncryption interface{} `yaml:"server_side_encryption,omitempty"`

	/*Host - Descr: Host of blobstore server used by simple blobstore plugin Default: <nil>
*/
	Host interface{} `yaml:"host,omitempty"`

	/*AccessKeyId - Descr: AWS access_key_id used by s3 blobstore plugin Default: <nil>
*/
	AccessKeyId interface{} `yaml:"access_key_id,omitempty"`

	/*S3SignatureVersion - Descr: Signature version of the blobstore used by s3 blobstore plugin (optional, if not provided the s3 client decides which version to use) Default: <nil>
*/
	S3SignatureVersion interface{} `yaml:"s3_signature_version,omitempty"`

	/*S3Port - Descr: Port of blobstore server used by s3 blobstore plugin Default: 443
*/
	S3Port interface{} `yaml:"s3_port,omitempty"`

}