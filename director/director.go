// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"open-match.dev/open-match/pkg/pb"

	agonesv1 "agones.dev/agones/pkg/apis/agones/v1"
	allocationv1 "agones.dev/agones/pkg/apis/allocation/v1"
	"agones.dev/agones/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	log.Println("Starting Director")
	for range time.Tick(time.Second) {
		if err := run(); err != nil {
			log.Println("Error running director:", err.Error())
		}
	}
}

func createOMBackendClient() (pb.BackendClient, func() error) {
	conn, err := grpc.Dial("om-backend.open-match.svc.cluster.local:50505", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewBackendClient(conn), conn.Close
}

func createAgonesClient() *versioned.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	agonesClient, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return agonesClient
}

// Customize the backend.FetchMatches request, the default one will return all tickets in the statestore
func createOMFetchMatchesRequest() *pb.FetchMatchesRequest {
	return &pb.FetchMatchesRequest{
		// om-function:50502 -> the internal hostname & port number of the MMF service in our Kubernetes cluster
		Config: &pb.FunctionConfig{
			Host: "mmf.default.svc.cluster.local",
			Port: 50502,
			Type: pb.FunctionConfig_GRPC,
		},
		Profiles: []*pb.MatchProfile{
			{
				Name: "1v1",
				Pools: []*pb.Pool{
					{
						Name: "everyone",
					},
				},
			},
		},
	}
}

func createAgonesGameServerAllocation() *allocationv1.GameServerAllocation {
	return &allocationv1.GameServerAllocation{
		Spec: allocationv1.GameServerAllocationSpec{
			Required: metav1.LabelSelector{
				MatchLabels: map[string]string{agonesv1.FleetNameLabel: "fleet-demogame-v1"},
			},
		},
	}
}

func createOMAssignTicketRequest(match *pb.Match, gsa *allocationv1.GameServerAllocation) *pb.AssignTicketsRequest {
	tids := []string{}
	for _, t := range match.GetTickets() {
		tids = append(tids, t.GetId())
	}

	return &pb.AssignTicketsRequest{
		TicketIds: tids,
		Assignment: &pb.Assignment{
			Connection: fmt.Sprintf("%s:%d", gsa.Status.Address, gsa.Status.Ports[0].Port),
		},
	}
}

func run() error {
	bc, closer := createOMBackendClient()
	defer closer()

	agonesClient := createAgonesClient()

	stream, err := bc.FetchMatches(context.Background(), createOMFetchMatchesRequest())
	if err != nil {
		return fmt.Errorf("fail to get response stream from backend.FetchMatches call: %w", err)
	}

	totalMatches := 0

	// Read the FetchMatches response. Each loop fetches an available game match that satisfies the match profiles.
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error streaming response from backend.FetchMatches call: %w", err)
		}

		gsa, err := agonesClient.AllocationV1().GameServerAllocations("default").Create(createAgonesGameServerAllocation())
		if err != nil {
			return fmt.Errorf("error requesting allocation: %w", err)
		}
		// TODO: This drops matches, instead of properly allocating them.  Tickets will only return to
		// the general pool after (iirc) one minute.  We should either tell OM that an assignment isn't
		// coming, or retry for a little while.
		if gsa.Status.State != allocationv1.GameServerAllocationAllocated {
			log.Printf("failed to allocate game server.\n")
			continue
		}

		if _, err = bc.AssignTickets(context.Background(), createOMAssignTicketRequest(resp.GetMatch(), gsa)); err != nil {
			// Corner case where we allocated a game server for players who left the queue after some waiting time.
			// Note that we may still leak some game servers when tickets got assigned but players left the queue before game frontend announced the assignments.
			if err = agonesClient.AgonesV1().GameServers("default").Delete(gsa.Status.GameServerName, &metav1.DeleteOptions{}); err != nil {
				return fmt.Errorf("error assigning tickets: %w", err)
			}
		}

		totalMatches++
	}

	log.Printf("Created and assigned %d matches", totalMatches)

	return nil
}
