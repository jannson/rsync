package rsync

import (
	"bytes"
	//"io"
	"testing"
)

func testSign(t *testing.T) {
	for _, s := range ss {
		dst := new(bytes.Buffer)
		err := testGenSign(s, dst, len(s), 3)
		if err != nil {
			t.Fail()
		} else {
			t.Logf("%x %v", dst.Bytes(), s)
			testLoadSign(t, dst)
		}
	}
}

func testGenSign(s string, dst *bytes.Buffer, srcLen, blockLen int) error {
	src := bytes.NewBufferString(s)
	err := GenSign(src, int64(srcLen), uint32(blockLen), dst)
	return err
}

func testLoadSign(t *testing.T, sigRd *bytes.Buffer) {
	sig, err := LoadSign(sigRd, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("src length: %d blocks: %d tlen: %d block len: %d sum len: %d magic: 0x%x\n",
		sig.flength, sig.count, sig.flength, sig.block_len, sig.strong_sum_len, sig.magic)
	for sum, block_sigs := range sig.block_sigs {
		t.Logf(" block sum: 0x%x:\n", sum)
		for _, block_sig := range block_sigs {
			t.Logf("    block index: %d block weak sum: 0x%x strong sum: %x\n", block_sig.i, block_sig.wsum, block_sig.ssum)
		}
	}
}
