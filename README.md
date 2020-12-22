# IPFS Batcher

Currently, FileCoin network miners are not that eager to make deals over the individual files, but rather want to deal in large pieces preferably more than 1 GiB in size. If you already have some files uploaded in pinned to your IPFS node making a directory from them could be problematic since you need to download files back and upload later as a folder. That is not that convenient afer all.

This repo is an implementation of batching that can be used if you want to make such batches with the files already stored and pinned in your IPFS node.


Ths repo is not intended to be used a library or dependency but holds possible implementation you can use as source to inspiration.

All the meaningful code you need can be found in [`ipfs.CreateBatch(...)`](https://github.com/kalambet/ipfsbatcher/blob/a3635b36abbfbfc48a9ca7575a7687fbf781ce4a/ipfs/ipfs.go#L28) method. Everything else is a proof-of-concept sandbox that checks the resulting piece size using [Lotus](https://github.com/filecoin-project/lotus) embedded methods.