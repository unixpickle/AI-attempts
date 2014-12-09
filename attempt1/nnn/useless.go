package nnn

func PruneUseless(network *Network, threshold CycleTime) {
	Prune(network, func(link *Link) bool {
		age := network.Time - link.Life.Creation
		timeSinceFire := network.Time - link.Life.LastUsed
		return age < threshold || timeSinceFire < threshold
	})
}
