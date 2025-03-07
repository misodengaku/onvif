package media

type GetMetadataConfigurationFunction struct{}

func (function *GetMetadataConfigurationFunction) Request() interface{} {
	return &GetMetadataConfiguration{}
}

func (function *GetMetadataConfigurationFunction) Response() interface{} {
	return &GetMetadataConfigurationResponse{}
}

type GetMetadataConfigurationsFunction struct{}

func (function *GetMetadataConfigurationsFunction) Request() interface{} {
	return &GetMetadataConfigurations{}
}

func (function *GetMetadataConfigurationsFunction) Response() interface{} {
	return &GetMetadataConfigurationsResponse{}
}

type AddMetadataConfigurationFunction struct{}

func (function *AddMetadataConfigurationFunction) Request() interface{} {
	return &AddMetadataConfiguration{}
}

func (function *AddMetadataConfigurationFunction) Response() interface{} {
	return &AddMetadataConfigurationResponse{}
}

type RemoveMetadataConfigurationFunction struct{}

func (function *RemoveMetadataConfigurationFunction) Request() interface{} {
	return &RemoveMetadataConfiguration{}
}

func (function *RemoveMetadataConfigurationFunction) Response() interface{} {
	return &RemoveMetadataConfigurationResponse{}
}

type SetMetadataConfigurationFunction struct{}

func (function *SetMetadataConfigurationFunction) Request() interface{} {
	return &SetMetadataConfiguration{}
}

func (function *SetMetadataConfigurationFunction) Response() interface{} {
	return &SetMetadataConfigurationResponse{}
}

type GetCompatibleMetadataConfigurationsFunction struct{}

func (function *GetCompatibleMetadataConfigurationsFunction) Request() interface{} {
	return &GetCompatibleMetadataConfigurations{}
}

func (function *GetCompatibleMetadataConfigurationsFunction) Response() interface{} {
	return &GetCompatibleMetadataConfigurationsResponse{}
}

type GetMetadataConfigurationOptionsFunction struct{}

func (function *GetMetadataConfigurationOptionsFunction) Request() interface{} {
	return &GetMetadataConfigurationOptions{}
}

func (function *GetMetadataConfigurationOptionsFunction) Response() interface{} {
	return &GetMetadataConfigurationOptionsResponse{}
}

type GetProfilesFunction struct{}

func (function *GetProfilesFunction) Request() interface{} {
	return &GetProfiles{}
}

func (function *GetProfilesFunction) Response() interface{} {
	return &GetProfilesResponse{}
}

type GetStreamUriFunction struct{}

func (function *GetStreamUriFunction) Request() interface{} {
	return &GetStreamUri{}
}

func (function *GetStreamUriFunction) Response() interface{} {
	return &GetStreamUriResponse{}
}

type GetVideoEncoderConfigurationFunction struct{}

func (function *GetVideoEncoderConfigurationFunction) Request() interface{} {
	return &GetVideoEncoderConfiguration{}
}

func (function *GetVideoEncoderConfigurationFunction) Response() interface{} {
	return &GetVideoEncoderConfigurationResponse{}
}

type GetVideoEncoderConfigurationsFunction struct{}

func (function *GetVideoEncoderConfigurationsFunction) Request() interface{} {
	return &GetVideoEncoderConfigurations{}
}

func (function *GetVideoEncoderConfigurationsFunction) Response() interface{} {
	return &GetVideoEncoderConfigurationsResponse{}
}

type SetVideoEncoderConfigurationFunction struct{}

func (function *SetVideoEncoderConfigurationFunction) Request() interface{} {
	return &SetVideoEncoderConfiguration{}
}

func (function *SetVideoEncoderConfigurationFunction) Response() interface{} {
	return &SetVideoEncoderConfigurationResponse{}
}

type GetVideoEncoderConfigurationOptionsFunction struct{}

func (function *GetVideoEncoderConfigurationOptionsFunction) Request() interface{} {
	return &GetVideoEncoderConfigurationOptions{}
}

func (function *GetVideoEncoderConfigurationOptionsFunction) Response() interface{} {
	return &GetVideoEncoderConfigurationOptionsResponse{}
}

type AddPTZConfigurationFunction struct{}

func (function *AddPTZConfigurationFunction) Request() interface{} {
	return &AddPTZConfiguration{}
}

func (function *AddPTZConfigurationFunction) Response() interface{} {
	return &AddPTZConfigurationResponse{}
}

type RemovePTZConfigurationFunction struct{}

func (function *RemovePTZConfigurationFunction) Request() interface{} {
	return &RemovePTZConfiguration{}
}

func (function *RemovePTZConfigurationFunction) Response() interface{} {
	return &RemovePTZConfigurationResponse{}
}
