#!/bin/bash

rm -r ~/.hellocli
rm -r ~/.hellod
rm -r ~/.appli

hellod init helloNode --chain-id helloworld

hellocli config keyring-backend test

hellocli keys add me --keyring-backend test
hellocli keys add you --keyring-backend test


hellod add-genesis-account $(hellocli keys show me -a) 10000000000msgcoin,10000000000000stake
hellod add-genesis-account $(hellocli keys show you -a) 1000msgcoin,10000000000000stake

hellocli config chain-id helloworld
hellocli config output json
hellocli config indent true
hellocli config trust-node true

hellod gentx --name me --keyring-backend test
hellod collect-gentxs