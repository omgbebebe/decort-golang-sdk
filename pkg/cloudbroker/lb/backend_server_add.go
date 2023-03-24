package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for add server definition to the backend
type BackendServerAddRequest struct {
	// ID of the load balancer instance to BackendServerAdd
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Must match one of the existing backens - name of the backend to add servers to
	// Required: true
	BackendName string `url:"backendName" json:"backendName" validate:"required"`

	// Must be unique among all servers defined for this backend - name of the server definition to add
	// Required: true
	ServerName string `url:"serverName" json:"serverName" validate:"required"`

	// IP address of the server
	// Required: true
	Address string `url:"address" json:"address" validate:"required"`

	// Port number on the server
	// Required: true
	Port uint64 `url:"port" json:"port" validate:"required"`

	// Set to disabled if this server should be used regardless of its state
	// Required: false
	Check string `url:"check,omitempty" json:"check,omitempty"`

	// Interval in milliseconds between two consecutive availability checks of the server that is considered available
	// Required: false
	Inter uint64 `url:"inter,omitempty" json:"inter,omitempty"`

	// Interval in milliseconds between two consecutive checks to restore
	// the availability of a server that is currently considered unavailable
	// Required: false
	DownInter uint64 `url:"downinter,omitempty" json:"downinter,omitempty"`

	// Number of checks that the server must pass in order to get
	// the available status and be included in the balancing scheme again
	// Required: false
	Rise uint64 `url:"rise,omitempty" json:"rise,omitempty"`

	// Number of consecutive failed availability checks,
	// after which the previously considered available server receives the status of unavailable and
	// is temporarily excluded from the balancing scheme
	// Required: false
	Fall uint64 `url:"fall,omitempty" json:"fall,omitempty"`

	// Interval in milliseconds from the moment the server receives the available status,
	// after which the number of actually allowed connections to this server will be returned to 100% of the set limit
	// Required: false
	SlowStart uint64 `url:"slowstart,omitempty" json:"slowstart,omitempty"`

	// Limit of simultaneous connections to the server. When this limit is reached, the server is temporarily excluded from the balancing scheme
	// Required: false
	MaxConn uint64 `url:"maxconn,omitempty" json:"maxconn,omitempty"`

	// Limit of connections waiting in the queue. When this limit is reached, all subsequent connections will be forwarded to other servers
	// Required: false
	MaxQueue uint64 `url:"maxqueue,omitempty" json:"maxqueue,omitempty"`

	// Server weight for use in weight balancing algorithms
	// Required: false
	Weight uint64 `url:"weight,omitempty" json:"weight,omitempty"`
}

// BackendServerAdd adds server definition to the backend on the specified load balancer
func (lb LB) BackendServerAdd(ctx context.Context, req BackendServerAddRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/backendServerAdd"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
