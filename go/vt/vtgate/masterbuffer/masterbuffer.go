/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package masterbuffer contains experimental logic to buffer master requests in VTGate.
Only statements outside of transactinos will be buffered (including the initial Begin
to start a transaction).

The reason why it might be useful to buffer master requests is during failovers:
the master vttablet can become unavailable for a few seconds. Upstream clients
(e.g., web workers) might not retry on failures, and instead may prefer for VTGate to wait for
a few seconds for the failover to complete. Thiis will block upstream callers for that time,
but will not return transient errors during the buffering time.
*/
package masterbuffer

import (
	"flag"
	"sync"
	"time"

	"vitess.io/vitess/go/stats"
	"vitess.io/vitess/go/vt/vterrors"

	querypb "vitess.io/vitess/go/vt/proto/query"
	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
	vtrpcpb "vitess.io/vitess/go/vt/proto/vtrpc"
)

var (
	enableFakeMasterBuffer = flag.Bool("enable_fake_master_buffer", false, "Enable fake master buffering.")
	bufferKeyspace         = flag.String("buffer_keyspace", "", "The name of the keyspace to buffer master requests on.")
	bufferShard            = flag.String("buffer_shard", "", "The name of the shard to buffer master requests on.")
	maxBufferSize          = flag.Int("max_buffer_size", 10, "The maximum number of master requests to buffer at a time.")
	fakeBufferDelay        = flag.Duration("fake_buffer_delay", 1*time.Second, "The amount of time that we should delay all master requests for, to fake a buffer.")

	bufferedRequestsAttempted  = stats.NewCounter("BufferedRequestsAttempted", "Count of buffered requests attempted")
	bufferedRequestsSuccessful = stats.NewCounter("BufferedRequestsSuccessful", "Count of successful buffered requests")
	// Use this lock when adding to the number of currently buffered requests.
	bufferMu         sync.Mutex
	bufferedRequests = stats.NewCounter("BufferedRequests", "Count of buffered requests")
)

// timeSleep can be mocked out in unit tests
var timeSleep = time.Sleep

// errBufferFull is the error returned a buffer request is rejected because the buffer is full.
var errBufferFull = vterrors.New(vtrpcpb.Code_UNAVAILABLE, "master request buffer full, rejecting request")

// FakeBuffer will pretend to buffer master requests in VTGate.
// Requests *will NOT actually be buffered*, they will just be delayed.
// This can be useful to understand what the impact of master request buffering will be
// on upstream callers. Once the impact is measured, it can be used to tweak parameter values
// for the best behavior.
// FakeBuffer should be called before a potential VtTablet Begin, otherwise it will increase transaction times.
func FakeBuffer(target *querypb.Target, inTransaction bool, attemptNumber int) error {
	if !*enableFakeMasterBuffer {
		return nil
	}
	// Don't buffer non-master traffic, requests that are inside transactions, or retries.
	if target.TabletType != topodatapb.TabletType_MASTER || inTransaction || attemptNumber != 0 {
		return nil
	}
	if target.Keyspace != *bufferKeyspace || target.Shard != *bufferShard {
		return nil
	}
	bufferedRequestsAttempted.Add(1)

	bufferMu.Lock()
	if int(bufferedRequests.Get()) >= *maxBufferSize {
		bufferMu.Unlock()
		return errBufferFull
	}
	bufferedRequests.Add(1)
	bufferMu.Unlock()

	defer bufferedRequestsSuccessful.Add(1)
	timeSleep(*fakeBufferDelay)
	// Don't need to lock for this, as there's no race when decrementing the count
	bufferedRequests.Add(-1)
	return nil
}
