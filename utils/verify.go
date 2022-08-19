package utils

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	SortIdVerify           = Rules{"SortId": {NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	LoginVerify            = Rules{"UserID": {NotEmpty()}, "Password": {NotEmpty()}}
	LoginAppVerify         = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"UserID": {NotEmpty()}, "Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "College": {NotEmpty()}, "Labels": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty(), Gt("0")}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "DataScope": {NotEmpty()}, "Level": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	UpdatePwdVerify        = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
	DeleteAuthorityVerify  = Rules{"AuthorityId": {NotEmpty(), Gt("3")}}
	DeleteUserVerify       = Rules{"ID": {NotEmpty(), Gt("2")}}
	BookLabelVerify        = Rules{"BookId": {NotEmpty()}}
)
