package utils

import "testing"

func TestSha256String(t *testing.T) {
	//t.Log(Md5String("9f735e0df9a1ddc702bf0a1a7b83033f9f7153a00c29de82cedadc9957289b05"))
	t.Log(HASHIDEncode(65))
	t.Log(HASHIDDecode("EgmJ4Ga7LQ"))
}
