package ipfs

import (
	"bufio"
	"context"
	"os"

	"github.com/filecoin-project/go-commp-utils/writer"
	"github.com/filecoin-project/lotus/api"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	ipfsfiles "github.com/ipfs/go-ipfs-files"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	ipfsIO "github.com/ipfs/go-unixfs/io"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipld/go-car"
	"github.com/multiformats/go-multiaddr"
)

func InitIPFSClient(addr string) (*httpapi.HttpApi, error) {
	ma, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	return httpapi.NewApi(ma)
}

func CreateBatch(ctx context.Context, i *httpapi.HttpApi, cids []cid.Cid) (cid.Cid, error) {
	batch := ipfsIO.NewDirectory(i.Dag())
	for _, c := range cids {
		node, err := i.Dag().Get(ctx, c)
		if err != nil {
			return cid.Undef, err
		}
		err = batch.AddChild(ctx, uuid.New().String(), node)
		if err != nil {
			return cid.Undef, err
		}
	}

	node, err := batch.GetNode()
	if err != nil {
		return cid.Undef, err
	}

	err = i.Dag().Add(ctx, node)
	if err != nil {
		return cid.Undef, err
	}

	err = i.Dag().Pinning().Add(ctx, node)
	if err != nil {
		return cid.Undef, err
	}
	return node.Cid(), nil
}

func ClientDealPieceCID(ctx context.Context, i *httpapi.HttpApi, root cid.Cid) (api.DataCIDSize, error) {
	dag := i.Dag()

	w := &writer.Writer{}
	bw := bufio.NewWriterSize(w, int(writer.CommPBuf))

	err := car.WriteCar(ctx, dag, []cid.Cid{root}, w)
	if err != nil {
		return api.DataCIDSize{}, err
	}

	if err := bw.Flush(); err != nil {
		return api.DataCIDSize{}, err
	}

	dataCIDSize, err := w.Sum()
	return api.DataCIDSize(dataCIDSize), err
}

func UploadAndPin(ctx context.Context, i *httpapi.HttpApi, pth string) (cid.Cid, error) {
	file, err := os.Open(pth)
	if err != nil {
		return cid.Undef, err
	}
	defer file.Close()

	resp, err := i.Unixfs().Add(ctx, ipfsfiles.NewReaderFile(file), options.Unixfs.Pin(true))
	if err != nil {
		return cid.Undef, err
	}
	return resp.Cid(), nil
}
