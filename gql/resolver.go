package gql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) ActivateNetwork(ctx context.Context, ssid string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) SetDefaultNetwork(ctx context.Context, ssid string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) ModifyWirelessNetwork(ctx context.Context, ssid string, b64password string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetWirelessNetworks(ctx context.Context) ([]*WirelessNetwork, error) {
	return []*WirelessNetwork{
		&WirelessNetwork{
			Ssid: "Schloss",
		},
		&WirelessNetwork{
			Ssid: "Hocturm",
		},
	}, nil
}
