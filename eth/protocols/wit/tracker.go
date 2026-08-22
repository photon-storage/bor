package wit

import (
	"time"

	"github.com/photon-storage/bor/p2p/tracker"
)

// requestTracker is a singleton tracker for wit/0 and newer request times.
var requestTracker = tracker.New(ProtocolName, 5*time.Minute)
