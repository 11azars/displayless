package gql

import (
	"context"
)

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
	panic("not implemented")
}
