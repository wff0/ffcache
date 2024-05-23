package ffCache

import pb "ffCache/ffcache/ffcachepb"

type PeerPicker interface {
	PickPeer(key string) (PeerGetter, bool)
}

type PeerGetter interface {
	//Get(group string, key string) ([]byte, error)
	Get(in *pb.Request, out *pb.Response) error
}
