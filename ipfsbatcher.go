package ipfsbatcher

import (
	"context"
	"ipfsbatcher/generator"
	"ipfsbatcher/ipfs"
	"log"
	"os"
	"path"
	"path/filepath"

	"code.cloudfoundry.org/bytefmt"
	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
)

func Upload(ctx context.Context, i *httpapi.HttpApi, paths []string) []cid.Cid {
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
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dataPath := path.Join(pwd, "data")
	err = os.Mkdir(dataPath, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	paths := []string{
		filepath.Join(dataPath, "test-file-1"),
		filepath.Join(dataPath, "test-file-2"),
		filepath.Join(dataPath, "test-file-3"),
		filepath.Join(dataPath, "test-file-4"),
		filepath.Join(dataPath, "test-file-5"),
		filepath.Join(dataPath, "test-file-6"),
	}

	err = generator.NewTestData(paths)
	if err != nil {
		return err
	}

	ctx := context.Background()
	i, err := ipfs.InitIPFSClient("/ip4/127.0.0.1/tcp/5001")
	if err != nil {
		return err
	}

	cids := Upload(ctx, i, paths)

	batch, err := ipfs.CreateBatch(ctx, i, cids)
	if err != nil {
		return err
	}
	log.Printf("batch cid: %s", batch.String())

	size, err := ipfs.ClientDealPieceCID(ctx, i, batch)
	if err != nil {
		return err
	}

	log.Printf("piece cid: %s", size.PieceCID)
	log.Printf("payload size: %s", bytefmt.ByteSize(uint64(size.PayloadSize)))
	log.Printf("piece size: %s", bytefmt.ByteSize(uint64(size.PieceSize)))
	return nil
}
