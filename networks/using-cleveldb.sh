#!/bin/bash

make install CREATA_BUILD_OPTIONS="cleveldb"

creatad init "t6" --home ./t6 --chain-id t6

creatad unsafe-reset-all --home ./t6

mkdir -p ./t6/data/snapshots/metadata.db

creatad keys add validator --keyring-backend test --home ./t6

creatad add-genesis-account $(creatad keys show validator -a --keyring-backend test --home ./t6) 100000000stake --keyring-backend test --home ./t6

creatad gentx validator 100000000stake --keyring-backend test --home ./t6 --chain-id t6

creatad collect-gentxs --home ./t6

creatad start --db_backend cleveldb --home ./t6
