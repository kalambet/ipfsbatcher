package ipfsbatcher

import (
	"context"
	"ipfsbatcher/ipfs"
	"log"

	"code.cloudfoundry.org/bytefmt"
	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
)

func Upload(ctx context.Context, i *httpapi.HttpApi) []cid.Cid {
	paths := []string{
		"./data/test-file-1",
		"./data/test-file-2",
		"./data/test-file-3",
		"./data/test-file-4",
	}

	var cids []cid.Cid
	for _, p := range paths {
		c, err := ipfs.UploadAndPin(ctx, i, p)
		if err != nil {
			log.Printf("error uploading file: %s", err)
			continue
		}
		cids = append(cids, c)
	}
	return cids
}

func Do() error {
	ctx := context.Background()
	i, err := ipfs.InitIPFSClient("/ip4/127.0.0.1/tcp/5001")
	if err != nil {
		return err
	}

	cids := Upload(ctx, i)

	batch, err := ipfs.CreateBatch(ctx, i, cids)
	if err != nil {
		return err
	}

	size, err := ipfs.ClientDealPieceCID(ctx, i, batch)
	if err != nil {
		return err
	}

	log.Printf("payload size: %s", bytefmt.ByteSize(uint64(size.PayloadSize)))
	log.Printf("piece size: %s", bytefmt.ByteSize(uint64(size.PieceSize)))
	return nil
}
