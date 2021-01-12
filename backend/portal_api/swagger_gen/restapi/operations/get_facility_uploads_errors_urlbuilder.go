// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetFacilityUploadsErrorsURL generates an URL for the get facility uploads errors operation
type GetFacilityUploadsErrorsURL struct {
	UploadID int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetFacilityUploadsErrorsURL) WithBasePath(bp string) *GetFacilityUploadsErrorsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetFacilityUploadsErrorsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetFacilityUploadsErrorsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/facility/uploads/{uploadId}/errors"

	uploadID := swag.FormatInt64(o.UploadID)
	if uploadID != "" {
		_path = strings.Replace(_path, "{uploadId}", uploadID, -1)
	} else {
		return nil, errors.New("uploadId is required on GetFacilityUploadsErrorsURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/divoc/admin/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetFacilityUploadsErrorsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetFacilityUploadsErrorsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetFacilityUploadsErrorsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetFacilityUploadsErrorsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetFacilityUploadsErrorsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetFacilityUploadsErrorsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}