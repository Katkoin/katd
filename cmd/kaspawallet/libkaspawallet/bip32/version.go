package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// katkoinMainnetPrivate is the version that is used for
// katkoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var katkoinMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// katkoinMainnetPublic is the version that is used for
// katkoin mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var katkoinMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// katkoinTestnetPrivate is the version that is used for
// katkoin testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var katkoinTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// katkoinTestnetPublic is the version that is used for
// katkoin testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var katkoinTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// katdevnetPrivate is the version that is used for
// katkoin devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var katdevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// katdevnetPublic is the version that is used for
// katkoin devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var katdevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// katkoinSimnetPrivate is the version that is used for
// katkoin simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var katkoinSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// katkoinSimnetPublic is the version that is used for
// katkoin simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var katkoinSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case katkoinMainnetPrivate:
		return katkoinMainnetPublic, nil
	case katkoinTestnetPrivate:
		return katkoinTestnetPublic, nil
	case katdevnetPrivate:
		return katdevnetPublic, nil
	case katkoinSimnetPrivate:
		return katkoinSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case katkoinMainnetPrivate:
		return true
	case katkoinTestnetPrivate:
		return true
	case katdevnetPrivate:
		return true
	case katkoinSimnetPrivate:
		return true
	}

	return false
}
