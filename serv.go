package supervisor

import context "context"

type nodeService struct {
	domain *domain
}

func newNodeService() *nodeService {
	return &nodeService{
		domain: newDomain(),
	}
}

func (s *nodeService) Add(ctx context.Context, req *AddNodeRequest) (*AddNodeResponse, error) {

}
