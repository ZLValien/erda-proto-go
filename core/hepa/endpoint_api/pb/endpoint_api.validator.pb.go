// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: endpoint_api.proto

package pb

import (
	fmt "fmt"
	math "math"

	_ "github.com/erda-project/erda-proto-go/core/hepa/pb"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ChangeEndpointRootResponse) Validate() error {
	return nil
}
func (this *ChangeEndpointRootRequest) Validate() error {
	if this.EndpointApi != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndpointApi); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndpointApi", err)
		}
	}
	return nil
}
func (this *DeleteEndpointApiRequest) Validate() error {
	return nil
}
func (this *DeleteEndpointApiResponse) Validate() error {
	return nil
}
func (this *UpdateEndpointApiResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateEndpointApiRequest) Validate() error {
	if this.EndpointApi != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndpointApi); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndpointApi", err)
		}
	}
	return nil
}
func (this *CreateEndpointApiResponse) Validate() error {
	return nil
}
func (this *EndpointApi) Validate() error {
	if this.Method != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Method); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Method", err)
		}
	}
	return nil
}
func (this *CreateEndpointApiRequest) Validate() error {
	if this.EndpointApi != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndpointApi); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndpointApi", err)
		}
	}
	return nil
}
func (this *GetEndpointApisRequest) Validate() error {
	return nil
}
func (this *GetEndpointApisResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteEndpointRequest) Validate() error {
	return nil
}
func (this *DeleteEndpointResponse) Validate() error {
	return nil
}
func (this *UpdateEndpointResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateEndpointRequest) Validate() error {
	if this.Endpoint != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Endpoint); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Endpoint", err)
		}
	}
	return nil
}
func (this *CreateEndpointResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateEndpointRequest) Validate() error {
	if this.Endpoint != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Endpoint); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Endpoint", err)
		}
	}
	return nil
}
func (this *GetEndpointsNameRequest) Validate() error {
	return nil
}
func (this *Endpoint) Validate() error {
	return nil
}
func (this *GetEndpointsNameResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetEndpointsRequest) Validate() error {
	return nil
}
func (this *GetEndpointsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetEndpointRequest) Validate() error {
	return nil
}
func (this *GetEndpointResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
