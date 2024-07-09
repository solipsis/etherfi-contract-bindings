package stakingManager

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

func DeployBlock() uint64 {
	return 17174441
}

func (s *StakingManager) BoundContract() *bind.BoundContract {
	return s.StakingManagerCaller.contract
}
