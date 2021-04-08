package types

import (
//ethcmn "github.com/ethereum/go-ethereum/common"
)

func (suite *GenesisTestSuite) TestTransactionLogsValidate() {
	testCases := []struct {
		name    string
		txLogs  TransactionLogs
		expPass bool
	}{
		{
			"valid log",
			TransactionLogs{
				Hash: suite.hash[:],
				Logs: []*Log{
					{
						Address:     suite.address,
						Topics:      []string{string(suite.hash[:])},
						Data:        []byte("data"),
						BlockNumber: 1,
						TxHash:      suite.hash[:],
						TxIndex:     1,
						BlockHash:   suite.hash[:],
						Index:       1,
						Removed:     false,
					},
				},
			},
			true,
		},
		{
			"empty hash",
			TransactionLogs{
				Hash: []byte{},
			},
			false,
		},
		{
			"invalid log",
			TransactionLogs{
				Hash: suite.hash[:],
				Logs: []*Log{nil},
			},
			false,
		},
		{
			"hash mismatch log",
			TransactionLogs{
				Hash: suite.hash[:],
				Logs: []*Log{
					{
						Address:     suite.address,
						Topics:      []string{string(suite.hash[:])},
						Data:        []byte("data"),
						BlockNumber: 1,
						TxHash:      []byte("other_hash"),
						TxIndex:     1,
						BlockHash:   suite.hash[:],
						Index:       1,
						Removed:     false,
					},
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.txLogs.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}

func (suite *GenesisTestSuite) TestValidateLog() {
	testCases := []struct {
		name    string
		log     *Log
		expPass bool
	}{
		{
			"valid log",
			&Log{
				Address:     suite.address,
				Topics:      []string{string(suite.hash[:])},
				Data:        []byte("data"),
				BlockNumber: 1,
				TxHash:      suite.hash[:],
				TxIndex:     1,
				BlockHash:   suite.hash[:],
				Index:       1,
				Removed:     false,
			},
			true,
		},
		{
			"empty log", &Log{}, false,
		},
		{
			"zero address",
			&Log{
				Address: "",
			},
			false,
		},
		{
			"empty block hash",
			&Log{
				Address:   suite.address,
				BlockHash: []byte{},
			},
			false,
		},
		{
			"zero block number",
			&Log{
				Address:     suite.address,
				BlockHash:   suite.hash[:],
				BlockNumber: 0,
			},
			false,
		},
		{
			"empty tx hash",
			&Log{
				Address:     suite.address,
				BlockHash:   suite.hash[:],
				BlockNumber: 1,
				TxHash:      []byte{},
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.log.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}
