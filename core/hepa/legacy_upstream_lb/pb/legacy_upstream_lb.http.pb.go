// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: legacy_upstream_lb.proto

package pb

import (
	context "context"
	http1 "net/http"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// UpstreamLbServiceHandler is the server API for UpstreamLbService service.
type UpstreamLbServiceHandler interface {
	// PUT /api/gateway/target/online
	TargetOnline(context.Context, *TargetOnlineRequest) (*TargetOnlineResponse, error)
	// PUT /api/gateway/target/offline
	TargetOffline(context.Context, *TargetOfflineRequest) (*TargetOfflineResponse, error)
}

// RegisterUpstreamLbServiceHandler register UpstreamLbServiceHandler to http.Router.
func RegisterUpstreamLbServiceHandler(r http.Router, srv UpstreamLbServiceHandler, opts ...http.HandleOption) {
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

	add_TargetOnline := func(method, path string, fn func(context.Context, *TargetOnlineRequest) (*TargetOnlineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*TargetOnlineRequest))
		}
		var TargetOnline_info transport.ServiceInfo
		if h.Interceptor != nil {
			TargetOnline_info = transport.NewServiceInfo("erda.core.hepa.legacy_upstream_lb.UpstreamLbService", "TargetOnline", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, TargetOnline_info)
				}
				r = r.WithContext(ctx)
				var in TargetOnlineRequest
				if err := h.Decode(r, &in.Lb); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
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

	add_TargetOffline := func(method, path string, fn func(context.Context, *TargetOfflineRequest) (*TargetOfflineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*TargetOfflineRequest))
		}
		var TargetOffline_info transport.ServiceInfo
		if h.Interceptor != nil {
			TargetOffline_info = transport.NewServiceInfo("erda.core.hepa.legacy_upstream_lb.UpstreamLbService", "TargetOffline", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, TargetOffline_info)
				}
				r = r.WithContext(ctx)
				var in TargetOfflineRequest
				if err := h.Decode(r, &in.Lb); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
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

	add_TargetOnline("PUT", "/api/gateway/target/online", srv.TargetOnline)
	add_TargetOffline("PUT", "/api/gateway/target/offline", srv.TargetOffline)
}
