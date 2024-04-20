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

// KatkoinMainnetPrivate is the version that is used for
// Katkoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var KatkoinMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// KatkoinMainnetPublic is the version that is used for
// Katkoin mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var KatkoinMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// KatkoinTestnetPrivate is the version that is used for
// Katkoin testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var KatkoinTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// KatkoinTestnetPublic is the version that is used for
// Katkoin testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var KatkoinTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// KatkoinDevnetPrivate is the version that is used for
// Katkoin devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var KatkoinDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// KatkoinDevnetPublic is the version that is used for
// Katkoin devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var KatkoinDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// KatkoinSimnetPrivate is the version that is used for
// Katkoin simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var KatkoinSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// KatkoinSimnetPublic is the version that is used for
// Katkoin simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var KatkoinSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case KatkoinMainnetPrivate:
		return KatkoinMainnetPublic, nil
	case KatkoinTestnetPrivate:
		return KatkoinTestnetPublic, nil
	case KatkoinDevnetPrivate:
		return KatkoinDevnetPublic, nil
	case KatkoinSimnetPrivate:
		return KatkoinSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case KatkoinMainnetPrivate:
		return true
	case KatkoinTestnetPrivate:
		return true
	case KatkoinDevnetPrivate:
		return true
	case KatkoinSimnetPrivate:
		return true
	}

	return false
}
