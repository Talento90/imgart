package health

import (
	"encoding/json"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// Status represents application health
type Status struct {
	Service         string            `json:"service"`
	Uptime          string            `json:"up_time"`
	StartTime       string            `json:"start_time"`
	MemoryAllocated uint64            `json:"memory_allocated"`
	HealthCheckers  map[string]string `json:"health_checkers"`
}

// Checker checks if the health of a service (database, external service)
type Checker interface {
	Check() error
}

// Health interface
type Health interface {
	// GetStatus return the current status of the application
	GetStatus() *Status
	// RegisterChecker register a service to check their health
	RegisterChecker(name string, check Checker)
	// ServeHTTP handler for http applications
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// New returns a new Health
func New() Health {
	return &health{
		mutex:     &sync.Mutex{},
		startTime: time.Now(),
		checkers:  map[string]Checker{},
	}
}

type health struct {
	mutex     *sync.Mutex
	startTime time.Time
	checkers  map[string]Checker
}

// RegisterChecker regist external dependencies health
func (h *health) RegisterChecker(name string, check Checker) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.checkers[name] = check
}

// GetStatus method returns the current application health status
func (h *health) GetStatus() *Status {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	checkers := map[string]string{}

	for k, c := range h.checkers {
		err := c.Check()

		if err != nil {
			checkers[k] = err.Error()
		} else {
			checkers[k] = "OK"
		}
	}

	return &Status{
		Service:         "gorpo",
		Uptime:          time.Now().Sub(h.startTime).String(),
		StartTime:       h.startTime.Format(time.RFC3339),
		MemoryAllocated: mem.Alloc,
		HealthCheckers:  checkers,
	}
}

// ServeHTTP that returns the health status
func (h *health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := h.GetStatus()

	bytes, _ := json.Marshal(status)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
