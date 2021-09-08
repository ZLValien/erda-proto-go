// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: org_client.proto

package pb

import (
	context "context"
	http1 "net/http"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// OrgClientServiceHandler is the server API for OrgClientService service.
type OrgClientServiceHandler interface {
	// POST /api/gateway/openapi/clients
	CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error)
	// DELETE /api/gateway/openapi/clients/{clientId}
	DeleteClient(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error)
	// GET /api/gateway/openapi/clients/{clientId}/credentials
	GetCredentials(context.Context, *GetCredentialsRequest) (*GetCredentialsResponse, error)
	// PATCH /api/gateway/openapi/clients/{clientId}/credentials
	UpdateCredentials(context.Context, *UpdateCredentialsRequest) (*UpdateCredentialsResponse, error)
	// POST /api/gateway/openapi/clients/{clientId}/packages/{packageId}
	GrantEndpoint(context.Context, *GrantEndpointRequest) (*GrantEndpointResponse, error)
	// DELETE /api/gateway/openapi/clients/{clientId}/packages/{packageId}
	RevokeEndpoint(context.Context, *RevokeEndpointRequest) (*RevokeEndpointResponse, error)
	// PUT /api/gateway/openapi/clients/{clientId}/packages/{packageId}/limits
	ChangeClientLimit(context.Context, *ChangeClientLimitRequest) (*ChangeClientLimitResponse, error)
}

// RegisterOrgClientServiceHandler register OrgClientServiceHandler to http.Router.
func RegisterOrgClientServiceHandler(r http.Router, srv OrgClientServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_CreateClient := func(method, path string, fn func(context.Context, *CreateClientRequest) (*CreateClientResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateClientRequest))
		}
		var CreateClient_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateClient_info = transport.NewServiceInfo("erda.core.hepa.org_client.OrgClientService", "CreateClient", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateClient_info)
				}
				r = r.WithContext(ctx)
				var in CreateClientRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["clientName"]; len(vals) > 0 {
					in.ClientName = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_DeleteClient := func(method, path string, fn func(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteClientRequest))
		}
		var DeleteClient_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteClient_info = transport.NewServiceInfo("erda.core.hepa.org_client.OrgClientService", "DeleteClient", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteClient_info)
				}
				r = r.WithContext(ctx)
				var in DeleteClientRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "clientId":
							in.ClientId = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetCredentials := func(method, path string, fn func(context.Context, *GetCredentialsRequest) (*GetCredentialsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetCredentialsRequest))
		}
		var GetCredentials_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetCredentials_info = transport.NewServiceInfo("erda.core.hepa.org_client.OrgClientService", "GetCredentials", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetCredentials_info)
				}
				r = r.WithContext(ctx)
				var in GetCredentialsRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "clientId":
							in.ClientId = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_UpdateCredentials := func(method, path string, fn func(context.Context, *UpdateCredentialsRequest) (*UpdateCredentialsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateCredentialsRequest))
		}
		var UpdateCredentials_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateCredentials_info = transport.NewServiceInfo("erda.core.hepa.org_client.OrgClientService", "UpdateCredentials", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateCredentials_info)
				}
				r = r.WithContext(ctx)
				var in UpdateCredentialsRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["clientSecret"]; len(vals) > 0 {
					in.ClientSecret = vals[0]
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "clientId":
							in.ClientId = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GrantEndpoint := func(method, path string, fn func(context.Context, *GrantEndpointRequest) (*GrantEndpointResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GrantEndpointRequest))
		}
		var GrantEndpoint_info transport.ServiceInfo
		if h.Interceptor != nil {
			GrantEndpoint_info = transport.NewServiceInfo("erda.core.hepa.org_client.OrgClientService", "GrantEndpoint", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GrantEndpoint_info)
				}
				r = r.WithContext(ctx)
				var in GrantEndpointRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "clientId":
							in.ClientId = val
						case "packageId":
							in.PackageId = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_RevokeEndpoint := func(method, path string, fn func(context.Context, *RevokeEndpointRequest) (*RevokeEndpointResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*RevokeEndpointRequest))
		}
		var RevokeEndpoint_info transport.ServiceInfo
		if h.Interceptor != nil {
			RevokeEndpoint_info = transport.NewServiceInfo("erda.core.hepa.org_client.OrgClientService", "RevokeEndpoint", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, RevokeEndpoint_info)
				}
				r = r.WithContext(ctx)
				var in RevokeEndpointRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "clientId":
							in.ClientId = val
						case "packageId":
							in.PackageId = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ChangeClientLimit := func(method, path string, fn func(context.Context, *ChangeClientLimitRequest) (*ChangeClientLimitResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ChangeClientLimitRequest))
		}
		var ChangeClientLimit_info transport.ServiceInfo
		if h.Interceptor != nil {
			ChangeClientLimit_info = transport.NewServiceInfo("erda.core.hepa.org_client.OrgClientService", "ChangeClientLimit", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ChangeClientLimit_info)
				}
				r = r.WithContext(ctx)
				var in ChangeClientLimitRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "clientId":
							in.ClientId = val
						case "packageId":
							in.PackageId = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateClient("POST", "/api/gateway/openapi/clients", srv.CreateClient)
	add_DeleteClient("DELETE", "/api/gateway/openapi/clients/{clientId}", srv.DeleteClient)
	add_GetCredentials("GET", "/api/gateway/openapi/clients/{clientId}/credentials", srv.GetCredentials)
	add_UpdateCredentials("PATCH", "/api/gateway/openapi/clients/{clientId}/credentials", srv.UpdateCredentials)
	add_GrantEndpoint("POST", "/api/gateway/openapi/clients/{clientId}/packages/{packageId}", srv.GrantEndpoint)
	add_RevokeEndpoint("DELETE", "/api/gateway/openapi/clients/{clientId}/packages/{packageId}", srv.RevokeEndpoint)
	add_ChangeClientLimit("PUT", "/api/gateway/openapi/clients/{clientId}/packages/{packageId}/limits", srv.ChangeClientLimit)
}
